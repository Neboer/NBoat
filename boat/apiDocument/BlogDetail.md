# *打开博客

返回一个网页，展示一篇博文的详细内容。

**URL** : `/blog/:blogid`

**方法** : `GET`

**身份等级** : 游客/我

## 请求成功

**状态码** : `200 OK`

## 请求错误

如果请求的博文不存在，返回404；请求不合法则返回400，注意，如果博文的delta值格式错误，
请求也不会接受。

**状态码** : `400/404`

`400`: 请求格式错误，即blogid字段不合规范

`404`：请求的博客不存在

认证用户“我”与游客看到的内容是不同的。

```json
{
  "error": "too long blog_name"
}
```

```json
{
  "error": "cover_picture_url is invalid"
}
```

```json
{
  "error": "blog content delta is invalid"
}
```
如果用户发出了未认证的请求，那么将返回`401 Unauthorized`，此响应没有响应体。