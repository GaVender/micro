package main

import (
	pb "github.com/GaVender/micro/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shipy"
	consignmentCollection = "consignment"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

func (repo *ConsignmentRepository) GetAll() (all []*pb.Consignment, err error) {
	err = repo.collection().Find(nil).All(&all)
	return
}

func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}
