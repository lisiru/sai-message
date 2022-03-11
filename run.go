package api_server_demo

import "sai/config"

func Run(cfg *config.Config,stopCh <-chan struct{}) error  {
	server,err:=createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.PrepareRun().Run(stopCh)
}
