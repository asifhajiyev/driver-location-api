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
	SaveDriverLocation(di core.DriverInfo) (*core.DriverInfo, *err.Error)
	SaveDriverLocationFile(di []core.DriverInfo) *err.Error
	GetNearDriversInfo(g core.Location, radius int) ([]*core.DriverInfo, *err.Error)
	createIndex(field string, i string) *err.Error
}
type driverLocationRepo struct {
	c *mongo.Collection
}

func NewDriverLocationRepo(m *db.MongoRepository) DriverLocationRepo {
	return &driverLocationRepo{c: m.GetCollection(collectionName)}
}

func (dlr driverLocationRepo) SaveDriverLocation(di core.DriverInfo) (*core.DriverInfo, *err.Error) {
	fmt.Println(di)
	_, e := dlr.c.InsertOne(context.Background(), di)
	if e != nil {
		log.Errorf("SaveDriverLocation.error: %v", e)
		return nil, err.ServerError("data could not be saved")
	}
	return &di, nil
}

func (dlr driverLocationRepo) SaveDriverLocationFile(di []core.DriverInfo) *err.Error {

	var d []interface{}
	for _, t := range di {
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

func (dlr driverLocationRepo) GetNearDriversInfo(l core.Location, radius int) ([]*core.DriverInfo, *err.Error) {
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

	var drivers []*core.DriverInfo
	cursor, e := dlr.c.Find(ctx, filter)

	if e != nil {
		return drivers, err.NotFoundError("no driver found near you")
	}
	e = cursor.All(ctx, &drivers)
	/*for cursor.Next(ctx) {
		fmt.Println("cursor raw", cursor.Current)
		var d *core.DriverInfo
		e = cursor.Decode(&d)
		if e != nil {
			return nil, err.NotFoundError("can not fetch drivers")
		}
		drivers = append(drivers, d)
	}*/

	if e != nil {
		return nil, err.NotFoundError("no driver found near you")
	}

	return drivers, nil
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
