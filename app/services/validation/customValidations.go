package validation

import (
	"fmt"
	"log"
	"strings"

	"github.com/ajangi/gAuthService/app/config"
	"github.com/jinzhu/gorm"
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
	fmt.Println(tableName)
	fmt.Println(columnName)
	fmt.Println(fl.Field().String())
	var count int
	db.Table(tableName).Where(columnName+" = ?", fl.Field().String()).Count(&count)
	if(count > 0){
		return false
	}
	return true
}