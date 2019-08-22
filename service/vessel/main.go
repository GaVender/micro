package main

import (
	"log"

	"github.com/GaVender/micro/db"
	pb "github.com/GaVender/micro/proto/vessel"
	"github.com/micro/go-micro"
)

func main() {
	defer db.CloseSession()

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
