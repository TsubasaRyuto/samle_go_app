package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "log"
  "github.com/TsubasaRyuto/go_sample_app/config"
)

func main() {
  log.Println("start server ...")

  log.Println(config.Config.Port)

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
      "message": "HelloGo",
    })
  })

  r.Run()
}
