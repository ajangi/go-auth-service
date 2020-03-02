package models

import (
	"github.com/jinzhu/gorm"
	"time"
  )

  // UserStatus as enum for user status like : active,banned,...
type UserStatus string
  // User model using gorm
type User struct {
	gorm.Model
	ID           uint         `gorm:"column:id;type:bigint(20) unsigned;primary_key;auto_increment;" json:"id"`
	Name         string       `gorm:"column:name;type:varchar(100);not null;" json:"name"`
	PhoneNumber  string       `gorm:"column:phone_number;type:varchar(11);unique_index" json:"phone_number"`
	UserName     string       `gorm:"column:username;type:varchar(255);unique;" json:"username"`
	Password     string       `gorm:"column:password;type:varchar(255);" json:"password"`
	Status       UserStatus   `gorm:"column:status;default:active" sql:"type:enum('active','banned')" json:"status"`
	CreatedAt    *time.Time
}