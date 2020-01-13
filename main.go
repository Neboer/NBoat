package main

import (
	"Nboat/boat"
	"Nboat/dbWork"
	"Nboat/nopiser"
	"Nboat/ritin"
	"github.com/dchenk/go-render-quill"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.Use(static.ServeRoot("/", "front"))
	boat.BindBoatRenderer(server.Group(""))
	mainServer := server.Group("/api")

	database := dbWork.ConnectionInit()

	nopiser.BindNopiser(mainServer, database)
	ritin.BindRitin(mainServer, database)

	server.GET("/er", func(ctx *gin.Context) {
		delta := `[{"insert":"This "},{"attributes":{"italic":true},"insert":"is"},
    {"insert":" "},{"attributes":{"bold":true},"insert":"great!"},{"insert":"\n"}]`

		html, err := quill.Render([]byte(delta))
		if err != nil {
			panic(err)
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", html)
	})
	_ = server.Run(":8080")
}
