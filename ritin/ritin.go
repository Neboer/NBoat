package ritin

import (
	quill "github.com/dchenk/go-render-quill"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func BindRitin(engine *gin.RouterGroup, database *mongo.Database) {
	ritinCollection := database.Collection("ritin")
	mainGroup := engine.Group("/ritin")
	// 添加delta。
	mainGroup.POST("/article", func(context *gin.Context) {
		upload := CreateNewArticleUpload{}
		_ = context.BindJSON(&upload)
		hexId := insertArticleDeltaIntoMongoCollection(upload.Content, ritinCollection)
		context.JSON(http.StatusOK, gin.H{"articleId": hexId})
	})

	mainGroup.GET("/article/:hexId", func(context *gin.Context) {
		queryId := context.Param("hexId")
		articleRecord := getArticleDeltaFromMongoCollection(queryId, ritinCollection)
		deltaContent := articleRecord.Content
		renderedHtml, _ := quill.Render([]byte(deltaContent))
		context.Data(http.StatusOK, "text/html; charset=utf-8", renderedHtml)
	})
}
