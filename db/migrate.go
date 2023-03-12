package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/TsubasaRyuto/go_sample_app/db"
  "github.com/TsubasaRyuto/go_sample_app/db/migrations"
)

func init() {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  migrate!()
}

func getDB() *gorm.DB {
  return db
}

func close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}

func migrate!() {
  db.AutoMigrate(&migrations.User{})
}
