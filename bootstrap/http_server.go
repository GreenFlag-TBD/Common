package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer interface {
	Start()
	AddRoutes(routes ...RouteInterface)
	AddMiddleware(middlewares ...fiber.Handler)
	GracefulShutdown()
}

type FiberHttpServer struct {
	Port        string
	routes      []RouteInterface
	app         *fiber.App
	middlewares []fiber.Handler
}

// NewFiberHttpServer creates a new instance of FiberHttpServer
func NewFiberHttpServer(port string) *FiberHttpServer {
	return &FiberHttpServer{Port: port}
}

// Start starts the http server with the providede routes
// It uses the fiber framework
func (f *FiberHttpServer) Start() {
	app := fiber.New()
	//Inject the routes
	for _, route := range f.routes {
		route.Register(app)
	}
	//Inject the middlewares
	for _, middleware := range f.middlewares {
		app.Use(middleware)
	}

	log.Fatal(app.Listen(":" + f.Port))
}

// AddRoutes adds routes to the http server
func (f *FiberHttpServer) AddRoutes(routes ...RouteInterface) {
	f.routes = append(f.routes, routes...)
}

// AddMiddleware adds middleware to the http server
func (f *FiberHttpServer) AddMiddleware(middlewares ...fiber.Handler) {
	f.middlewares = append(f.middlewares, middlewares...)

}

// GracefulShutdown stops the http server

func (f *FiberHttpServer) GracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	if err := f.app.Shutdown(); err != nil {
		log.Fatal("Error shutting down the server: ", err)
	}

}
