package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", show)
	router.POST("/", create)
	return router
}

func show(c *gin.Context) {
	c.String(http.StatusOK, "Posture Data")
}

func create(c *gin.Context) {
	var posture string = c.PostForm("posture")
	c.String(http.StatusCreated, "Created posture=%s", posture)
}
