package bootstrap

type App struct {
	HttpServer   HttpServer
	GRPCOperator *GRPCServerOperator
}

func (a *App) Run() {
	go a.HttpServer.Start()
	if a.GRPCOperator != nil {
		go a.GRPCOperator.Start()
	}
	a.HttpServer.GracefulShutdown()
}
