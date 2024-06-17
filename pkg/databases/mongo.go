package databases

import (
	"context"
	"log"
	"time"

	"github.com/ferhateroglu/go-fiber/internal/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDatabase struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoDatabase(cfg *configs.Config) (*MongoDatabase, error) {
	clientOptions := options.Client().ApplyURI(cfg.Database.MongoURI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Printf("Cannot connect to MongoDB: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Cannot ping to MongoDB: %v", err)
		client.Disconnect(ctx)
		return nil, err
	}

	db := client.Database(cfg.Database.MongoDBName)

	return &MongoDatabase{
		Client: client,
		DB:     db,
	}, nil
}

func (m *MongoDatabase) GetDatabase() *mongo.Database {
	return m.DB
}

func (m *MongoDatabase) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
