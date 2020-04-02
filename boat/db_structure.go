package boat

import (
	"time"
)

// type BlogSubject
type BlogSubject struct {
	Info    BlogSubjectInfo `bson:"info"`
	Meta    BlogSubjectMeta `bson:"meta"`
	Article []Article       `bson:"article"`
}

type BlogSubjectMeta struct {
	CreateTime     time.Time `bson:"create_time"`
	LastModifyTime time.Time `bson:"last_modify_time"`
	ReadCount      int       `bson:"read_count"`
}

// 用户可编辑属性，用来创建和更新博客项目本身。
type BlogSubjectInfo struct {
	Title           string   `bson:"title"`
	Introduction    string   `bson:"introduction"`
	CoverPictureURL string   `bson:"cover_picture"`
	Sort            []string `bson:"sort"`
}

type BlogSubjectBriefList []struct {
	ID   string          `bson:"_id"`
	Info BlogSubjectInfo `bson:"info"`
	Meta BlogSubjectMeta `bson:"meta"`
}

//type Article
type Article struct {
	Index      int         `bson:"index"` // 第一个article的index为0
	Info       ArticleInfo `bson:"info"`
	Meta       ArticleMeta `bson:"meta"`
	Content    string      `bson:"content"`     // 一般，这里都会保存HTML结构化数据，用来直接展示博客文章内容
	RawContent string      `bson:"raw_content"` // 这里储存原始的quill.js的delta和markdown原文件，便于渲染。
}

type ArticleMeta struct {
	CreateTime     time.Time `bson:"create_time"`
	LastModifyTime time.Time `bson:"last_modify_time"`
	Editor         string    `bson:"editor"` // 支持markdown和quill，在创建article的时候就应该指定并不再可以修改。
}

type ArticleInfo struct {
	Name            string   `bson:"name"`
	CoverPictureURL string   `bson:"cover_picture_url"`
	Key             []string `bson:"key"`
	Draft           bool     `bson:"is_draft"` // 是否为草稿
}

// 插入新的文章或者整个更新文章时要用到
type ArticleInput struct {
	Info       ArticleInfo
	Content    string
	RawContent string
	Editor     string // 注意，editor是单独传递的，目的只有一个就是为了区分不同的渲染平台……
}

// 用户必须先创建博客项目，然后一篇一篇的写入文章。
func (info BlogSubjectInfo) toBlogSubject() BlogSubject {
	return BlogSubject{
		Info: info,
		Meta: BlogSubjectMeta{
			CreateTime:     time.Now(),
			LastModifyTime: time.Now(),
			ReadCount:      0,
		},
		Article: []Article{},
	}
}

func (input ArticleInput) toArticle(index int) Article {
	return Article{
		Info: input.Info,
		Meta: ArticleMeta{
			CreateTime:     time.Now(),
			LastModifyTime: time.Now(),
			Editor:         input.Editor,
		},
		Index:      index,
		Content:    input.Content,
		RawContent: input.RawContent,
	}
}
