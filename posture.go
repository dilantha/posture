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
	return router
}

func show(c *gin.Context) {
	c.String(http.StatusOK, "Posture Data")
}
