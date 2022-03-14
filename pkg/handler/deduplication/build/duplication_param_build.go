package build

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/marmotedu/errors"
	"sai/common"
	"sai/pkg/code"
	"sai/pkg/util"
)

type build interface {
	paramBuild(deduplicationConfig string, info common.TaskInfo) (common.DeduplicationParam, error)
}

type abstractBuild struct {
	build
}

func NewAbstractBuild(deDuplicationType int) *abstractBuild {
	return &abstractBuild{build: SelectParamBuild(deDuplicationType)}
}

func (a *abstractBuild) Build(deduplicationConfig string, taskInfo common.TaskInfo) (common.DeduplicationParam, error) {
	return a.paramBuild(deduplicationConfig, taskInfo)
}

var DE_DUPLICATION_CONFIG_PRE = "deduplication_%s"

func (b *abstractBuild) getParamsFromConfig(deduplicationType int, deduplicationConfig string, taskInfo common.TaskInfo) (common.DeduplicationParam, error) {
	//deduplicationParam := common.DeduplicationParam{}

	var res map[string]common.DeduplicationParam

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(deduplicationConfig), &res)
	if err != nil {
		return common.DeduplicationParam{}, errors.WithCode(code.ErrParamNotValid, err.Error())

	}
	key:=fmt.Sprintf(DE_DUPLICATION_CONFIG_PRE,util.IntToString(deduplicationType))
	currentParam := res[key]
	currentParam.TaskInfo = taskInfo
	return currentParam, nil

	// 解析deduplicationConfig json字符串
	//js, _ := simplejson.NewJson([]byte(deduplicationConfig))
	//coruentJson, err := js.Get(DE_DUPLICATION_CONFIG_PRE + string(deduplicationType)).MarshalJSON()
	//if err != nil {
	//
	//}
	//err = jsoniter.Unmarshal(coruentJson, &deduplicationParam)
	//if err != nil {
	//
	//}
	//deduplicationParam.TaskInfo = taskInfo
	//return deduplicationParam

}
