package service

import (
	"context"
	"encoding/json"
	"github.com/marmotedu/errors"
	"regexp"
	"sai/common"
	"sai/model"
	"sai/pkg/code"
	"sai/pkg/util"
	"sai/pkg/util/kafka"
	"sai/store"
	"strings"
)

type Processor interface {
	Process(ctx context.Context, processContext *common.ProcessContext) error
}

type PreParamCheckAction struct{}

func (p *PreParamCheckAction) Process(ctx context.Context, processContext *common.ProcessContext) error {
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

func (a *AfterParamCheckAction) Process(ctx context.Context, processContext *common.ProcessContext) error {
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

func (asseble *AssembleAction) Process(ctx context.Context, processContext *common.ProcessContext) error {
	messageTemplateId := processContext.SendTaskModel.MessageTemplateId
	where := make(map[string]interface{})
	where["id"] = messageTemplateId

	messageTemplateInfo, err := asseble.store.MessageTemplate().GetMessageTemplate(ctx, where)
	if err != nil {
		return err
	}
	if messageTemplateInfo.IsDeleted == common.TRUE {
		return errors.WithCode(code.ErrMessageTemplateNotFound, "")

	}
	//
	taskInfoList, err:= asseble.buildTaskInfo(processContext.SendTaskModel, messageTemplateInfo)
	if err != nil {
		return errors.WithCode(code.ErrMessageTemplateNotFound, "")
	}
	processContext.SendTaskModel.TaskInfo = taskInfoList
	return nil

}

func (assemble *AssembleAction) buildTaskInfo(sendTaskModel common.SendTaskModel, messageTemplateInfo *model.MessageTemplate) ([]common.TaskInfo,error) {
	messageParamList := sendTaskModel.MessageParamList
	var taskInfoList []common.TaskInfo
	for _, v := range messageParamList {
		contentValue,err:=assemble.getContentValue(messageTemplateInfo,v)
		if err != nil {
			return nil, err
		}
		taskInfoList = append(taskInfoList, common.TaskInfo{
			MessageTemplateId: messageTemplateInfo.Id,
			BusinessId:        "",
			Receiver:          strings.Split(v.Receiver, ","),
			IdType:            messageTemplateInfo.IdType,
			SendChannel:       messageTemplateInfo.SendChannel,
			MsgType:           messageTemplateInfo.MsgType,
			Content:           contentValue,
			SendAccount:       messageTemplateInfo.SendAccount,
		})
	}
	return taskInfoList,nil

}

func (assemble *AssembleAction) getContentValue(messageTemplateInfo *model.MessageTemplate, messageParam common.MessageParam) (common.Content,error) {
	content:=common.Content{}
	channel := messageTemplateInfo.SendChannel
	switch channel {
	case common.CHANNEL_TYPE_SMS:
		contentModel:=common.SmsContent{}
		err := json.Unmarshal([]byte(messageTemplateInfo.MsgContent), &contentModel)
		if err != nil {
			return content, errors.WithCode(code.ErrParamNotValid,"")
		}
		contentModel.Content=messageParam.Variable[strings.ReplaceAll(contentModel.Content,common.CONTENT_REPLACE_OLD_STR,common.CONTENT_REPLACE_NEW_STR)]
		contentModel.Expire=messageParam.Variable[strings.ReplaceAll(contentModel.Url,"?","")]
		content.SmsContent=contentModel
		url:=content.SmsContent.Url
		if len(url) !=0{
			content.SmsContent.Url=util.GenerateUrl(url,messageTemplateInfo.Id,messageTemplateInfo.TemplateType)
		}
	default:

	}
	return content,nil
	//contentStruct := common.ChannelContentMap[channel]
	////variables := messageParam.Variable
	////msgContent := json.Unmarshal([]byte(messageTemplateInfo.MsgContent), contentStruct)
	//structType := reflect.TypeOf(contentStruct)
	//fieldsLen := structType.NumField()
	//structValue := reflect.ValueOf(contentStruct)
	//for i := 0; i < fieldsLen; i++ {
	//	columnName := structType.Field(i).Name
	//	if columnName != "Url" {
	//		continue
	//	}
	//	urlValue := structValue.FieldByName(columnName)
	//	if len(urlValue.String()) != 0 {
	//		resultUrl := util.GenerateUrl(urlValue.String(), messageTemplateInfo.Id, messageTemplateInfo.TemplateType)
	//		urlValue.SetString(resultUrl)
	//	}
	//
	//}
	//return contentStruct

}

type SendMqAction struct{}

func (s *SendMqAction) Process(ctx context.Context, processContext *common.ProcessContext) error {
	message, err := json.Marshal(processContext.SendTaskModel.TaskInfo)
	if err != nil {
		return errors.WithCode(code.ErrParamNotValid, "")
	}
	producer:=kafka.NewProducer(ctx)
	producer.Send(kafka.TopicName,string(message))
	return nil

}
