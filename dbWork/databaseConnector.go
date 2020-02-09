package dbWork

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"log"
	"time"
)

func ConnectionInit(MongoLocationString string, Username string, Password string) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+MongoLocationString).SetAuth(options.Credential{
		AuthSource: "admin",
		Username:   Username,
		Password:   Password,
	}))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("nboat")
}

// 按顺序返回三个collection。分别为nboat ritin和nopiser
func GetCollection(database mongo.Database) (*mongo.Collection, *mongo.Collection, *mongo.Collection) {
	return database.Collection("nboat"), database.Collection("ritin"), database.Collection("nopiser")
}
