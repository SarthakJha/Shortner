package shortner

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConn struct
type MongoConn struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// SeedDb initialises and connects the db
func SeedDb() (MongoConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		cancel()
		return MongoConn{}, err
	}
	err1 := client.Connect(ctx)
	if err1 != nil {
		cancel()
		return MongoConn{}, err
	}
	db := client.Database("shawty")
	cancel()
	return MongoConn{
		Client: client,
		Db:     db,
	}, nil

}
