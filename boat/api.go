// boat有两套api，一套是“输入”标准，一套是“输出”标准，输入标准是新建博客时，前后端向boat插入的数据。输出标准是boat返回给前后端的数据，用来渲染前端页面。
package boat

import (
	"Nboat/ritin"
	quill "github.com/dchenk/go-render-quill"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// 作为博客列表的一项输出。
type BlogInfoItem struct {
	BlogName string
	// 封面图片网址，这个应该在上传博客的时候就已经指定了
	CoverPictureURL string
	// 博客的内容。我们认为博客创建的时间和修改的时间就是博客正文内容改变的时间。
	CreateTime       time.Time
	BlogBriefContent string
}

// 作为插入的一个博文。如果前端想要创建一个新可插入博文，就需要填写以下结构。注意，这个结构不包含时间，对博文分类暂不实现。
type BlogIn struct {
	BlogName         string
	CoverPictureURL  string
	BlogDeltaContent string
}

// 对于每个博文，都请求一次数据库查找对应的ritin内容，这样的操作极大的降低了效率。不过这样虽然带来了额外的开销，但我坚信它是有好处的。
func GetBlogBriefList(boatCollection *mongo.Collection, ritinCollection *mongo.Collection) []BlogInfoItem {
	blogList := getBlogListFromMongoCollection(boatCollection)
	for _, blog := range blogList {
		article, _ := ritin.GetArticle(blog.RelativeRitinID, ritinCollection)
		currentBlogInfo := BlogInfoItem{
			BlogName:         blog.BlogName,
			CoverPictureURL:  blog.CoverPictureURL,
			CreateTime: article.CreateTime,
			BlogBriefContent: quill.Class,
		}
	}
	for index, article := range articleList {
		blogInfoList = append(blogInfoList, BlogInfoItem{
			BlogName:,
			CoverPictureURL: "",
			BlogContent:     "",
		})
	}

	blogInfoList := make([]BlogInfoItem, 0)
	articleList := ritin.GetArticleList(boatCollection)

}

func InsertBlog(blog BlogInsert, collection mongo.Collection) {

}
