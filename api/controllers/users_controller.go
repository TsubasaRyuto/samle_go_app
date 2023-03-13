package controllers

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "github.com/TsubasaRyuto/go_sample_app/api/models"
)

func GetUsers(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "name": "TsubasaRyuto",
    "email": "tsubasa.ryuto@gmail.com",
  });
}

func GetUser(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("userid"))

  if err != nil {
    // error
  }

  user := models.GetUserById(id)
  c.JSON(http.StatusOK, gin.H{
    "name": user.Name,
    "email": user.Email,
  });
}

func CreateUser(c *gin.Context) {
  name := c.PostForm("name")
  email := c.PostForm("email")
  password := c.PostForm("password")
  models.CreateUser(name, email, password)
  c.JSON(http.StatusOK, gin.H{
    "message": "ok",
  });
}
