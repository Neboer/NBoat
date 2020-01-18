package ritin

import (
	"Nboat/dbWork"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ArticleRecord struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Content      string             `bson:"content"`
	CreateTime   primitive.DateTime `bson:"createTime"`
	LastModified primitive.DateTime `bson:"lastModified"`
}

func insertArticleDeltaIntoMongoCollection(delta string, ritinMongoCollection *mongo.Collection) string {
	// delta本应是一个json，但传输和存储都按照string来看待。
	hexId, _ := dbWork.InsertStructureDataIntoCollection(ritinMongoCollection, bson.M{"content": delta, "createTime": time.Now(), "lastModified": time.Now()})
	return hexId
}

func getArticleFromMongoCollection(articleIdHexString string, ritinMongoCollection *mongo.Collection) ArticleRecord {
	result := ArticleRecord{}
	_ = dbWork.FindDataInMongoWithCollectionId(ritinMongoCollection, articleIdHexString, &result)
	return result
}

func updateArticleFromMongoCollection(newDelta string, articleIdHexString string, ritinMongoCollection *mongo.Collection) {
	_ = dbWork.UpdateStructureDataFromCollection(ritinMongoCollection, articleIdHexString, bson.M{"content": newDelta, "lastModified": time.Now()})
}

func getArticleListFromMongoCollection(ritinMongoCollection *mongo.Collection) []ArticleRecord {
	result := make([]ArticleRecord, 0)
	_ = dbWork.GetDocumentList(ritinMongoCollection, &result)
	return result
}
