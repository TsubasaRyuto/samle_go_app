package main

import (
  // "net/http"
  // "github.com/gin-gonic/gin"
  "log"
  "github.com/TsubasaRyuto/go_sample_app/config"

  // "fmt"
  // "gorm.io/gorm"
  // "gorm.io/driver/postgres"
)

// type Test struct {
//   gorm.Model
//   Hoge string
// }

func main() {
  log.Println("start server ...")

  log.Println(config.Config.Port)


  // dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  // if err != nil {
  //   panic("failed to connect database")
  // }

  // // Migrate the schema
  // db.AutoMigrate(&Test{})

  // // Create
  // db.Create(&Test{ Hoge: "hogehoge" })

  // // Read
  // var test Test
  // db.First(&test, 2)

  config.ServerStart()

}
