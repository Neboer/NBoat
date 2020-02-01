// 用来给出一个article的brief content，可以作为博客的简单介绍。
package ritin

import (
	"encoding/json"
	"strings"
)

type Op []map[string]string

func GetBriefTextOfArticle(length int, deltaString string) string {
	opList := Op{}
	briefText := ""
	_ = json.Unmarshal([]byte(deltaString), &opList)
	for _, item := range opList {
		valueGoingToBeInsert := strings.ReplaceAll(item["insert"], "\n", "")
		if len(briefText)+len(valueGoingToBeInsert) > length {
			briefText += valueGoingToBeInsert[0 : length-len(briefText)]
		} else {
			briefText += valueGoingToBeInsert
		}
	}
	return briefText
}
