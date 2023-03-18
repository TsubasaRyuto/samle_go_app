package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/TsubasaRyuto/go_sample_app/api/models"
)

func Login(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")

  user, err := models.Login(email, password)

  if err != nil {
    // error
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "ok",
    "user": user,
  })
}
