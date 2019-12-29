package npc

import (
	"Nboat/dbWork"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func BindNPC(engine *gin.RouterGroup, database *mongo.Database) {
	mainGroup := engine.Group("/npc")
	pictureCollection := dbWork.GetNPCCollection(database)
	mainGroup.POST("/picture", func(context *gin.Context) {
		uploadImageHandler(context, pictureCollection)
	})
	mainGroup.GET("/picture/:objectHex", func(context *gin.Context) {
		getImageHandler(context, pictureCollection)
	})
}

func uploadImageHandler(ctx *gin.Context, npcMongoCollection *mongo.Collection) {
	header, _ := ctx.FormFile("image")
	imageContent := make([]byte, header.Size)
	file, _ := header.Open()
	_, _ = file.Read(imageContent)
	imageId := InsertPictureIntoMongoCollection(imageContent, npcMongoCollection)
	ctx.JSON(http.StatusOK, gin.H{"url": "/api/npc/picture/" + imageId})
}

func getImageHandler(ctx *gin.Context, npcMongoCollection *mongo.Collection) {
	objectHex := ctx.Param("objectHex")
	pictureContent := GetPictureContentFromMongoCollection(objectHex, npcMongoCollection)
	ctx.Data(http.StatusOK, "image/png", pictureContent)
}

//server.GET("/api/npc/picture/:objectHex", func(ctx *gin.Context) {
//	objectHex := ctx.Param("objectHex")
//
//	objectId, _ := primitive.ObjectIDFromHex(objectHex)
//
//	type Language struct {
//		ID      primitive.ObjectID
//		Content primitive.Binary
//	}
//
//	mongoContext, _ := context.WithTimeout(context.Background(), 2*time.Second)
//	searchResult := pictureCollection.FindOne(mongoContext, bson.M{"_id": objectId})
//	station := Language{}
//
//	err = searchResult.Decode(&station)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ctx.Data(http.StatusOK, "image/png", station.Content.Data)
//})
//
//server.POST("/api/npc/picture", func(ctx *gin.Context) {
//	header, _ := ctx.FormFile("image")
//	imageContent := make([]byte, header.Size)
//	file, _ := header.Open()
//	_, _ = file.Read(imageContent)

//})
