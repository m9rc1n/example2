package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/maszhin/example2/handler"
	"github.com/maszhin/example2/subscriber"

	example2 "github.com/maszhin/example2/proto/example2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.example2"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example2.RegisterExample2Handler(service.Server(), new(handler.Example2))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.example2", service.Server(), new(subscriber.Example2))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.example2", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
