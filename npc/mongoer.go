package npc

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
	"time"
)

func InsertPictureIntoMongoCollection(pictureContent []byte, pictureCollection *mongo.Collection) string {
	mongoContext := context.Background()
	insertResult, _ := pictureCollection.InsertOne(mongoContext, bson.M{"content": pictureContent}) // ObjectID("5e07e36b5ff8b45d022be1d6")
	return insertResult.InsertedID.(primitive.ObjectID).Hex()
}

func GetPictureContentFromMongoCollection(pictureIdHexString string, pictureCollection *mongo.Collection) []byte {
	objectId, _ := primitive.ObjectIDFromHex(pictureIdHexString)

	type pieceOfPictureStructure struct {
		ID      primitive.ObjectID
		Content primitive.Binary
	}

	mongoContext, _ := context.WithTimeout(context.Background(), 2*time.Second)
	searchResult := pictureCollection.FindOne(mongoContext, bson.M{"_id": objectId})
	result := pieceOfPictureStructure{}

	err := searchResult.Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result.Content.Data
}
