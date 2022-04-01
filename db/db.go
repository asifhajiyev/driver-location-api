package db

import (
	"context"
	"driver-location-api/util"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	fmt.Println("cs", cs)
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

/*func connectToDB() *mongo.Database {
	url := getConnectionString()
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(os.Getenv("DB_NAME"))
	return db
}*/

/*func closeConnection(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}*/

/*func (m MongoRepository) GetCollection(c string) *mongo.Collection {
	return m.Client.Database(m.Db).Collection(c)
}*/

func (m MongoRepository) GetMongoDB() *mongo.Database {
	return m.Client.Database(m.Db)
}
