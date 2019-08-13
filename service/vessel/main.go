package main

import (
	"errors"
	"log"

	pb "github.com/GaVender/micro/proto/vessel"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, v := range repo.vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("not available vessel")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, res *pb.Response) error {
	v, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	res.Vessel = v
	return nil
}

func main() {
	vessel := pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500}
	vessels := []*pb.Vessel{&vessel}
	repo := &VesselRepository{vessels}
	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
