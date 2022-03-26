package repository

import (
	"context"
	"driver-location-api/db"
	err "driver-location-api/error"
	"driver-location-api/model/core"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "driver_locations"

type DriverLocationRepo interface {
	SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error)
	SaveDriverLocationFile(dl []core.DriverLocation) *err.Error
	GetNearestDriver() *err.Error
}
type driverLocationRepo struct {
	c *mongo.Collection
}

func NewDriverLocationRepo(m *db.MongoRepository) DriverLocationRepo {
	return &driverLocationRepo{c: m.GetCollection(collectionName)}
}

func (dlr driverLocationRepo) SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error) {
	fmt.Println(dl)
	_, e := dlr.c.InsertOne(context.Background(), dl)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError("data could not be saved")
	}
	return &dl, nil
}

func (dlr driverLocationRepo) SaveDriverLocationFile(dl []core.DriverLocation) *err.Error {

	var d []interface{}
	for _, t := range dl {
		d = append(d, t)
	}
	_, e := dlr.c.InsertMany(context.Background(), d)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return err.ServerError("data could not be saved")
	}
	return nil
}

func (dlr driverLocationRepo) GetNearestDriver() *err.Error {
	return nil
}
