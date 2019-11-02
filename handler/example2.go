package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	example2 "github.com/maszhin/example2/proto/example2"
)

type Example2 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example2) Call(ctx context.Context, req *example2.Request, rsp *example2.Response) error {
	log.Log("Received Example2.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example2) Stream(ctx context.Context, req *example2.StreamingRequest, stream example2.Example2_StreamStream) error {
	log.Logf("Received Example2.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example2.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example2) PingPong(ctx context.Context, stream example2.Example2_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example2.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
