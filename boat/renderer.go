// 渲染html页面
package boat

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// 这里负责渲染博客的各个页面。这里直接绑定到服务器路径上。这是服务器的真正主页。
func BindBoatRenderer(engine *gin.RouterGroup, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) {
	// 用户请求主页。主页上有一些博文。
	engine.GET("/", func(context *gin.Context) {
		BlogList := GetBlogBriefList(boatCollection, ritinCollection)

		//content := make([]map[string]string, 0)
		//content1 := map[string]string{"Title": "this is title", "Time": "Aug 28", "Brief": "this is a sample brief of the whole article."}
		//content2 := map[string]string{"Title": "this second title", "Time": "Aug 29", "Brief": "brief of the whole article."}
		//content = append(content, content1, content2)
		r := render.New(render.Options{Directory: "front", Layout: "layout", RequirePartials: true})
		_ = r.HTML(context.Writer, http.StatusOK, "home", BlogList)
	})

	engine.GET("/blog/:HexId", func(context *gin.Context) {
		hexId := context.Params.ByName("HexId")
		context.JSON(http.StatusOK, gin.H{"articleId": hexId})
	})

	engine.POST("/blog", func(context *gin.Context) {
		upload := struct {
			Content string `json:"content"`
		}{}
		_ = context.BindJSON(&upload)
	})

	engine.GET("/editor", func(context *gin.Context) {
		r := render.New(render.Options{Directory: "front", Layout: "layout", RequirePartials: true})
		_ = r.HTML(context.Writer, http.StatusOK, "editor", "")
	})
}
