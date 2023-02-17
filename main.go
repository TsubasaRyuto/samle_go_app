package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "log"
)


func main() {
  log.Println("staart server ...")

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
      "message": "HelloGo",
    })
  })

  r.Run()
}
