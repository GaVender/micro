package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pbConsignment "github.com/GaVender/micro/proto/consignment"
	pbVessel "github.com/GaVender/micro/proto/vessel"
	microClient "github.com/micro/go-micro/client"
)

const (
	defaultFileName = "consignment.json"
)

func parseFile(file string) (*pbConsignment.Consignment, error) {
	var consignment pbConsignment.Consignment
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
	vesselClient := pbVessel.NewVesselServiceClient("go.micro.srv.vessel", microClient.DefaultClient)
	respVessel, err := vesselClient.Create(context.Background(), &pbVessel.Vessel{
		Id: "vessel001",
		Name: "Kane's Salty Secret",
		MaxWeight: 200000,
		Capacity: 500,
	})
	if err != nil {
		panic(err)
	}
	log.Println(respVessel)

	consignmentClient := pbConsignment.NewShippingServiceClient("go.micro.srv.consignment", microClient.DefaultClient)

	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		panic(consignment)
	}

	consignmentResp, err := consignmentClient.CreateConsignment(context.Background(), consignment)
	if err != nil {
		panic(err)
	}
	log.Println(consignmentResp)

	getAll, err := consignmentClient.GetAll(context.Background(), &pbConsignment.EmptyRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
