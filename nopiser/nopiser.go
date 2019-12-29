package nopiser

import (
	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func BindNopiser(engine *gin.RouterGroup, database *mongo.Database) {
	mainGroup := engine.Group("/nopiser")
	pictureCollection := database.Collection("nopiser")
	mainGroup.POST("/picture", func(context *gin.Context) {
		uploadImageHandler(context, pictureCollection)
	})
	mainGroup.GET("/picture/:objectHex", func(context *gin.Context) {
		getImageHandler(context, pictureCollection)
	})
}

func uploadImageHandler(ctx *gin.Context, nopiserMongoCollection *mongo.Collection) {
	header, _ := ctx.FormFile("image")
	imageContent := make([]byte, header.Size)
	file, _ := header.Open()
	_, _ = file.Read(imageContent)

	imageType, _ := filetype.Match(imageContent)
	imageId := InsertPictureIntoMongoCollection(imageContent, imageType.MIME.Value, nopiserMongoCollection)
	ctx.JSON(http.StatusOK, gin.H{"url": "/api/nopiser/picture/" + imageId})
}

func getImageHandler(ctx *gin.Context, nopiserMongoCollection *mongo.Collection) {
	objectHex := ctx.Param("objectHex")
	pictureRecord := GetPictureRecordFromMongoCollection(objectHex, nopiserMongoCollection)
	ctx.Data(http.StatusOK, pictureRecord.MIME, pictureRecord.PictureContent.Data)
}
