package boat

import (
	"Nboat/dbWork"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"testing"
)

type SecretContent struct {
	IdentityPassword string `json:"identity_password"`
	MongoLocation    string `json:"mongo_ip:port"`
	Username         string `json:"mongo_username"`
	Password         string `json:"mongo_password"`
}

var SecretObject SecretContent

func ReadSecretFileToObject(secretJson []byte) {
	err := json.Unmarshal(secretJson, &SecretObject)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMongoCollection() *mongo.Collection {
	secretFile, err := ioutil.ReadFile("../secret.json")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		ReadSecretFileToObject(secretFile)
		if SecretObject.IdentityPassword == "" {
			log.Fatal("empty admin password!")
		}
	}
	database := dbWork.ConnectionInit(SecretObject.MongoLocation, SecretObject.Username, SecretObject.Password)
	testCollection := database.Collection("nboat")
	return testCollection
}

func TestFullWorkflow(t *testing.T) {
	testCollection := GetMongoCollection()
	blogSubjectID, err := CreateEmptyBlogSubject(testCollection, BlogSubjectInfo{
		Title:        "a test blog obj",
		Introduction: "仅仅是供我们进行测试用的！",
		Sort:         []string{"测试", "fgo", "jsp"},
	})
	if err != nil {
		fmt.Println("CreateEmptyBlogSubject error", err)
	}
	newArticle := ArticleInput{
		Info: ArticleInfo{
			Name:            "my article",
			CoverPictureURL: "cover",
			Key:             []string{"good", "uefi"},
			Draft:           false,
			Editor:          "quill",
		},
		Content: "this is the content of the inserted article",
	}
	id, err := InsertArticle(testCollection, newArticle, blogSubjectID)
	if err != nil {
		fmt.Println("InsertArticle error", err)
	}
	newArticleInfo := ArticleInfo{
		Name:            "new info is here",
		CoverPictureURL: "new cover!",
		Key:             []string{"dddd", "ddddder"},
		Draft:           false,
		Editor:          "markdown",
	}
	err = UpdateArticleInfo(testCollection, id, blogSubjectID, newArticleInfo)
	if err != nil {
		fmt.Println("UpdateArticleInfo error", err)
	}
	blogBriefList, err := GetBlogSubjectList(testCollection)
	if err != nil {
		fmt.Println("GetBlogSubjectList error", err)
	}
	blogSubject, err := getBlogSubject(testCollection, blogBriefList[0].ID)
	if err != nil {
		fmt.Println("getBlogSubject error", err)
	}
	fmt.Println(blogSubject.Article[0].Content)
	err = UpdateArticleContent(testCollection, id, blogSubjectID, "this is the new content of the article!")
	if err != nil {
		fmt.Println("UpdateArticleContent error", err)
	}
	err = DeleteArticle(testCollection, id, blogSubjectID)
	if err != nil {
		fmt.Println("DeleteArticle error", err)
	}

}
