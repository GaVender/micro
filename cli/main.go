package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/GaVender/micro/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:5001"
	defaultFileName = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := json.Unmarshal(data, &consignment); err != nil {
		log.Println(err)
		return nil, err
	}
	return &consignment, nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)
	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		panic(consignment)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		panic(err)
	}
	log.Println(r)

	getAll, err := client.GetAll(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
