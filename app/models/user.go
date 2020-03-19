package models

import (
	"fmt"
	"log"

	"github.com/ajangi/gAuthService/app/types/requests"
	"github.com/ajangi/gAuthService/app/utils/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserStatus as enum for user status like : active,banned,...
type userStatus string

// User model using gorm
type User struct {
	gorm.Model
	FirstName   string     `gorm:"column:first_name;type:varchar(100);not null;" json:"first_name"`
	LastName    string     `gorm:"column:last_name;type:varchar(100);not null;" json:"last_name"`
	PhoneNumber string     `gorm:"column:phone_number;type:varchar(11);unique_index" json:"phone_number"`
	Email       string     `gorm:"column:email;type:varchar(255);unique_index" json:"email"`
	UserName    string     `gorm:"column:username;default:null;type:varchar(255);unique;" json:"username"`
	Password    string     `gorm:"column:password;default:null;type:varchar(255);" json:"password"`
	Status      userStatus `gorm:"column:status;type:varchar(25);default:'active'" json:"status"`
	//CreatedAt   time.Time
}

// Migrate is the method to make user model migration
func migrate() {
	database, err := db.GetDB()
	if err != nil {
		log.Panic(err)
	}
	database.AutoMigrate(&User{})
	defer database.Close()
}

// TableName is to define user model table name
func (User) TableName() string {
	return "users"
}

// CreateUser is to create user
// Exports User model
func CreateUser(req *requests.RegisterUserRequest) (User,error) {
	var user = User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		UserName: req.UserName,
		Password: req.Password,
	}
	database, err := db.GetDB()
	db.MigrateTable("users",migrate)
	if err != nil {
		log.Panic(err)
		return user,err
	}
	dbError := database.Create(&user)
	if dbError != nil {
		return user,dbError.Error
	}
	fmt.Println("result : ",err)
	defer database.Close()
	return user,nil
}
