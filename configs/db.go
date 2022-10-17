package configs
import (
	"context"
	"fmt"
	"log"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection() *mongo.Client{
	URL_DB := os.Getenv("MONGO_URL")
	client, err:= mongo.NewClient(options.Client().ApplyURI(URL_DB))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo DB")
	return client
}