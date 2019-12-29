package ritin

import (
	"Nboat/dbWork"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ArticleRecord struct {
	ID           primitive.ObjectID `json:"_id"`
	Content      string             `json:"content"`
	CreateTime   primitive.DateTime `json:"create_time"`
	LastModified primitive.DateTime `json:"last_modified"`
}

func insertArticleDeltaIntoMongoCollection(delta string, ritinMongoCollection *mongo.Collection) string {
	// delta本应是一个json，但传输和存储都按照string来看待。
	hexId, _ := dbWork.InsertStructureDataIntoCollection(ritinMongoCollection, bson.M{"content": delta, "createTime": time.Now(), "lastModified": time.Now()})
	return hexId
}

func getArticleDeltaFromMongoCollection(articleIdHexString string, ritinMongoCollection *mongo.Collection) ArticleRecord {
	result := ArticleRecord{}
	_ = dbWork.FindDataInMongoWithCollectionId(ritinMongoCollection, articleIdHexString, &result)
	return result
}
