package main

import (
	"log"

	"github.com/GaVender/micro/db"
	pb "github.com/GaVender/micro/proto/consignment"
	vesselProto "github.com/GaVender/micro/proto/vessel"
	"github.com/micro/go-micro"
)

func main() {
	defer db.CloseSession()

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{vesselClient})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
