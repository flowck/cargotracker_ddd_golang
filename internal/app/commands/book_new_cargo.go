package commands

import (
	"context"
	"time"
)

type BookNewCargo struct{}

type bookNewCargoHandler struct {
}

func NewBookNewCargo() bookNewCargoHandler {
	return bookNewCargoHandler{}
}

func (b bookNewCargoHandler) Execute(ctx context.Context, cmd BookNewCargo) error {
	time.Sleep(time.Second * 2)
	return nil
}
