package ritin

import (
	"fmt"
	"testing"
)

func TestBrief(t *testing.T) {
	delta := `[{"insert":"嘛，确实有点不容易啦。这篇博客要用到一些高级样式什么的了。\n\n"},
{"attributes":{"size":"large"},"insert":"用很大的字体来 做一个板块。"},{"insert":"\n"},
{"attributes":{"size":"small"},"insert":"巨大的第一标题"},
{"attributes":{"header":1},"insert":"\n"}
,{"insert":"\n不要被这个标题吓到，我们还可以插入很多有趣的东西。\n "}
,{"insert":{"image":"/api/nopiser/picture/5e34f4b49ecb753f53490d98"}},{"insert":"\n这个是跟进的！虽然图片很大，但是很可爱，卡哇伊\n\n新型冠状病毒\n"}]`
	fmt.Println(GetBriefTextOfArticle(450, delta))
}
