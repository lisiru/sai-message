package service

import (
	"context"
	"encoding/json"
	"github.com/marmotedu/errors"
	"reflect"
	"regexp"
	"strings"
	"sai/common"
	"sai/model"
	"sai/pkg/code"
	"sai/pkg/util"
	"sai/pkg/util/kafka"
	"sai/store"
)

type Processor interface {
	Process(ctx context.Context, processContext common.ProcessContext) error
}

type PreParamCheckAction struct{}

func (p *PreParamCheckAction) Process(ctx context.Context, processContext common.ProcessContext) error {
	messageParamList := processContext.SendTaskModel.MessageParamList
	if len(messageParamList) == 0 {
		return errors.WithCode(code.ErrParamNotValid, "", nil)
	}
	// 过滤reciver == nil的messageParam
	newMessageParamList := make([]common.MessageParam, 0, len(messageParamList))
	for _, v := range messageParamList {
		if v.Receiver == "" {
			continue
		}
		newMessageParamList = append(newMessageParamList, v)
	}
	if len(newMessageParamList) == 0 {
		return errors.WithCode(code.ErrParamNotValid, "", nil)

	}
	processContext.SendTaskModel.MessageParamList = newMessageParamList
	return nil

}

type AfterParamCheckAction struct {
	store store.Factory
}

func (a *AfterParamCheckAction) Process(ctx context.Context, processContext common.ProcessContext) error {
	taskInfo := processContext.SendTaskModel.TaskInfo
	idType := taskInfo[0].IdType
	sendChannel := taskInfo[0].SendChannel

	// todo 测试是否有修改原来结构体的值
	if idType == common.ID_TYPE_PHONE && sendChannel == common.CHANNEL_TYPE_SMS {
		// 过滤不符合规则的手机号
		for _, task := range taskInfo {
			index := 0
			for _, reveiver := range task.Receiver {
				if isMatch, err := regexp.Match(common.PHONE_REGEX_EXP, []byte(reveiver)); err == nil && isMatch {
					task.Receiver[index] = reveiver
					index++
				}
			}
			task.Receiver = task.Receiver[:index]
			taskIndex := 0
			if len(task.Receiver) != 0 {
				taskInfo[taskIndex] = task
				taskIndex++
			}
		}
	}
	if len(taskInfo) == 0 {
		return errors.WithCode(code.ErrParamNotValid, "")

	}
	return nil

}

type AssembleAction struct {
	store store.Factory
}

func (asseble *AssembleAction) Process(ctx context.Context, processContext common.ProcessContext) error {
	messageTemplateId := processContext.SendTaskModel.MessageTemplateId
	where := make(map[string]interface{})
	where["message_template_id"] = messageTemplateId

	messageTemplateInfo, err := asseble.store.MessageTemplate().GetMessageTemplate(ctx, where)
	if err != nil {
		return err
	}
	if messageTemplateInfo.IsDeleted == common.TRUE {
		return errors.WithCode(code.ErrMessageTemplateNotFound, "")

	}
	//
	taskInfoList := asseble.buildTaskInfo(processContext.SendTaskModel, messageTemplateInfo)
	processContext.SendTaskModel.TaskInfo = taskInfoList
	return nil

}

func (assemble *AssembleAction) buildTaskInfo(sendTaskModel common.SendTaskModel, messageTemplateInfo *model.MessageTemplate) []common.TaskInfo {
	messageParamList := sendTaskModel.MessageParamList
	var taskInfoList []common.TaskInfo
	for _, v := range messageParamList {
		taskInfoList = append(taskInfoList, common.TaskInfo{
			MessageTemplateId: messageTemplateInfo.Id,
			BusinessId:        "",
			Receiver:          strings.Split(v.Receiver, ","),
			IdType:            messageTemplateInfo.IdType,
			SendChannel:       messageTemplateInfo.SendChannel,
			MsgType:           messageTemplateInfo.MsgType,
			Content:           assemble.getContentValue(messageTemplateInfo, v),
			SendAccount:       messageTemplateInfo.SendAccount,
		})
	}
	return taskInfoList

}

func (assemble *AssembleAction) getContentValue(messageTemplateInfo *model.MessageTemplate, messageParam common.MessageParam) interface{} {
	channel := messageTemplateInfo.SendChannel
	contentStruct := common.ChannelContentMap[channel]
	//variables := messageParam.Variable
	//msgContent := json.Unmarshal([]byte(messageTemplateInfo.MsgContent), contentStruct)
	structType := reflect.TypeOf(contentStruct)
	fieldsLen := structType.NumField()
	structValue := reflect.ValueOf(contentStruct)
	for i := 0; i < fieldsLen; i++ {
		columnName := structType.Field(i).Name
		if columnName != "Url" {
			continue
		}
		urlValue := structValue.FieldByName(columnName)
		if len(urlValue.String()) != 0 {
			resultUrl := util.GenerateUrl(urlValue.String(), messageTemplateInfo.Id, messageTemplateInfo.TemplateType)
			urlValue.SetString(resultUrl)
		}

	}
	return contentStruct

}

type SendMqAction struct{}

func (s *SendMqAction) Process(ctx context.Context, processContext common.ProcessContext) error {
	message, err := json.Marshal(processContext.SendTaskModel.TaskInfo)
	if err != nil {
		return errors.WithCode(code.ErrParamNotValid, "")
	}
	kafka.InitProducer()
	defer kafka.ProduceClose()
	kafka.Send(kafka.TopicName, string(message))
	return nil

}
