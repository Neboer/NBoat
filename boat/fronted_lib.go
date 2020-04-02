package boat

import (
	"github.com/gin-gonic/gin"
)

func parseQueryIndex(ctx *gin.Context, maxIndex int) int {
	parseStruct := struct {
		Index int `form:"page"`
	}{}
	err := ctx.Bind(&parseStruct)
	if err != nil {
		_ = ctx.Error(err)
	}
	if parseStruct.Index > maxIndex {
		return maxIndex
	} else if parseStruct.Index == 0 {
		return 1
	} else {
		return parseStruct.Index
	}
}

func parseQuerySort(ctx *gin.Context) string {
	parseStruct := struct {
		Sort string `form:"sort"`
	}{}
	err := ctx.Bind(&parseStruct)
	if err != nil {
		_ = ctx.Error(err)
	}
	return parseStruct.Sort
}

// 这里应该在mongo里面进行查询的。但是博客这里设计偏向于查询的扩展性了。如果直接在mongo里检索未尝不可，但是我们并不考虑数据库脱机的情况，而且我认为
// golang的执行效率要快于mongo的javascript引擎。因此在这里我们单独设计一个搜索函数。
// all match为一个flag，设置为True的时候，只有当所有tag都符合的时候才会返回。

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func searchBlogByKey(allBlogList *BlogSubjectBriefList, keys []string, allMatch bool) BlogSubjectBriefList {
	searchResult := BlogSubjectBriefList{}
	if !allMatch {
		for _, blog := range *allBlogList {
			for _, userKey := range keys {
				if contains(blog.Info.Sort, userKey) {
					searchResult = append(searchResult, blog)
					goto CC
				}
			}
		CC:
		}
	} else {
		for _, blog := range *allBlogList {
			keyCount := 0
			for _, userKey := range keys {
				if contains(blog.Info.Sort, userKey) {
					searchResult = append(searchResult, blog)
					goto DD
				}
			}
		DD:
		}
	}

}
