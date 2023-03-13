package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/TsubasaRyuto/go_sample_app/database/migrations"
)

func Migrate() {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&migrations.User{})
}
