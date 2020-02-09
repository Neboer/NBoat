package main

import (
	"Nboat/boat"
	"Nboat/cookieauth"
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
	// 身份验证，把身份绑定到请求参数上。
	server.Use(cookieauth.AuthenticGate(SecretObject.IdentityPassword))
	// 访问指定的密码页面，会收到同样的密码cookie。
	cookieauth.SetIdentity(server, SecretObject.IdentityPassword)
	server.Use(static.ServeRoot("/", "front"))
	// 这么底层的操作怎么能让陌生人来做呢？駄目でう！
	boat.BindBoatRenderer(server.Group(""), nboatCollection, ritinCollection)
	apiServer := server.Group("/api")
	// 目前我们不开放公共图床，所以图床的上传需要控制。
	nopiser.BindNopiser(apiServer, nopiserCollection)
	// Ritin模块禁止访问。
	ritin.BindRitin(apiServer, ritinCollection, cookieauth.OnlyAllowAuthor(""))
	boat.BindBoatBackend(apiServer, nboatCollection, ritinCollection)

	_ = server.Run(":8080")
}

func ReadSecretFileToObject(secretJson []byte) {
	err := json.Unmarshal(secretJson, &SecretObject)
	if err != nil {
		log.Fatal(err.Error())
	}
}
