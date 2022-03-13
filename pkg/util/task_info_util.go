package util

import (
	"fmt"
	"strings"
	"time"
)

var typeFlag = 1000000
func GenerateUrl(url string,templateId int64,templateType int) string {
	businessId:=GenBusinessId(templateId,templateType)
	if strings.Index(url,"?")==-1 {
		return url+"?track_code_bid="+businessId
	}
	return url+"&track_code_bid="+businessId


}

// 生成businessId

func GenBusinessId(templateId int64,templateType int) string {
	today:=time.Now().Format("20060503")
	return fmt.Sprintf("%d%s",templateType*typeFlag+int(templateId),today)
}
