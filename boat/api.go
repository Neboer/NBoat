// boat有两套api，一套是“输入”标准，一套是“输出”标准，输入标准是新建博客时，前后端向boat插入的数据。输出标准是boat返回给前后端的数据，用来渲染前端页面。
package boat

type BlogInfo struct {
	BlogName string
	// 封面图片网址，这个应该在上传博客的时候就已经指定了
	CoverPictureURL string
	// 博客的内容。我们认为博客创建的时间和修改的时间就是博客正文内容改变的时间。
	BlogContent string
}
