// 渲染html页面
package boat

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"net/http"
)

func BindBoatRenderer(engine *gin.RouterGroup) {
	engine.GET("/", func(context *gin.Context) {
		content := make([]map[string]string, 0)
		content1 := map[string]string{"Title": "this is title", "Time": "Aug 28", "Brief": "this is a sample brief of the whole article."}
		content2 := map[string]string{"Title": "this second title", "Time": "Aug 29", "Brief": "brief of the whole article."}
		content = append(content, content1, content2)
		r := render.New(render.Options{Directory: "front", Layout: "layout", RequirePartials: true})
		_ = r.HTML(context.Writer, http.StatusOK, "home", content)
	})

	engine.GET("/editor", func(context *gin.Context) {
		r := render.New(render.Options{Directory: "front", Layout: "layout", RequirePartials: true})
		_ = r.HTML(context.Writer, http.StatusOK, "editor", "")
	})
}
