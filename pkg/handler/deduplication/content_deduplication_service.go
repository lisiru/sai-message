package deduplication

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sai/common"
)

type ContentDeDuplication struct {

}

func init()  {
	TypeMap[common.DE_DUPLICATION_TYPE_CONTENT]=&ContentDeDuplication{}
}

/**
内容去重构建key
key: md5(templateId+receiver+content)
 */
func (contentDe *ContentDeDuplication) deduplicationSingleKey(taskInfo common.TaskInfo,receiver string) string {

	contentString, _ :=json.Marshal(taskInfo.Content)
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(string(taskInfo.MessageTemplateId) + receiver + string(contentString))))
	return md5String
}

