package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
	_ "net/http/pprof"
	"sai/pkg/logger"
	"sai/pkg/middleware"
	"time"
)

type GenericAPIServer struct {
	middlewares []string
	mode        string

	InsecureServingInfo *InsecureServingInfo


	*gin.Engine
	healthz         bool
	// wrapper for gin.Engine

	insecureServer *http.Server
}

func initGenericAPIServer(s *GenericAPIServer) {
	// do some setup
	// s.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}
// InstallAPIs install generic apis.
func (s *GenericAPIServer) InstallAPIs() {
	// install healthz handler
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {

		})
	}


}

func (s *GenericAPIServer) Setup() {
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}
func (s *GenericAPIServer) InstallMiddlewares() {


	for _, m := range s.middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			logger.Warnf("can not find middleware: %s", m)

			continue
		}

		logger.Infof("install middleware: %s", m)
		s.Use(mw)
	}
}

func (s *GenericAPIServer) Run(stopCh <-chan struct{}) error {
	// For scalability, use custom HTTP configuration mode here
	s.insecureServer = &http.Server{
		Addr:    s.InsecureServingInfo.Address,
		Handler: s,


	}

	go func() {
		_ = http.ListenAndServe(":7070", nil)
	}()

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	eg.Go(func() error {
		logger.Infof("Start to listening the incoming requests on http address: %s", s.InsecureServingInfo.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("启动服务出错 %s",err.Error())

		}
		return nil


	})


	// Ping the server to make sure the router is working.
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if s.healthz {
		fmt.Print("健康检查")
	}
	<-stopCh
	if err:=s.insecureServer.Shutdown(timeoutCtx);err!=nil{

	}
	if err:=eg.Wait();err!=nil{

	}
	return nil


}
