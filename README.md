# Neboer's Boat
Neboer's Boat是"Neboer's Blog isn't Only About Technique"
的缩略形式。

这个项目是Neboer的个人博客。使用技术栈Native javascript + golang-gin + Mongodb

本项目包含三个子项目：Neboer's Boat、RTP-core和NPC。


## 流程介绍
项目中使用到了Quilljs作为前端的富文本编辑器，编辑完成后，浏览器会上传delta内容。

服务器根据Quilljs传入的delta，使用go-render-quill这种模块对HTML内容进行服务器端渲染
，最终将渲染结果与quill delta本身一起保存在数据库中。

一般用户访问时，服务器仅仅负责返回渲染后的HTML数据，当用户需要编辑的时候，再发送delta。

## Pastebin
项目本质上就是一个Pastebin，只是对权限加以限制而已。因此，Neboer's Boat可以轻易变成Neboer's
 Blog-Drived Pastebin, 可以称其为share-exchange中的一个rich text pastebin(RTP)

这个Nboat将会实现一个RTP-core,其中包含最基本的一些Pastebin功能。项目的其余部分都是围绕着core展开的。

可以说，博客最重要的内容部分是RTP处理的，这个项目很大一部分也是在实现RTP核心模块。

## 图床
这个项目当然也需要保存图片数据等等。在作者上传图片的时候，会异步请求到服务器，服务器端保存图片并返回其地址。
这也就相当于一个图床了。这个图床可以起名Neboer's Picture Server(NPC)

项目将会部署到我的github个人主页，至于其HTTPS后端……再另寻服务器吧。一个域名、一个证书什么的，足够了。

感谢开源世界为我们提供了如此多的的选择。