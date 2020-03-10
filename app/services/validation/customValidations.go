package validation

import (
	"log"
	"strings"

	"github.com/ajangi/gAuthService/app/config"
	"github.com/jinzhu/gorm"
	// using to connect to mysql database!
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
)
var validate *validator.Validate
// UniqueInDB is a custom validatation to check data is uniqueu in database
func UniqueInDB(fl validator.FieldLevel) bool{
	db, err := gorm.Open("mysql", config.DbConfig)
	if err != nil {
		log.Panic(err)
	}
	param := strings.Split(fl.Param(), `:`)
	tableName := param[0]
	columnName := param[1]
	var count int
	db.Table(tableName).Where(columnName+" = ?", fl.Field().String()).Count(&count)
	if(count > 0){
		return false
	}
	return true
}