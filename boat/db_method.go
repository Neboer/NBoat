// 就是专门保存数据库方法的地方。涉及到什么增删查改都到这里来找吧！
package boat

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"time"
)

// 内部方法，严禁外泄
func getBlogSubject(collection *mongo.Collection, id primitive.ObjectID) (BlogSubject, error) {
	resultBlogSubject := BlogSubject{}
	mongoContext := context.Background()
	searchResult := collection.FindOne(mongoContext, bson.M{"_id": id})
	//fmt.Println(searchResult.err.Error())
	err := searchResult.Decode(&resultBlogSubject)
	if err != nil {
		return resultBlogSubject, err // mongo.ErrNoDocuments
	}
	return resultBlogSubject, nil
}

func GetBlogSubject(collection *mongo.Collection, HexID string) (BlogSubject, error) {
	objectId, err := primitive.ObjectIDFromHex(HexID)
	if err != nil {
		return BlogSubject{}, err
	}
	return getBlogSubject(collection, objectId)
}

func GetBlogSubjectList(collection *mongo.Collection) (BlogSubjectBriefList, error) {
	resultList := BlogSubjectBriefList{}
	mongoContext := context.Background()
	resultCursor, _ := collection.Find(mongoContext, primitive.M{})
	err := resultCursor.All(mongoContext, &resultList)
	return resultList, err
}

//
func CreateEmptyBlogSubject(collection *mongo.Collection, info BlogSubjectInfo) (string, error) {
	mongoContext := context.Background()
	insertResult, err := collection.InsertOne(mongoContext, info.toBlogSubject()) // ObjectID("5e07e36b5ff8b45d022be1d6")
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), err
}

// 注意，插入是按照最大+1的原则，blog的index本身并没有参考价值！
func InsertArticle(collection *mongo.Collection, inputArticle ArticleInput, toBlogSubjectID string) (int, error) {
	mongoContext := context.Background()
	objectId, err := primitive.ObjectIDFromHex(toBlogSubjectID)
	if err != nil {
		return 0, err
	}
	blogSubjectToUpdate, err := getBlogSubject(collection, objectId)
	if err != nil {
		return 0, err
	}
	maxValue := 0
	for _, article := range blogSubjectToUpdate.Article {
		if article.Index > maxValue {
			maxValue = article.Index
		}
	}
	newArticle := inputArticle.toArticle(maxValue + 1)
	_, err = collection.UpdateOne(mongoContext, bson.M{"_id": objectId}, bson.M{"$push": bson.M{"article": newArticle}})
	if err != nil {
		return 0, err
	}
	return maxValue + 1, nil
}

// article被删去也不会改变index
func DeleteArticle(collection *mongo.Collection, articleIndex int, fromBlogSubjectID string) error {
	objectId, err := primitive.ObjectIDFromHex(fromBlogSubjectID)
	if err != nil {
		return err
	}
	mongoContext := context.Background()
	deleteResult, err := collection.UpdateOne(mongoContext, bson.M{"_id": objectId}, bson.M{"$pull": bson.M{"article": bson.M{"index": articleIndex}}})
	if err != nil {
		return err
	}
	if deleteResult.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	} else if deleteResult.ModifiedCount == 0 {
		// 要删除的index根本就不存在
		return mongo.ErrInvalidIndexValue
	} else {
		return nil
	}
}

// Update article
func UpdateArticleInfo(collection *mongo.Collection, articleIndex int, fromBlogSubjectID string, newInfo ArticleInfo) error {
	objectId, err := primitive.ObjectIDFromHex(fromBlogSubjectID)
	if err != nil {
		return err
	}
	mongoContext := context.Background()
	_, err = getBlogSubject(collection, objectId)
	if err != nil {
		return err
	}
	updateResult, err := collection.UpdateOne(mongoContext, bson.M{"_id": objectId, "article.index": articleIndex},
		bson.M{"$set": bson.M{"article.$.info": newInfo, "article.$.meta.last_modify_time": time.Now()}})
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return mongo.ErrInvalidIndexValue
	} else {
		return nil
	}
}

func UpdateArticleContent(collection *mongo.Collection, articleIndex int, fromBlogSubjectID string, newContent string) error {
	objectId, err := primitive.ObjectIDFromHex(fromBlogSubjectID)
	if err != nil {
		return err
	}
	mongoContext := context.Background()
	_, err = getBlogSubject(collection, objectId)
	if err != nil {
		return err
	}
	updateResult, err := collection.UpdateOne(mongoContext, bson.M{"_id": objectId, "article.index": articleIndex},
		bson.M{"$set": bson.M{"article.$.content": newContent, "article.$.meta.last_modify_time": time.Now()}})
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return mongo.ErrInvalidIndexValue
	} else {
		return nil
	}
}
