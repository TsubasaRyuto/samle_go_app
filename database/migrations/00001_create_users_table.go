package migrations

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name string `gorm:"not null"`
  Email string `gorm:"not null; unique"`
  PasswordHash string `gorm:not null"`
}
