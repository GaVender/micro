package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GaVender/micro/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5001"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(c *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, c)
	return c, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetAll(context.Context, *pb.EmptyRequest) (*pb.Response, error) {
	return &pb.Response{Consignments: s.repo.GetAll()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	repo := Repository{}
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{&repo})
	reflection.Register(s)

	log.Println("grpc serve...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
