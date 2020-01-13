package boat

import (
	"Nboat/dbWork"
	"Nboat/ritin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlogList() []BlogInfo {
	BlogInfoList := make([]BlogInfo, 0)

}

/* 上传博客。这里的BlogInfo是一个完整的博客。
上传时，将博客的所有内容都填满，除了id之外。
*/

func InsertBlogIntoDatabase(collection *mongo.Collection, info BlogInfo) error {

	InsertStruct := primitive.M{"BlogName": info.BlogName, "CoverPictureURL": info.CoverPictureURL, "RitinId": info.BlogContentRitinInfo.ID}
	_, _ = dbWork.InsertStructureDataIntoCollection(collection)
}
