package dbWork

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func FindDataInMongoWithCollectionId(collection *mongo.Collection, hexId string, resultStructure interface{}) error {
	objectId, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return err
	}
	mongoContext := context.Background()
	searchResult := collection.FindOne(mongoContext, bson.M{"_id": objectId})
	err = searchResult.Decode(resultStructure)
	if err != nil {
		return err
	}
	return nil
}

func InsertStructureDataIntoCollection(collection *mongo.Collection, structuredData primitive.M) (string, error) {
	mongoContext := context.Background()
	insertResult, err := collection.InsertOne(mongoContext, structuredData) // ObjectID("5e07e36b5ff8b45d022be1d6")
	if err != nil {
		return "", nil
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), err
}
