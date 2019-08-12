package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/GaVender/micro/proto/consignment"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
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
	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microClient.DefaultClient)
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
