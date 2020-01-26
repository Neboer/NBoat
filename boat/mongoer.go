package boat

import (
	"Nboat/dbWork"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// 这是mongoer模块查询向上层返回的结果的结构。同时也注意，这是mongo数据库中存储内容的真实结构。
type BlogRecord struct {
	ID              primitive.ObjectID `bson:"_id"`
	BlogName        string             `bson:"blogName"`
	CoverPictureURL string             `bson:"coverPictureURL"`

	RelativeRitinID primitive.ObjectID `bson:"articleID"`
}

// 插入一篇博文时，api层面进行ritin的操作，然后传递如下格式给boat的mongoer。
type BlogInRecord struct {
	BlogName        string
	CoverPictureURL string

	RelativeRitinHexID string
}

// 在api层面应该完成插入ritin的操作，然后构建BlogRecord对象。
func insertBlogToMongoCollection(blog BlogInRecord, collection *mongo.Collection) string {
	blogHexID, _ := dbWork.InsertStructureDataIntoCollection(collection, primitive.M{
		"blogName":        blog.BlogName,
		"coverPictureURL": blog.CoverPictureURL,
		"articleID":       primitive.ObjectIDFromHex(blog.RelativeRitinHexID),
	})
	return blogHexID
}

func getBlogListFromMongoCollection(collection *mongo.Collection) []BlogRecord {
	BlogRecordList := make([]BlogRecord, 0)
	_ = dbWork.GetDocumentList(collection, &BlogRecordList)
	return BlogRecordList
}

func getBlogFromMongoCollection(blogID string, collection *mongo.Collection) (BlogRecord, error) {
	result := BlogRecord{}
	err := dbWork.FindDataInMongoWithCollectionId(collection, blogID, &result)
	// err == mongo.ErrNoDocuments or mongo.
	return result, err
}

// 你更新的是blog还是article？这一点十分重要，正常情况下，blog本身是不应该做改动的。但是也可以改。前端收到“修改”的请求的时候，应该直接查询articleID
// 然后就修改article就可以了，没必要操作blog本身。
func updateBlogFromMongoCollection(blogID string, record BlogInRecord, collection *mongo.Collection) {
	_ = dbWork.UpdateStructureDataFromCollection(collection, blogID, primitive.M{
		"blogName":        record.BlogName,
		"coverPictureURL": record.CoverPictureURL,
		"articleID":       record.RelativeRitinHexID,
	})
}
