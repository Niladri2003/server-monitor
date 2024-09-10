package database

import (
	"context"
	"github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var db *mongo.Database

func GetDbCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDb() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	uri, err := utils.ConnectionURLBuilder("mongodb")
	if err != nil {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	db = client.Database("go_demo")
	return nil
}

func CloseDb() error {
	return db.Client().Disconnect(context.Background())
}
