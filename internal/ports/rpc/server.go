package rpc

import (
	"fmt"
	"net"

	"github.com/flowck/cargotracker_ddd_golang/internal/common/logs"
	pb "github.com/flowck/cargotracker_ddd_golang/internal/ports/rpc/static"
	"google.golang.org/grpc"
)

type Port struct {
	server *grpc.Server
	logger *logs.Logger
}

func NewPort(logger *logs.Logger) Port {
	server := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterCargoTrackerServer(server, handlers{})

	return Port{
		server: server,
		logger: logger,
	}
}

func (p Port) Start(port int16) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	p.logger.WithFields(logs.Fields{"host": fmt.Sprintf("http://localhost:%d", port)}).Info("grpc server will run shortly:")
	return p.server.Serve(listener)
}

func (p Port) Stop() {
	p.server.GracefulStop()
	p.logger.Info("grpc port has been stopped gracefully")
}
