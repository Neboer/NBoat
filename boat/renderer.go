// 渲染html页面
package boat

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"net/http"
)

// 这里负责渲染博客的各个页面。这里直接绑定到服务器路径上。这是服务器的真正主页。
func BindBoatRenderer(engine *gin.RouterGroup, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) {
	r := render.New(render.Options{Directory: "front", Layout: "layout", RequirePartials: true})
	// 用户请求“最新的”，这个上有一个博文总列表。
	engine.GET("/newest", func(context *gin.Context) {
		BlogList := GetBlogBriefList(boatCollection, ritinCollection)
		_ = r.HTML(context.Writer, http.StatusOK, "newest", BlogList)
	})

	engine.GET("/", func(context *gin.Context) {
		_ = r.HTML(context.Writer, 200, "home", nil)
	})

	// /blog/:hexid 返回渲染完毕的博客内容。
	engine.GET("/blog/:HexId", func(context *gin.Context) {
		hexId := context.Params.ByName("HexId")
		//	if hexId == "test" {
		//		delta := `[{"insert":"This "},{"attributes":{"italic":true},"insert":"is"},
		//{"insert":" "},{"attributes":{"bold":true},"insert":"great!"},{"insert":"\n"}]`
		//		html, _ := quill.Render([]byte(delta))
		//		renderedBlogHTML := template.HTML(string(html))
		//		_ = r.HTML(context.Writer, http.StatusOK, "blog", renderedBlogHTML)
		//		return
		//	}
		BlogOutput, err := GetBlog(hexId, boatCollection, ritinCollection)
		if err == mongo.ErrNoDocuments {
			context.String(404, "404 Blog Not Found")
			context.Abort()
		} else if err != nil {
			context.String(400, "400 Bad Request")
			context.Abort()
		} else {
			_ = r.HTML(context.Writer, http.StatusOK, "blog", BlogOutput)
		}
	})

	// 如果渲染编辑器的话，需要生成一个编辑器页面EditorPage对象。这里会返回某个特定博客的编辑界面。
	engine.GET("/editor/:HexId", func(context *gin.Context) {
		blogId := context.Params.ByName("HexId")
		if blogId == "" {
			context.AbortWithStatus(400)
		}
		blogEditObject, err := GetBlogDelta(blogId, boatCollection, ritinCollection)
		if err == mongo.ErrNoDocuments {
			context.AbortWithStatus(404)
		} else if err != nil {
			context.AbortWithStatus(400)
		} else {
			_ = r.HTML(context.Writer, 200, "existBlogEditor",
				struct {
					BlogDeltaContent template.JS
					BlogId           string
					BlogName         string
				}{
					template.JS(blogEditObject.BlogDeltaContent),
					blogId,
					blogEditObject.BlogName,
				})
		}
	})

	// 返回通用编辑页面，创建博客的页面。
	engine.GET("/editor", func(context *gin.Context) {
		_ = r.HTML(context.Writer, http.StatusOK, "newBlogEditor", nil)
	})

	engine.GET("/sort", func(context *gin.Context) {
		_ = r.HTML(context.Writer, 200, "sort", nil)
	})
}
