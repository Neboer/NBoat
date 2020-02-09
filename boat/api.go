// boat有两套api，一套是“输入”标准，一套是“输出”标准，输入标准是新建博客时，前后端向boat插入的数据。输出标准是boat返回给前后端的数据，用来渲染前端页面。
package boat

import (
	"Nboat/ritin"
	quill "github.com/dchenk/go-render-quill"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
)

// 作为博客列表的一项输出，注意这个输出的目标是直接参与渲染的。
type BlogInfoItem struct {
	BlogHexID string
	BlogName  string
	// 封面图片网址，这个应该在上传博客的时候就已经指定了
	CoverPictureURL template.URL
	// 我们认为博客创建的时间和修改的时间就是博客正文内容改变的时间。
	CreateTimeString string
	BlogBriefContent string // 这是博客的简短内容表述，由ritin模块生成。
}

// 作为插入的一个博文。如果前端想要创建一个新可插入博文，就需要填写以下结构。注意，这个结构不包含时间，对博文分类暂不实现。
type BlogIn struct {
	BlogName         string `json:"blog_name"`
	CoverPictureURL  string `json:"cover_picture_url"`
	BlogDeltaContent string `json:"blog_delta_content"`
}

type BlogOut struct {
	BlogHexId          string
	BlogName           string
	CoverPictureURL    template.URL
	BlogArticleHTML    template.HTML // 这是编译之后的delta内容，是html形式的哦
	CreateTimeString   string
	LastModifiedString string
}

// 用来编辑的blog ,相当于作者模式对象。
type BlogEdit struct {
	BlogName         string
	BlogDeltaContent string
}

// 对于每个博文，都请求一次数据库查找对应的ritin内容，这样的操作极大的降低了效率。不过这样虽然带来了额外的开销，但我坚信它是有好处的。
func GetBlogBriefList(boatCollection *mongo.Collection, ritinCollection *mongo.Collection) []BlogInfoItem {
	blogBriefList := make([]BlogInfoItem, 0)
	blogList := getBlogListFromMongoCollection(boatCollection)
	for _, blog := range blogList {
		article, _ := ritin.GetArticle(blog.RelativeRitinID.Hex(), ritinCollection)
		currentBlogInfo := BlogInfoItem{
			BlogHexID:        blog.ID.Hex(),
			BlogName:         blog.BlogName,
			CoverPictureURL:  template.URL(blog.CoverPictureURL),
			CreateTimeString: article.CreateTime.String()[:19],
			BlogBriefContent: ritin.GetBriefTextOfArticle(80, article.Content), // 准备做一个提取quill delta文本内容的生成器。// 已经做出来了！
		}
		blogBriefList = append(blogBriefList, currentBlogInfo)
	}
	return blogBriefList
}

func InsertBlog(blog BlogIn, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) string {
	ritinArticleContent := blog.BlogDeltaContent
	articleID := ritin.InsertArticle(ritinArticleContent, ritinCollection)
	blogID := insertBlogToMongoCollection(BlogInRecord{
		BlogName:           blog.BlogName,
		CoverPictureURL:    blog.CoverPictureURL,
		RelativeRitinHexID: articleID,
	}, boatCollection)
	return blogID
}

// 注意：这个getblog返回的是直接可以供顶层使用的blogout对象，这个对象不应该用在除了前端渲染器之外的任何地方，在大多数情况下，调用boat的mongoer里的api来解决问题
func GetBlog(blogHexID string, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) (BlogOut, error) {
	blogRecord, err := getBlogFromMongoCollection(blogHexID, boatCollection)
	if err != nil {
		return BlogOut{}, err
	} else {
		ritinArticleHexID := blogRecord.RelativeRitinID.Hex()
		ritinArticle, _ := ritin.GetArticle(ritinArticleHexID, ritinCollection)
		articleHTMLbytes, _ := quill.Render([]byte(ritinArticle.Content))
		outputBlog := BlogOut{
			BlogHexId:          blogHexID,
			BlogName:           blogRecord.BlogName,
			CoverPictureURL:    template.URL(blogRecord.CoverPictureURL),
			BlogArticleHTML:    template.HTML(string(articleHTMLbytes)),
			CreateTimeString:   ritinArticle.CreateTime.String()[0:19],
			LastModifiedString: ritinArticle.LastModifyTime.String()[0:19],
		}
		return outputBlog, nil
	}
}

// 当用户希望编辑一个博客内容的时候，应该返回其delta了。
func GetBlogDelta(blogHexID string, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) (BlogEdit, error) {
	blogRecord, err := getBlogFromMongoCollection(blogHexID, boatCollection)
	if err != nil {
		return BlogEdit{}, err
	} else {
		ritinArticleHexID := blogRecord.RelativeRitinID.Hex()
		ritinArticle, _ := ritin.GetArticle(ritinArticleHexID, ritinCollection)
		return BlogEdit{
			BlogName:         blogRecord.BlogName,
			BlogDeltaContent: ritinArticle.Content,
		}, nil
	}
}

// 这个操作仅仅可以更新博文文章内容。
func UpdateBlogContent(blogHexID string, newArticleDeltaContent string, boatCollection *mongo.Collection, ritinCollection *mongo.Collection) error {
	blogRecord, err := getBlogFromMongoCollection(blogHexID, boatCollection)
	if err != nil {
		return err
	} else {
		ritin.UpdateArticle(blogRecord.RelativeRitinID.Hex(), ritinCollection, newArticleDeltaContent)
		return nil
	}
}
