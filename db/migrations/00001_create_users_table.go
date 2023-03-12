package migrations

import (
  "gorm.io/gorm"
)

type CreateUserTable struct {}

func (m *CreateUserTable) Up(db *gorm.DB) error {
  return db.CreateTable(&User{}).Error
}

func (m *CreateUserTable) Down(db *gorm.DB) error {
  return db.DropTableIfExists(&User{}).Error
}

type User struct {
  gorm.Model
  Name string `gorm:"not null"`
  Email string `gorm:"not null; unique"`
  PasswordHash string `gorm:not null"`
}
