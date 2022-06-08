package coder

import (
	"github.com/chinaboard/habada/pkg/urlshortener"
	"github.com/chinaboard/habada/storage"
	"github.com/chinaboard/habada/storage/database"
	"github.com/chinaboard/habada/storage/memory"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

var (
	memStorage storage.Storage
	dbStorage  storage.Storage
)

func init() {
	memStorage = memory.New()
	dbStorage = database.New()

}

func Encode(c *gin.Context) {
	longUrl := c.Query("longUrl")
	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	tinyUrl := urlshortener.Generate()

	if _, err = dbStorage.Set(tinyUrl, longUrl); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "server internal error"})
		return
	}

	if _, err = memStorage.Set(tinyUrl, longUrl); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "server internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tiny_url": tinyUrl})
}
func Decode(c *gin.Context) {
	tinyUrl := c.Query("tinyUrl")
	if len(tinyUrl) != urlshortener.CodeLength {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	longUrl, _ := memStorage.Get(tinyUrl)
	if longUrl != "" {
		c.JSON(http.StatusOK, gin.H{"data": longUrl})
		return
	}

	dblongUrl, err := dbStorage.Get(tinyUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "server internal error"})
		return
	}

	if dblongUrl != "" {
		memStorage.Set(tinyUrl, dblongUrl)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"long_url": dblongUrl})
		return
	}

	c.Status(http.StatusNotFound)
}

func Redirect(c *gin.Context) {
	tinyUrl := c.Param("tinyUrl")

	if len(tinyUrl) != urlshortener.CodeLength {
		c.Status(http.StatusNotFound)
		return
	}
	longUrl, _ := memStorage.Get(tinyUrl)
	if longUrl != "" {
		c.Redirect(http.StatusTemporaryRedirect, longUrl)
		return
	}

	dbLongUrl, err := dbStorage.Get(tinyUrl)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if dbLongUrl != "" {
		memStorage.Set(tinyUrl, dbLongUrl)
	}

	c.Redirect(http.StatusTemporaryRedirect, dbLongUrl)
}
