// ritin 的web api。通过调用api.go的函数，做基于mongoer.go的抽象，mongoer.go再调用dbwork库进行操作。
package ritin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func BindRitin(engine *gin.RouterGroup, ritinCollection *mongo.Collection) {
	mainGroup := engine.Group("/ritin")
	// 添加delta。
	mainGroup.POST("/article", func(context *gin.Context) {
		upload := struct {
			Content string `json:"content"`
		}{}
		_ = context.BindJSON(&upload)
		hexId := InsertArticle(upload.Content, ritinCollection)
		context.JSON(http.StatusOK, gin.H{"articleId": hexId})
	})

	mainGroup.GET("/article", func(context *gin.Context) {
		articleList := GetArticleList(ritinCollection)
		context.JSON(200, articleList)
	})

	mainGroup.GET("/article/:hexId", func(context *gin.Context) {
		queryId := context.Param("hexId")
		article, err := GetArticle(queryId, ritinCollection)
		if err != nil {
			handleErr(context, err)
		} else {
			context.JSON(200, article)
		}
		//articleRecord := getArticleFromMongoCollection(queryId, ritinCollection)
		//deltaContent := articleRecord.Content
		//renderedHtml, _ := quill.Render([]byte(deltaContent))
		//context.Data(http.StatusOK, "text/html; charset=utf-8", renderedHtml)
	})

	mainGroup.GET("/edit/:hexId", func(context *gin.Context) {
		queryId := context.Param("hexId")
		articleRecord, _ := getArticleFromMongoCollection(queryId, ritinCollection)
		deltaContent := articleRecord.Content
		context.JSON(http.StatusOK, gin.H{"delta": deltaContent})
	})

	mainGroup.PUT("/article/:hexId", func(context *gin.Context) {
		queryId := context.Param("hexId")
		upload := struct {
			Content string `json:"content"`
		}{}
		_ = context.BindJSON(&upload)
		UpdateArticle(queryId, ritinCollection, upload.Content)
		_ = context.BindJSON(&upload)
		updateArticleFromMongoCollection(upload.Content, queryId, ritinCollection)
		context.JSON(200, gin.H{"result": "success."})
	})
}

func handleErr(ctx *gin.Context, err error) {
	if err == mongo.ErrNoDocuments {
		ctx.String(404, "no such article")
	} else {
		_ = ctx.AbortWithError(400, err)
	}
}
