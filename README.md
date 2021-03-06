# Neboer's Boat

## 本项目已经完成重构，请移步新仓库[Nboat2](https://github.com/neboer/Nboat2)，网址不变，谢谢。

https://www.neboer.site

> **武汉肺炎** - 这注定是一个不平凡的2020,愿一切安好。

<p align="center"><img src="https://bit.ly/2TMhT93" width="250"></p>

<h4 align="center">宮水 三葉</h4>

> miyamizu mitsuha (kimi no na wa) drawn by takahiro-rikky

> 这并不是这个博客项目的图标！！！仅仅是因为现在logo没做好用这个凑数罢了QAQ

Neboer's Boat是"Neboer's Blog isn't Only About Technique"
的缩略形式。

这个项目是Neboer的个人博客。使用技术栈Native javascript + golang-gin + Mongodb

本项目包含三个子项目：Neboer's Boat、Ritin和Nopiser。

## 写在前面

请注意，这是一个**个人博客项目**。

这就意味着，这并不是一个通用的网站项目，你当然可以自由的在协议下使用它的源代码，
但是这个博客从设计上就已经加入了大量的“个性化”元素，并且关于它的各个功能，并不配套有详细的文档。

如果你需要使用这个项目做进一步的开发和应用，或者按照这个模板做你自己的博客软件，当然也是没有问题的，
但你很难从这个仓库里直接获得很多技术支持，可能需要自行阅读并理解源代码。代码里的注释比较丰富，
这也使得你在阅读过程中不会遇到太多困难。

需要任何技术支持，欢迎随时邮件与我联系。作者Neboer是一个热心肠的人，非常愿意解决你的问题。


## Neboer

Neboer这个名字取自于“初生牛犊”的 *New(ly) born ...* 每个词的前两个字，并在结尾加上“er”表示“人”。
这是一个很奇怪的拼合词，严格来说并不符合拼读规范。你可以把它读成“neighbor”这个词，这也是我推荐的读法。

Neboer是一个不太极端的技术爱好者，现居中国大陆。在技术上偏向网页设计，全栈懵逼，对前端技术较为热衷，有一定的js水平，
也喜欢linux系列，写过一些实用程序和网站，是一个开源软件倡导者。

爱好上应该最喜欢的就是程序的设计和实现，目前有在学习日本语，单身，喜欢几个动漫角色。最喜欢的动画是电磁炮，
生活里的朋友不多，喜欢听一些日风轻音乐和享受独自一人的感觉。

你可以用中文或英语和我交流，发送邮件时请优先选择我的gmail邮箱。

没什么好说的，也没什么好了解的。关于Neboer我也所知不多，大概就是一只程序猿而已。

## 重构计划

*本项目从即日起开始重构*。在重构结束后，这个说明将会删除。此次重构旨在删去一些冗余的设计，使其扩展性更加强大。

- *Quill编辑器不再传出delta，改为传出html源文档。在编辑的时候，直接把html文件内容渲染进编辑器框内。
- *nopiser会进行一些处理，使之更适合上传图片。
- **打算使用阿里云oss作为图床（略有昂贵！）
- ritin和nboat存在一些冲突的地方。很难十分明确的界定二者之间的区别了。我的建议是发挥mongo的优势，改名字。
一条nboat记录可以不代表“一篇博文”了，我们把它叫做“一个博客项目”。一个博客项目中可以包含很多篇文章，它们可以随时增减。
博客项目单独储存在mongodb中，文章储存在对应的博客项目里，就不单独拿出来放在ritin中了。
到时候，在访问一个博客项目的时候，可以切换其中的一页一页的博文。
- 博客提交界面将会进行一些改动，比如可能会增加上传进度的提示。以及多种上传选择。

分割线下是重构后将持续开发的内容
<hr>

- 博客将会支持博客分类，到时候“分类”页面就可以使用了。
- 博客会支持关联博文……
- 博客还打算和[telegra.ph](https://telegra.ph)对接一下之类。具体对接的形式很有可能是我们在本地使用telegraph的简单编辑器写文章，
然后交到博客后，博客将会发到telegraph。用户请求这个博文的时候，博客将会从telegraph请求并返回其内容。这个属于不依赖数据库的存储形式，可以用来存一些永久存在的东西之类。
因为telegra.ph这个地方真的很优秀，所以我想最起码一定要用一下啊。
- 适配移动端。注意Bootstrap本来就是移动端优先适配的，但是博客这里并没有彻底发挥出它的优势，因此我们还是适配一下比较好。
- 除了telegra.ph，博客还可以支持多种不同的形式编辑、上传博文，比如wiki、enhanced_markdown或github_markdown之类。
- 增加评论区。评论也可以使用开源项目，正好也可以存在mongodb的“博文”对象中。当然你也可以对一个博客项目评论。评论将支持讨论版的形式，具体呈现形式暂定。

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
## 部署
执行这句话，以便于下载全部依赖到指定的区域。
```shell script
wget -i getlibrary.txt --directory-prefix ./front/library
``` 
## future
项目已经部署到我的主页中，我会一边开发维护一边写博客的……


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
- 虽然部署，但是还有问题。
- 重构计划
- 明确了数据库存储模型。
- 重新定义了数据库存储模型
- 实现了后端
