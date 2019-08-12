package main

import (
	"context"
	"log"

	pb "github.com/GaVender/micro/proto/consignment"
	micro "github.com/micro/go-micro"
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

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetAll(c context.Context, req *pb.EmptyRequest, res *pb.Response) error {
	res.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	/*lis, err := net.Listen("tcp", port)
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
	}*/

	repo := &Repository{}
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		)
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
