package rpc

import (
	"context"

	"github.com/flowck/cargotracker_ddd_golang/internal/app/commands"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/logs"

	"github.com/flowck/cargotracker_ddd_golang/internal/app"
	pb "github.com/flowck/cargotracker_ddd_golang/internal/ports/rpc/static"
)

type handlers struct {
	application *app.App
	pb.UnimplementedCargoTrackerServer
}

func (h handlers) BookNewCargo(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	logs.Info(h.application.Commands.BookNewCargo.Execute(ctx, commands.BookNewCargo{}))
	return &pb.Empty{}, nil
}
