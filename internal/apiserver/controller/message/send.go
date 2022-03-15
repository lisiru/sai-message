package message

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"sai/common"
	"sai/pkg/code"
	"sai/pkg/response"
)




func (me *MessageController) Send(c *gin.Context)  {
	sendRequestParam:=&common.SendRequestParam{}
	if err:=c.ShouldBindJSON(sendRequestParam);err!=nil{
		response.WriteResponse(c,errors.WithCode(code.ErrParamNotValid,err.Error()),nil)
		return
	}
	err:=me.service.Message().SendMessage(c,*sendRequestParam)
	if err != nil {
		response.WriteResponse(c,err,nil)
		return
	}
	response.WriteResponse(c,nil,nil)


}

func (me *MessageController) SendKafkaTest(c *gin.Context)  {
	me.service.Message().SendTest(c)

}
