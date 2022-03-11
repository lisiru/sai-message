package response

import (
	"sai/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string`json:"message"`
	Data interface{} `json:"data"`
}

func WriteResponse(c *gin.Context,err error,data interface{})  {
	if err != nil {
		logger.Errorf("%#+v",err)
		coder:=errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(),Response{
			Code: coder.Code(),
			Message: coder.String(),
			Data: coder.Reference(),
		})
		return
	}
	c.JSON(http.StatusOK,Response{
		Code: 0,
		Message: "",
		Data: data,
	})

}
