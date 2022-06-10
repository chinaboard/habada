package service

import (
	"github.com/chinaboard/habada/pkg/bininfo"
	"github.com/chinaboard/habada/service/coder"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/encode", coder.Encode)
		api.GET("/decode", coder.Decode)
	}

	r.GET("/:tinyUrl", coder.Redirect)
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"BuildTime": bininfo.BuildTime})
	})

	cwd, _ := os.Getwd()

	r.LoadHTMLFiles(path.Join(cwd, "view/index.html"))
	r.Static("/static", path.Join(cwd, "static"))
	r.StaticFile("/favicon.ico", path.Join(cwd, "static/favicon.ico"))

	return r
}
