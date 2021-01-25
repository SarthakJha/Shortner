package shortner

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		log.Fatalln(err)
	}

	err1 := client.Connect(ctx)
	if err1 != nil {
		log.Fatalln(err1)
	}

	// to check if connection was successfull
	err9 := client.Ping(ctx, readpref.Primary())
	if err9 != nil {
		log.Fatalln(err9)
	}
	db := client.Database(os.Getenv("DB_NAME"))

	defer cancel()
	fmt.Println("Connected to mongo...")
	return MongoConn{
		Client: client,
		Db:     db,
	}, nil

}
