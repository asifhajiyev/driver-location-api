package repository

import (
	"driver-location-api/db"
	err "driver-location-api/error"
	"driver-location-api/model/core"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "driver_locations"

type DriverLocationRepo interface {
	SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error)
	UploadDriverLocationFile() *err.Error
	GetNearestDriver() *err.Error
}
type DriverLocationRepoImpl struct {
	c *mongo.Collection
}

func NewDriverLocationRepo(m *db.MongoRepository) DriverLocationRepo {
	return &DriverLocationRepoImpl{c: m.GetCollection(collectionName)}
}

func (dlr *DriverLocationRepoImpl) SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error) {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	/*_, e := collection.InsertOne(ctx, dl)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError("data could not be saved")
	}*/
	return &dl, nil
}

func (dlr *DriverLocationRepoImpl) UploadDriverLocationFile() *err.Error {
	return nil
}

func (dlr *DriverLocationRepoImpl) GetNearestDriver() *err.Error {
	return nil
}
