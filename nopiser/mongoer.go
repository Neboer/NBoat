package nopiser

import (
	"Nboat/dbWork"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PictureRecord struct {
	ID             primitive.ObjectID `bson:"_id"`
	MIME           string             `bson:"MIME"`
	PictureContent primitive.Binary   `bson:"PictureContent"`
}

func InsertPictureIntoMongoCollection(pictureContent []byte, pictureMIMEValue string, pictureCollection *mongo.Collection) string {
	hexId, _ := dbWork.InsertStructureDataIntoCollection(pictureCollection, bson.M{"PictureContent": pictureContent, "MIME": pictureMIMEValue})
	return hexId
}

func GetPictureRecordFromMongoCollection(pictureIdHexString string, pictureCollection *mongo.Collection) PictureRecord {
	result := PictureRecord{}
	_ = dbWork.FindDataInMongoWithCollectionId(pictureCollection, pictureIdHexString, &result)
	return result
}
