package models

import (
	"log"
	"fmt"

	"github.com/ajangi/gAuthService/app/utils/requests"
	"github.com/ajangi/gAuthService/app/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserStatus as enum for user status like : active,banned,...
type UserStatus string

// User model using gorm
type User struct {
	gorm.Model
	//ID          uint       `gorm:"column:id;type:bigint(20) unsigned;primary_key;auto_increment;" json:"id"`
	FirstName   string     `gorm:"column:first_name;type:varchar(100);not null;" json:"first_name"`
	LastName    string     `gorm:"column:last_name;type:varchar(100);not null;" json:"last_name"`
	PhoneNumber string     `gorm:"column:phone_number;type:varchar(11);unique_index" json:"phone_number"`
	Email       string     `gorm:"column:email;type:varchar(255);unique_index" json:"email"`
	UserName    string     `gorm:"column:username;default:null;type:varchar(255);unique;" json:"username"`
	Password    string     `gorm:"column:password;default:null;type:varchar(255);" json:"password"`
	Status      UserStatus `gorm:"column:status;type:varchar(25);default:'active'" json:"status"`
	//CreatedAt   time.Time
}

// Migrate is the method to make user model migration
func Migrate() {
	db, err := gorm.Open("mysql", config.DbConfig)
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&User{})
	defer db.Close()
}

// DropTable is the method to drop user model table
func DropTable() {
	db, err := gorm.Open("mysql", config.DbConfig)
	if err != nil {
		log.Panic(err)
	}
	db.DropTableIfExists(&User{}, "users")
	defer db.Close()
}

// TableName is to define user model table name
func (User) TableName() string {
	return "users"
}

// CreateUser is to create user
func CreateUser(req *requests.RegisterUserRequest) (User,error) {
	var user = User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		UserName: req.UserName,
		Password: req.Password,
	}
	db, err := gorm.Open("mysql", config.DbConfig)
	if(!db.HasTable("users")){
		Migrate()
	}
	if err != nil {
		log.Panic(err)
		return user,err
	}
	dbError := db.Create(&user)
	if dbError != nil {
		return user,dbError.Error
	}
	fmt.Println("result : ",err)
	defer db.Close()
	return user,nil
}
