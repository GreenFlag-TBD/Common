package bootstrap

type App struct {
	HttpServer HttpServer
}

func (a *App) Run() {
	go a.HttpServer.Start()
	a.HttpServer.GracefulShutdown()
}
