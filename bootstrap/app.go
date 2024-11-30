package bootstrap

type App struct {
	HttpServer HttpServer
	GrpcServer GRPCServer
}

func (a *App) Run() {
	// Its concurrent because the grpc server is blocking
	if a.GrpcServer != nil {
		go a.GrpcServer.Start()
	}
	a.HttpServer.Start()
}
