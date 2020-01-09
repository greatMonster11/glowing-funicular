package main

import (
	"context"
	"log"
	"time"
	hello "../proto"
	"github.com/micro/go-micro"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rep *hello.Response) error {
	log.Print("Recieved Say.Hello request - second greeting service")
	rep.Msg = "Hello" + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	hello.RegisterSayHandler(service.Server(), new(Say))
	if err := service.Say(); err != nil {
		log.Fatal("errro starting service: ", err)
		return
	}
}