package message

import (
	"github.com/gin-gonic/gin"
	"sai/common"
)

type BatchSendRequest struct {
	Code string `json:"code"` // 执行业务类型(默认填写 "send")
	MessageTemplateId string `json:"message_template_id" binding:"required"` // 消息模板Id
	MessageParamList []common.MessageParam `json:"message_param_list"` // 消息相关的参数

}

func (me *MessageController) BatchSend(c *gin.Context)  {

}
