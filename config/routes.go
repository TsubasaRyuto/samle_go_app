package config

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/TsubasaRyuto/go_sample_app/api/controllers"
)

func ServerStart() {
  router := gin.Default()
  router.Use(CORS())
  RegisterRoutes(router)
  router.Run()
}

func userRoutes(router *gin.RouterGroup) {

  router.GET("", controllers.GetUsers)
  router.GET("/:userid", controllers.GetUser)
  router.POST("", controllers.CreateUser)
  // router.PUT("/users/:id", usersController.Update)
  // router.DELETE("/users/:id", usersController.Update)
}

func sessionRouters(router *gin.RouterGroup) {
  router.POST("/login", controllers.Login)
}

func RegisterRoutes(router *gin.Engine) {
  apiRoutes := router.Group("/api")
  {
    userRoutes(apiRoutes.Group("/users"))
    sessionRouters(apiRoutes.Group(""))
  }
}

func CORS() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(http.StatusNoContent)
      return
    }

    c.Next()
  }
}
