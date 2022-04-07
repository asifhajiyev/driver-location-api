package db

import (
	"context"
	"driver-location-api/util"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func getConnectionString(dbName string) string {
	username := os.Getenv(fmt.Sprintf("%s_DB_USERNAME", dbName))
	password := os.Getenv(fmt.Sprintf("%s_DB_PASSWORD", dbName))
	host := os.Getenv(fmt.Sprintf("%s_DB_HOST", dbName))

	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		username, password, host, dbName)
}

type MongoRepository struct {
	Client  *mongo.Client
	Db      string
	Timeout time.Duration
}

func newMongoClient(mongoServerURL string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServerURL))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewMongoRepository(dbName string, dbTimeout string) (*MongoRepository, error) {
	timeout := util.StringToInt(dbTimeout)
	cs := getConnectionString(dbName)
	mongoClient, err := newMongoClient(cs, timeout)

	repo := &MongoRepository{
		Client:  mongoClient,
		Db:      dbName,
		Timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	return repo, nil
}

func CloseConnection(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (m MongoRepository) GetMongoDB() *mongo.Database {
	return m.Client.Database(m.Db)
}
