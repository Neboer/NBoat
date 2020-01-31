# 上传博客

经过认证的用户可以上传一篇博客。这里是上传一篇博客的后端，这里仅仅是“上传”的过程，“编辑”新博文的页面请参考
[编辑博文](EditBlog.md)

**URL** : `/api/boat/blog`

**方法** : `POST`

**请求体**：

```json
{
  "blog_name": "博客名字",
  "cover_picture_url": "封面图片地址",
  "blog_delta_content": "博客文章delta内容"
}
```

**身份等级** : 我


**请求样例**

```json
{
  "blog_name": "第一篇博文",
  "cover_picture_url": "https://www.baidu.com/picture/XXXXXX",
  "blog_delta_content": "[{\"insert\":\"This \"},{\"attributes\":{\"italic\":true},\"insert\":\"is\"},\n    {\"insert\":\" \"},{\"attributes\":{\"bold\":true},\"insert\":\"great!\"},{\"insert\":\"\\n\"}]"
}
```

## 请求成功

**状态码** : `200 OK`

**响应体** ： 

```json
{
  "blog_id": "博客的hex id"
}
```

## 请求错误


**状态码** : `400/401`

如果是请求体格式错误，`400`不带响应体。如果是name不合法/url错误/delta格式错误，
`400`响应体为如下格式。

```json
{
  "error": "too long blog name"
}
```

```json
{
  "error": "cover picture url is invalid"
}
```

```json
{
  "error": "blog content delta is invalid"
}
```
如果用户发出了未认证的请求，那么将返回`401 Unauthorized`，此响应没有响应体。