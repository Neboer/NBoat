package dbWork

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNPCCollection(mongoDatabaseConnection *mongo.Database) *mongo.Collection {
	return mongoDatabaseConnection.Collection("npc")
}
