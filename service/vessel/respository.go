package main

import (
	"errors"

	pb "github.com/GaVender/micro/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shipy"
	vesselCollection = "vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (v *pb.Vessel, err error) {
	err = repo.collection().Find(bson.M{
		"capacity": bson.M{"$gte": spec.Capacity},
		"weight":   bson.M{"$gte": spec.Capacity},
	}).One(&v)
	if err == mgo.ErrNotFound {
		err = errors.New("not available vessel")
	}
	return
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	err := repo.collection().Insert(vessel)
	return err
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}
