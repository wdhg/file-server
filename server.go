package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/files/*path", func(c *gin.Context) {
		contents, err := Get(c.Param("path"))

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		c.String(http.StatusOK, contents)
	})

	r.POST("/files/*path", func(c *gin.Context) {
		err := Create(c.Param("path"), c.Query("contents"))

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
	})

	r.PUT("/files/*path", func(c *gin.Context) {
		err := Update(c.Param("path"), c.Query("contents"))

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
	})

	r.DELETE("/files/*path", func(c *gin.Context) {
		err := Delete(c.Param("path"))

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
	})

	return r
}

func main() {
	os.Mkdir(dir, os.ModePerm)
	CreateRouter().Run(":8000")
}
