package repositories

import (
	"context"
	"driver-location-api/db"
	"driver-location-api/domain/constants"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	"driver-location-api/domain/repository"
	err "driver-location-api/error"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const collectionDriverLocation = "driver_locations"

type driverRepository struct {
	db *mongo.Database
}

func NewDriverRepository(m *db.MongoRepository) repository.DriverRepository {
	return &driverRepository{db: m.GetMongoDB()}
}

func (dr driverRepository) SaveDriverLocation(di model.DriverInfo) (*model.DriverInfo, *err.Error) {
	fmt.Println(di)
	_, e := dr.db.Collection(collectionDriverLocation).InsertOne(context.Background(), di)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError(constants.DataNotSaved)
	}
	return &di, nil
}

func (dr driverRepository) SaveDriverLocationFile(di []model.DriverInfo) *err.Error {
	var d []interface{}
	for _, t := range di {
		d = append(d, t)
	}
	_, e := dr.db.Collection(collectionDriverLocation).InsertMany(context.Background(), d)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return err.ServerError(constants.DataNotSaved)
	}
	dr.createIndex("location", "2dsphere")
	return nil
}

func (dr driverRepository) GetNearDrivers(location core.Location, radius int) ([]*model.DriverInfo, *err.Error) {
	ctx := context.Background()

	filter := bson.D{
		{"location",
			bson.D{
				{"$nearSphere", bson.D{
					{"$geometry", location},
					{"$maxDistance", radius},
				}},
			}},
	}

	var drivers []*model.DriverInfo
	cursor, e := dr.db.Collection(collectionDriverLocation).Find(ctx, filter)

	if e != nil {
		return drivers, err.ServerError(constants.CouldNotGetDriverData)
	}
	e = cursor.All(ctx, &drivers)
	if e != nil {
		return drivers, err.NotFoundError(constants.CouldNotGetDriverData)
	}

	return drivers, nil
}

func (dr driverRepository) createIndex(field string, indexType string) *err.Error {
	_, e := dr.db.Collection(collectionDriverLocation).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{Key: field, Value: bsonx.String(indexType)}},
	})
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return err.ServerError(constants.IndexNotCreated)
	}
	return nil
}
