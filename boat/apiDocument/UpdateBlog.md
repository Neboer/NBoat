# 更新博文内容

首先，进入博客更新页面，这时编辑器会自动跳出，在编辑器中完成编辑后，单击“上传”完成博客的更新。
这里的更新接口仅仅是“上传”的过程，“编辑”页面请参见[编辑博文](EditBlog.md)页面

**URL** : `/api/blog/:blogid`

**方法** : `PUT`

**请求体**：

```json
{
  "blog_delta_content": "博客文章的新的delta内容"
}
```

**身份等级** : 我


**请求样例**

```json
{
  "blog_delta_content": "[{\"insert\":\"This \"},{\"attributes\":{\"italic\":true},\"insert\":\"is\"},\n    {\"insert\":\" \"},{\"attributes\":{\"bold\":true},\"insert\":\"great!\"},{\"insert\":\"\\n\"}]"
}
```

## 请求成功

**状态码** : `200 OK`

如果更新成功，则响应体为空。

## 请求错误

**状态码** : `400/401/404`

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

如果用户更新的博文不存在，则返回`404 blog not found`，此响应没有响应体。