package repositories

import (
	"context"
	"driver-location-api/db"
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

type driverLocationRepo struct {
	db *mongo.Database
}

func NewDriverLocationRepo(m *db.MongoRepository) repository.DriverLocationRepo {
	return &driverLocationRepo{db: m.GetMongoDB()}
}

func (dlr driverLocationRepo) SaveDriverLocation(di model.DriverInfo) (*model.DriverInfo, *err.Error) {
	fmt.Println(di)
	_, e := dlr.db.Collection(collectionDriverLocation).InsertOne(context.Background(), di)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError("data could not be saved")
	}
	return &di, nil
}

func (dlr driverLocationRepo) SaveDriverLocationFile(di []model.DriverInfo) *err.Error {
	var d []interface{}
	for _, t := range di {
		d = append(d, t)
	}
	_, e := dlr.db.Collection(collectionDriverLocation).InsertMany(context.Background(), d)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return err.ServerError("data could not be saved")
	}
	dlr.createIndex("location", "2dsphere")
	return nil
}

func (dlr driverLocationRepo) GetNearDrivers(l core.Location, radius int) ([]*model.DriverInfo, *err.Error) {
	ctx := context.Background()
	fmt.Println("it works")
	fmt.Println(l)
	fmt.Println("it works")

	filter := bson.D{
		{"location",
			bson.D{
				{"$nearSphere", bson.D{
					{"$geometry", l},
					{"$maxDistance", radius},
				}},
			}},
	}

	var drivers []*model.DriverInfo
	cursor, e := dlr.db.Collection(collectionDriverLocation).Find(ctx, filter)

	if e != nil {
		return drivers, err.NotFoundError("no driver found near you")
	}
	e = cursor.All(ctx, &drivers)
	if e != nil {
		return nil, err.NotFoundError("no driver found near you")
	}

	return drivers, nil
}

func (dlr driverLocationRepo) createIndex(field string, i string) *err.Error {
	_, e := dlr.db.Collection(collectionDriverLocation).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{Key: field, Value: bsonx.String(i)}},
	})
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return err.ServerError("index could not be created")
	}
	return nil
}
