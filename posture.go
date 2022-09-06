package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Posture Data")
  })
  return router
}

func main() {
  router := setupRouter()
  router.Run()
}