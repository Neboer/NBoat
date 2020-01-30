package boat

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// 博客的大前端。从理论上来讲，用户根本不应该操作除了这里规定的api之外的其他api。暴露其他接口仅仅是为了不时之需。

func BindBoatBackend(apiEngine *gin.RouterGroup, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) {
	// post /api/boat/blog
	apiEngine.POST("/blog", func(context *gin.Context) {

	})
}
