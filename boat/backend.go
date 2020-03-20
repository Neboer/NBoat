package boat

import (
	"Nboat/cookieauth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

func BindBoatBackend(apiEngine *gin.RouterGroup, mongoCollection *mongo.Collection) {
	apiEngine.Use(cookieauth.OnlyAllowAuthor("/api/blog"))
	// 创建一个空的博客项目
	apiEngine.POST("/blog", func(context *gin.Context) {
		info := BlogSubjectInfo{}
		err := context.BindJSON(&info)
		if err != nil || info.Title == "" {
			_ = context.AbortWithError(400, err)
		}
		newBlogSubjectID, err := CreateEmptyBlogSubject(mongoCollection, info)
		if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.JSON(200, gin.H{"blog_subject_id": newBlogSubjectID})
		}
	})
	// 向博客项目中插入新的文章
	apiEngine.POST("/blog/:blogID", func(context *gin.Context) {
		article := ArticleInput{}
		err := context.BindJSON(&article)
		if err != nil || article.Content == "" || article.Info.Name == "" {
			_ = context.AbortWithError(400, err)
		}
		newArticleIndex, err := InsertArticle(mongoCollection, article, context.Param("blogID"))
		if err == mongo.ErrNoDocuments || err == mongo.ErrInvalidIndexValue {
			_ = context.AbortWithError(404, err)
		} else if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.JSON(200, gin.H{"article_id": newArticleIndex})
		}
	})
	// 更新博客的info
	apiEngine.PUT("/blog/:blogID", func(context *gin.Context) {
		newBlogSubjectInfo := BlogSubjectInfo{}
		err := context.BindJSON(&newBlogSubjectInfo)
		if err != nil || newBlogSubjectInfo.Title == "" {
			_ = context.AbortWithError(400, err)
		}
		err = UpdateBlogSubjectInfo(mongoCollection, newBlogSubjectInfo, context.Param("blogID"))
		if err == mongo.ErrNoDocuments || err == mongo.ErrInvalidIndexValue {
			context.AbortWithStatus(404)
		} else if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.Status(200)
			context.Done()
		}
	})
	// 更新文章的info
	apiEngine.PUT("/blog/:blogID/:articleID/info", func(context *gin.Context) {
		articleInfo := ArticleInfo{}
		err := context.BindJSON(&articleInfo)
		if err != nil || articleInfo.Name == "" {
			_ = context.AbortWithError(400, err)
		}
		articleIndex := context.GetInt(":articleID")
		err = UpdateArticleInfo(mongoCollection, articleIndex, context.Param("blogID"), articleInfo)
		if err == mongo.ErrNoDocuments || err == mongo.ErrInvalidIndexValue {
			_ = context.AbortWithError(404, err)
		} else if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.Status(200)
			context.Done()
		}
	})
	// 更新文章的content
	apiEngine.PUT("/blog/:blogID/:articleID/content", func(context *gin.Context) {
		articleInfo := ArticleInfo{}
		err := context.BindJSON(&articleInfo)
		if err != nil || articleInfo.Name == "" {
			_ = context.AbortWithError(400, err)
		}
		articleIndex, err := strconv.Atoi(context.Param("articleID"))
		if err != nil {
			_ = context.AbortWithError(400, err)
		}
		byteContent, err := context.GetRawData()
		if err != nil {
			_ = context.AbortWithError(400, err)
		}
		err = UpdateArticleContent(mongoCollection, articleIndex, context.Param("blogID"), string(byteContent))
		if err == mongo.ErrNoDocuments || err == mongo.ErrInvalidIndexValue {
			_ = context.AbortWithError(404, err)
		} else if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.Status(200)
			context.Done()
		}
	})
	// 删除一篇文章
	apiEngine.DELETE("/blog/:blogID/:articleID", func(context *gin.Context) {
		blogID := context.GetString("blogID")
		articleIndex := context.GetInt("articleID")
		err := DeleteArticle(mongoCollection, articleIndex, blogID)
		if err == mongo.ErrNoDocuments || err == mongo.ErrInvalidIndexValue {
			_ = context.AbortWithError(404, err)
		} else if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.Status(200)
			context.Done()
		}
	})
	// 删除一个博客项目
	apiEngine.DELETE("/blog/:blogID", func(context *gin.Context) {
		blogID := context.GetString("blogID")
		err := DeleteBlogSubject(mongoCollection, blogID)
		if err == mongo.ErrNoDocuments {
			_ = context.AbortWithError(404, err)
		} else if err != nil {
			_ = context.AbortWithError(400, err)
		} else {
			context.Status(200)
			context.Done()
		}
	})
}
