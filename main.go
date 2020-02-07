package main

import (
	"Nboat/auth"
	"Nboat/boat"
	"Nboat/dbWork"
	"Nboat/nopiser"
	"Nboat/ritin"
	"encoding/json"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type SecretContent struct {
	IdentityPassword string `json:"identity_password"`
}

var SecretObject SecretContent

func init() {
	secretFile, err := ioutil.ReadFile("secret.json")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		ReadSecretFileToObject(secretFile)
		if SecretObject.IdentityPassword == "" {
			log.Fatal("empty admin password!")
		}
	}
}

func main() {
	server := gin.Default()
	database := dbWork.ConnectionInit()
	nboatCollection, ritinCollection, nopiserCollection := dbWork.GetCollection(*database)
	// 身份验证～
	server.Use(auth.AuthenticGate(SecretObject.IdentityPassword))
	// 博客的开放页面，也不是谁都能访问的！这个开放页面是选择性渲染的哦！
	server.Use(static.ServeRoot("/", "front"))
	// 这么底层的操作怎么能让陌生人来做呢？駄目でう！
	boat.BindBoatRenderer(server.Group(""), nboatCollection, ritinCollection)
	apiServer := server.Group("/api")
	nopiser.BindNopiser(apiServer, nopiserCollection)
	ritin.BindRitin(apiServer, ritinCollection, auth.OnlyAllowAuthor())
	boat.BindBoatBackend(apiServer, nboatCollection, ritinCollection)

	_ = server.Run(":8080")
}

func ReadSecretFileToObject(secretJson []byte) {
	err := json.Unmarshal(secretJson, &SecretObject)
	if err != nil {
		log.Fatal(err.Error())
	}
}
