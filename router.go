package api_server_demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"sai/cache/redis"
	"sai/internal/apiserver/controller/message"
	"sai/pkg/response"
	"sai/store/mysql"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}
func installController(g *gin.Engine) *gin.Engine {
	g.GET("test", func(context *gin.Context) {
		res := make(map[string]string)
		res["user"] = "lisr"
		response.WriteResponse(context, nil, res)

	})
	g.GET("testTime", func(context *gin.Context) {
		time.Sleep(10 * time.Second)
		context.String(http.StatusOK, "success")
	})

	// 获取mysql的
	storeInstance, _ := mysql.GetMySQLFactoryOr(nil)
	cacheInstance, _ := redis.NewRedisFactoryOr(nil)

	messageGroup := g.Group("/message")
	{
		messageController := message.NewMessageController(storeInstance, cacheInstance)
		messageGroup.POST("/send", messageController.Send)
		messageGroup.POST("/batch_send",messageController.BatchSend)
		messageGroup.GET("/send_test",messageController.SendKafkaTest)


	}

	return g
}
