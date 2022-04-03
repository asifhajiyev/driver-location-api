package repositories

import (
	"context"
	"driver-location-api/db"
	"driver-location-api/domain/constants"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	"driver-location-api/domain/repository"
	err "driver-location-api/error"
	"driver-location-api/logger"
	"fmt"
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
	logger.Info("SaveDriverLocation.begin")
	fmt.Println(di)
	_, e := dr.db.Collection(collectionDriverLocation).InsertOne(context.Background(), di)
	if e != nil {
		logger.Error("SaveDriverLocation.error", e)
		return nil, err.ServerError(constants.ErrorDataNotSaved)
	}
	logger.Info("SaveDriverLocation.end", &di)
	return &di, nil
}

func (dr driverRepository) SaveDriverLocationFile(di []model.DriverInfo) *err.Error {
	logger.Info("SaveDriverLocationFile.begin")
	var d []interface{}
	for _, t := range di {
		d = append(d, t)
	}
	_, e := dr.db.Collection(collectionDriverLocation).InsertMany(context.Background(), d)
	if e != nil {
		logger.Error("SaveDriverLocationFile.error", e)
		return err.ServerError(constants.ErrorDataNotSaved)
	}
	dr.createIndex("location", "2dsphere")
	logger.Info("SaveDriverLocationFile.end")
	return nil
}

func (dr driverRepository) GetNearDrivers(location core.Location, radius int) ([]*model.DriverInfo, *err.Error) {
	logger.Info("GetNearDrivers.begin")
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
		logger.Error("GetNearDrivers.error", e)
		return drivers, err.ServerError(constants.ErrorCouldNotGetDriverData)
	}
	e = cursor.All(ctx, &drivers)
	if e != nil {
		logger.Error("GetNearDrivers.error", e)
		return drivers, err.NotFoundError(constants.ErrorCouldNotGetDriverData)
	}
	logger.Info("GetNearDrivers.end", drivers)
	return drivers, nil
}

func (dr driverRepository) createIndex(field string, indexType string) *err.Error {
	_, e := dr.db.Collection(collectionDriverLocation).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{Key: field, Value: bsonx.String(indexType)}},
	})
	if e != nil {
		logger.Error("createIndex.error", e)
		return err.ServerError(constants.ErrorIndexNotCreated)
	}
	return nil
}
