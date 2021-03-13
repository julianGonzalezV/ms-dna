package storageconn

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect returns a new connection to storage target
func Connect(addr string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		fmt.Println("Error1!")
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error2!")
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error3!", err)
		log.Fatal(err)
	}
	fmt.Println("Ping!")
	return client
}
