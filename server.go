package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func routeFilesRequest(c *gin.Context) {
	path := c.Param("path")
	contentsParam := c.Query("contents")
	response := ""
	var err error

	switch c.Request.Method {
	case http.MethodGet:
		response, err = Get(path)
	case http.MethodPost:
		err = Create(path, contentsParam)
	case http.MethodPut:
		err = Update(path, contentsParam)
	case http.MethodDelete:
		err = Delete(path)
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, response)
}

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/files/*path", routeFilesRequest)
	r.POST("/files/*path", routeFilesRequest)
	r.PUT("/files/*path", routeFilesRequest)
	r.DELETE("/files/*path", routeFilesRequest)

	return r
}

func main() {
	os.Mkdir(dir, os.ModePerm)
	CreateRouter().Run(":8000")
}
