package models

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

type User struct {
  gorm.Model
  Name string `gorm:"not null"`
  Email string `gorm:"not null; unique"`
  PasswordHash string `gorm:not null"`
}

func CreateUser(name string, email string, passwordHash string) {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  db.Create(&User{Name: name, Email: email, PasswordHash: passwordHash})
}

func GetUserById(id int) {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  var user User
  db.First(&user, id)
}
