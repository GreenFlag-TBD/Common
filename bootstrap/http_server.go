package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HttpServer interface {
	Start()
	AddRoutes(routes ...RouteInterface)
	AddMiddleware(middlewares ...fiber.Handler)
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
func (f *FiberHttpServer) Start(cfg ...fiber.Config) {
	app := fiber.New(cfg...)
	//Inject the routes
	for _, route := range f.routes {
		route.Register(app)
	}
	//Inject the middlewares
	for _, middleware := range f.middlewares {
		app.Use(middleware)
	}

	log.Fatal(app.Listen(f.Port))
}

// AddRoutes adds routes to the http server
func (f *FiberHttpServer) AddRoutes(routes ...RouteInterface) {
	f.routes = append(f.routes, routes...)
}

// AddMiddleware adds middleware to the http server
func (f *FiberHttpServer) AddMiddleware(middlewares ...fiber.Handler) {
	f.middlewares = append(f.middlewares, middlewares...)

}
