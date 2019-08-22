package main

import (
	"github.com/GaVender/micro/db"
	pb "github.com/GaVender/micro/proto/vessel"
	"golang.org/x/net/context"
)

type service struct {}

func (s *service) GetRepo() Repository {
	return &VesselRepository{db.GetSession()}
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, res *pb.Response) (err error) {
	repo := s.GetRepo()
	defer repo.Close()
	res.Vessel, err = repo.FindAvailable(spec)
	return
}

func(s *service) Create(ctx context.Context, ves *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(ves); err != nil {
		return err
	}

	res.Created = true
	res.Vessel = ves
	return nil
}
