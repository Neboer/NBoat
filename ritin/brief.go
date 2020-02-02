// 用来给出一个article的brief content，可以作为博客的简单介绍。
package ritin

import (
	"encoding/json"
	utf8str "golang.org/x/exp/utf8string"
	"strings"
)

type Op []map[string]string

func GetBriefTextOfArticle(length int, deltaString string) string {
	opList := Op{}
	briefText := ""
	totalLength := 0
	_ = json.Unmarshal([]byte(deltaString), &opList)
	for _, item := range opList {
		valueGoingToBeInsert := strings.ReplaceAll(item["insert"], "\n", "")
		encStrToInsert := utf8str.NewString(valueGoingToBeInsert)
		smallLength := encStrToInsert.RuneCount()
		if totalLength+smallLength > length {
			briefText += encStrToInsert.Slice(0, length-totalLength)
			break
		} else {
			briefText += encStrToInsert.String()
		}
		totalLength += smallLength
	}
	return briefText
}
