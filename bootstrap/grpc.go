package bootstrap

import (
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GRPCServer interface {
	Register(server *grpc.Server)
}

type GRPCServerOperator struct {
	port       string
	server     GRPCServer
	opts       []grpc.ServerOption
	grpcServer *grpc.Server
}

func NewGRPCOperator(
	port string,
	server GRPCServer,
	opts []grpc.ServerOption,
) *GRPCServerOperator {
	return &GRPCServerOperator{
		port:   port,
		server: server,
		opts:   opts,
	}
}

func (g *GRPCServerOperator) Start() {
	log.Infof("Starting GRPC server on port %s", g.port)
	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		log.Fatalf("Failed to start GRPC server on port: %v", err)
	}
	grpcServer := grpc.NewServer(g.opts...)
	g.grpcServer = grpcServer
	g.server.Register(grpcServer)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (g *GRPCServerOperator) GracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	g.grpcServer.GracefulStop()
}
