// ritin的IO规范，是boat和ritin模块交互的方法。
package ritin

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UploadedArticle struct {
	Content string `json:"content"`
}

// 完整的一篇文章
type Article struct {
	Id             string
	CreateTime     time.Time
	LastModifyTime time.Time
	Content        string
}

func InsertArticle(article Article, collection *mongo.Collection) string {
	hexId := insertArticleDeltaIntoMongoCollection(article.Content, collection)
	return hexId
}

func GetArticle(articleHexID string, collection *mongo.Collection) Article {
	articleRecord := getArticleFromMongoCollection(articleHexID, collection)
	return Article{
		Id:             articleRecord.ID.Hex(),
		CreateTime:     articleRecord.CreateTime.Time(),
		LastModifyTime: articleRecord.LastModified.Time(),
		Content:        articleRecord.Content,
	}
}

func UpdateArticle(articleHexID string, collection *mongo.Collection, newArticle Article) {

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
