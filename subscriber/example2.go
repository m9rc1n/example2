package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	example2 "github.com/maszhin/example2/proto/example2"
)

type Example2 struct{}

func (e *Example2) Handle(ctx context.Context, msg *example2.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example2.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
