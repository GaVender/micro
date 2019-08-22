package main

import (
	"log"

	"github.com/GaVender/micro/db"
	pb "github.com/GaVender/micro/proto/consignment"
	vesselProto "github.com/GaVender/micro/proto/vessel"
	"golang.org/x/net/context"
)

type service struct {
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{db.GetSession()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.GetWeight(),
		Capacity:  int32(len(req.GetContainers())),
	})
	log.Println("Found vessel: ", vesselResponse)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.GetVessel().GetId()
	err = repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetAll(c context.Context, req *pb.EmptyRequest, res *pb.Response) (err error) {
	repo := s.GetRepo()
	defer repo.Close()

	res.Consignments, err = repo.GetAll()
	return
}
