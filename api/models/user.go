package models

import (
  "errors"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  gorm.Model
  Name string `gorm:"not null"`
  Email string `gorm:"not null; unique"`
  PasswordHash string `gorm:not null"`
}

func GetAllUsers() []User {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  var users []User
  db.Find(&users)
  return users
}

func CreateUser(name string, email string, passwordHash string) {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  encryptPass, err := bcrypt.GenerateFromPassword([]byte(passwordHash), bcrypt.DefaultCost)
  if err != nil {
    // error
  }

  db.Create(&User{Name: name, Email: email, PasswordHash: string(encryptPass)})
}

func GetUserById(id int) User {
  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  var user User
  db.First(&user, id)
  return user
}

func Login(email, password string) (*User, error) {
  user := User{}

  dsn := "host=db user=app_user password=password dbname=sample_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  db.Where("email = ?", email).First(&user)

  if user.ID == 0 {
    err := errors.New("ログインに失敗しました")
    return nil, err
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

  if err != nil {
    return nil, err
  }

  return &user, nil
}
