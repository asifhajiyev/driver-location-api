package repository

import (
	"context"
	"driver-location-api/db"
	err "driver-location-api/error"
	"driver-location-api/model/core"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "driver_locations"

type DriverLocationRepo interface {
	SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error)
	UploadDriverLocationFile() *err.Error
	GetNearestDriver() *err.Error
}
type driverLocationRepo struct {
	c *mongo.Collection
}

func NewDriverLocationRepo(m *db.MongoRepository) DriverLocationRepo {
	return &driverLocationRepo{c: m.GetCollection(collectionName)}
}

func (dlr driverLocationRepo) SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error) {
	_, e := dlr.c.InsertOne(context.Background(), dl)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError("data could not be saved")
	}
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	/*_, e := collection.InsertOne(ctx, dl)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError("data could not be saved")
	}*/
	return &dl, nil
}

func (dlr driverLocationRepo) UploadDriverLocationFile() *err.Error {
	return nil
}

func (dlr driverLocationRepo) GetNearestDriver() *err.Error {
	return nil
}
