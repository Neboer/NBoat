package boat

import "github.com/gin-gonic/gin"

func parseQueryIndex(ctx *gin.Context, maxIndex int) int {
	parseStruct := struct {
		Index int `form:"page"`
	}{}
	err := ctx.Bind(&parseStruct)
	if err != nil {
		_ = ctx.Error(err)
	}
	if parseStruct.Index > maxIndex {

	}
	return parseStruct.Index
}
