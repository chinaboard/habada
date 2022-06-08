package service

import (
	"github.com/chinaboard/habada/service/coder"
	"github.com/gin-gonic/gin"
	"net/http"
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
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("habada!"))
		c.Status(http.StatusOK)
	})

	return r
}
