# 上传博客

经过认证的用户可以上传一篇博客。

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
  "blog_id": "博客的hex id。"
}
```

**Content example** : Response will reflect back the updated information. A
User with `id` of '1234' sets their name, passing `UAPP` header of 'ios1_2':

```json
{
    "id": 1234,
    "first_name": "Joe",
    "last_name": "Bloggs",
    "email": "joe25@example.com",
    "uapp": "ios1_2"
}
```

## Error Response

**Condition** : If provided data is invalid, e.g. a name field is too long.

**Code** : `400 BAD REQUEST`

**Content example** :

```json
{
    "first_name": [
        "Please provide maximum 30 character or empty string",
    ]
}
```

## Notes

* Endpoint will ignore irrelevant and read-only data such as parameters that
  don't exist, or fields that are not editable like `id` or `email`.
* Similar to the `GET` endpoint for the User, if the User does not have a
  UserInfo instance, then one will be created for them.
