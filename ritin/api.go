// ritin的IO规范，是boat和ritin模块交互的方法。
package ritin

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// 完整的一篇文章
type Article struct {
	Id             string
	CreateTime     time.Time
	LastModifyTime time.Time
	Content        string
}

// 外部调用可以向内部写入的接口
// 别装了，article本质就是string。外部传入内部的只能是string。

// 在这个api层面，对顶层暴露尽可能简单的接口，但对底层严苛。
func InsertArticle(article string, collection *mongo.Collection) string {
	hexId := insertArticleDeltaIntoMongoCollection(article, collection)
	return hexId
}

func GetArticle(articleHexID string, collection *mongo.Collection) (Article, error) {
	articleRecord, err := getArticleFromMongoCollection(articleHexID, collection)
	resultArticle := Article{
		Id:             articleRecord.ID.Hex(),
		CreateTime:     articleRecord.CreateTime.Time(),
		LastModifyTime: articleRecord.LastModified.Time(),
		Content:        articleRecord.Content,
	}
	return resultArticle, err
}

func UpdateArticle(articleHexID string, collection *mongo.Collection, newArticle string) {
	updateArticleFromMongoCollection(newArticle, articleHexID, collection)
}

func GetArticleList(collection *mongo.Collection) []Article {
	ArticleRecordList := getArticleListFromMongoCollection(collection)
	ArticleList := make([]Article, 0)
	for singleArticleRecordIndex := range ArticleRecordList {
		currentItem := ArticleRecordList[singleArticleRecordIndex]
		ArticleList = append(ArticleList, Article{
			Id:             currentItem.ID.Hex(),
			CreateTime:     currentItem.CreateTime.Time(),
			LastModifyTime: currentItem.LastModified.Time(),
			Content:        currentItem.Content,
		})
	}
	return ArticleList
}
