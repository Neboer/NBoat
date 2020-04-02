package boat

import (
	"Nboat/cookieauth"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/mongo"
)

type FrontendSettings struct {
	CountOfBlogSubjectShowInPerPage int
}

func BindBoatFrontend(apiEngine *gin.RouterGroup, collection *mongo.Collection, settings FrontendSettings) {
	r := render.New(render.Options{Directory: "front", Layout: "layout"})
	apiEngine.GET("/", func(context *gin.Context) {
		_ = r.HTML(context.Writer, 200, "home", struct{ IsAuthed bool }{IsAuthed: cookieauth.IsAuthed(context)})
	})
	// “最新的”博客列表，每页展示条数可以在main中设定。
	apiEngine.GET("/newest", func(context *gin.Context) {
		pageIndex := parseQueryIndex(context, settings.CountOfBlogSubjectShowInPerPage)
		wholeBlogBriefList, _ := GetBlogSubjectList(collection)
		blogListToShow := BlogSubjectBriefList{}
		totalPageCount := len(wholeBlogBriefList)/settings.CountOfBlogSubjectShowInPerPage + 1
		if pageIndex > totalPageCount {
			// 返回错误页面
		} else if pageIndex == totalPageCount {
			// 用户请求了最后一页
			blogListToShow = wholeBlogBriefList[(pageIndex-1)*settings.CountOfBlogSubjectShowInPerPage:]
		} else {
			blogListToShow = wholeBlogBriefList[(pageIndex-1)*settings.CountOfBlogSubjectShowInPerPage : pageIndex*settings.CountOfBlogSubjectShowInPerPage]
		}
		newestPage := struct {
			IsAuthed       bool
			TotalPageCount int
			CurrentPage    int
			BlogBriefList  BlogSubjectBriefList
		}{
			IsAuthed:       true,
			TotalPageCount: totalPageCount,
			CurrentPage:    pageIndex,
			BlogBriefList:  blogListToShow,
		}
		_ = r.HTML(context.Writer, 200, "newest", newestPage)
	})

	apiEngine.GET("/sort", func(context *gin.Context) {
		_ = r.HTML(context.Writer, 200, "sort", nil)
	})

	apiEngine.GET("/newBlog", func(ctx *gin.Context) {

	})

	apiEngine.GET("/editor.js", func(context *gin.Context) {
		context.Render()
	})
}
