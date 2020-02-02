package boat

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// 博客的前后端。从理论上来讲，用户根本不应该操作除了这里规定的api之外的其他api。暴露其他接口仅仅是为了不时之需。

func BindBoatBackend(apiEngine *gin.RouterGroup, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) {
	// post /api/boat/blog
	apiEngine.POST("/blog", func(context *gin.Context) {
		CominBlogTemplate := BlogIn{}
		err := context.BindJSON(&CominBlogTemplate)
		if err != nil {
			context.AbortWithStatus(400)
		}
		hexId := InsertBlog(CominBlogTemplate, boatCollection, ritinCollection)
		context.JSON(200, gin.H{"blog_id": hexId})
	})

	apiEngine.PUT("/blog/:hexId", func(context *gin.Context) {
		blogId := context.Params.ByName("hexId")
		newBlogContent := struct {
			BlogDeltaContent string `json:"blog_delta_content"`
		}{}
		err := context.BindJSON(&newBlogContent)
		if err != nil {
			_ = context.AbortWithError(400, err)
			//context.AbortWithStatus(400)
		} else {
			err := UpdateBlogContent(blogId, newBlogContent.BlogDeltaContent, boatCollection, ritinCollection)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					context.AbortWithStatus(404)
				} else {
					_ = context.AbortWithError(400, err)
					//context.AbortWithStatus(400)
				}
			} else {
				context.Status(200)
				context.Abort()
			}
		}
	})
}
