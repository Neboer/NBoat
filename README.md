# Neboer's Boat

> 这注定是一个不平凡的2020,愿一切安好。

<p align="center"><img src="https://bit.ly/2TMhT93" width="250"></p>

<h4 align="center">宮水 三葉</h4>

> miyamizu mitsuha (kimi no na wa) drawn by takahiro-rikky

> 这并不是这个博客项目的图标！！！仅仅是因为现在logo没做好用这个凑数罢了QAQ

Neboer's Boat是"Neboer's Blog isn't Only About Technique"
的缩略形式。

这个项目是Neboer的个人博客。使用技术栈Native javascript + golang-gin + Mongodb

本项目包含三个子项目：Neboer's Boat、Ritin和Nopiser。


## 开发初衷
这是我的个人主页，也是个人的博客。博客是挺平常的一种东西，就是文章而已，本身也只是一个载体。
这个项目基于众多的开源项目开发，目前应该还在持续集成。本博客完全开源，欢迎各位向项目中贡献代码。

## Neboer's Boat
博客的大前端。

Neboer's Boat以Ritbin作为内容的编辑器和阅读器，发表博文的过程视为向Ritbin提交。
项目本身仅仅是Ritin+index而已，多了一个视图外壳和一个目录index，更多的效果和后端的访问控制。

仅此而已。

nopiser、ritin、boat的后后端都是函数库；后前端渲染器负责渲染页面；
单独的前后端负责迎接用户的直接请求、调用各个函数并根据函数的运行结果渲染页面本身，最终返回给用户。
## Ritin
Ritin是NBoat的rich text pastebin。

Ritin使用Quill.js作为前端的富文本编辑器，编辑完成后，浏览器会上传delta内容。
对图片等多媒体内容，Ritin会上传到图床Nopiser，并在本地留存其reference。

服务器根据客户端发送的、由Quill.js生成的delta，在内部对HTML内容进行服务器端渲染，填充至模板然后发送。

一般用户访问时，服务器仅仅负责返回渲染后的HTML数据，当用户需要编辑的时候，再发送delta。

Ritin是博客的核心，每篇博文本质上都是一次Paste的过程。
## Nopiser
这个项目当然也需要保存图片数据等等。在作者上传图片的时候，会异步请求到服务器，服务器端保存图片并返回其地址。
这也就相当于一个图床了。这个图床可以起名Neboer's Picture Server(Nopiser)
## future
项目将会部署到我的主页中。至于其HTTPS后端……再另寻服务器吧。一个域名、一个证书什么的，足够了。


### current state
（越新的状态就越会放在下面）
- 感谢开源世界为我们提供了如此多的的选择。
- 委托某朋友设计Logo
- 打算函数化核心库
- 完成核心函数的更新，并留出api接口
- 完成前后端的设计与开发
- 正在鉴权
- 正在记忆五十音图（我在干啥？
- 鉴权成功，平假名差不多熟了吧。。。