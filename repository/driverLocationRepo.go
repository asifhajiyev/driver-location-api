package repository

import (
	"context"
	"driver-location-api/db"
	err "driver-location-api/error"
	"driver-location-api/model/core"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const collectionName = "driver_locations"

type DriverLocationRepo interface {
	SaveDriverLocation(dl core.DriverLocation) (*core.DriverLocation, *err.Error)
	SaveDriverLocationFile(dl []core.DriverLocation) *err.Error
	GetNearDriversLocation(g core.Geometry, radius int) (*[]core.DriverLocation, *err.Error)
	createIndex(field string, i string) *err.Error
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
	dlr.createIndex("location", "2dsphere")
	return nil
}

func (dlr driverLocationRepo) GetNearDriversLocation(g core.Geometry, radius int) (*[]core.DriverLocation, *err.Error) {
	ctx := context.Background()
	//point := core.NewPoint(-73.9667, 40.78)

	filter := bson.D{
		{"location",
			bson.D{
				{"$nearSphere", bson.D{
					{"$geometry", g},
					{"$maxDistance", radius},
				}},
			}},
	}

	var drivers []core.DriverLocation
	cursor, e := dlr.c.Find(ctx, filter)
	fmt.Println("cursor", cursor)

	if e != nil {
		return &drivers, err.NotFoundError("no driver found near you")
	}
	e = cursor.All(ctx, &drivers)

	if e != nil {
		return nil, err.NotFoundError("no driver found near you")
	}

	return &drivers, nil
}

func (dlr driverLocationRepo) createIndex(field string, i string) *err.Error {
	_, e := dlr.c.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{Key: field, Value: bsonx.String(i)}},
	})
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return err.ServerError("index could not be created")
	}
	return nil
}
