package main

import (
	"log"

	pb "github.com/GaVender/micro/proto/consignment"
	vesselProto "github.com/GaVender/micro/proto/vessel"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) Create(c *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, c)
	return c, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.GetWeight(),
		Capacity:  int32(len(req.GetContainers())),
	})
	log.Println("Found vessel: ", vesselResponse)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.GetVessel().GetId()
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

	repo := &ConsignmentRepository{}
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
