package models

import (
	"github.com/ajangi/gAuthService/app/services/verification"
	"github.com/ajangi/gAuthService/app/utils/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// VerificationCode model using gorm
type VerificationCode struct {
	gorm.Model
	User User     `gorm:"foreignkey:UserID"`
	UserID uint    `gorm:"column:user_id;" json:"user_id"`
	Code string   `gorm:"column:code;type:varchar(4);not null;" json:"code"`
	Status string `gorm:"column:status;type:varchar(25);default:'fresh'" json:"status"`
}

// Migrate is the method to make user model migration
func (VerificationCode) migrate() {
	database, err := db.GetDB()
	if err != nil {
		log.Panic(err)
	}
	database.AutoMigrate(&VerificationCode{})
	defer database.Close()
}

// TableName is to define user model table name
func (VerificationCode) TableName() string {
	return "verification_codes"
}

// CreateCode is to create user
func CreateCode(user User) error {
	database, err := db.GetDB()
	if err != nil {
		log.Panic(err)
	}
	code := verification.GenerateVerificationCode(4)
	var verificationCode = VerificationCode{
		User:   user,
		UserID: user.ID,
		Code:   code,
		Status: "fresh",
	}
	db.MigrateTable(verificationCode.TableName(),verificationCode.migrate)
	dbError := database.Create(&verificationCode)
	if dbError != nil {
		return dbError.Error
	}
	defer database.Close()
	return nil
}