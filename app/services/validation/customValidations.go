package validation

import (
	"log"
	"strings"

	"github.com/ajangi/gAuthService/app/utils/db"
	"github.com/ajangi/gAuthService/app/utils/env"
	// using to connect to mysql database!
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
)
// UniqueInDB is a custom validation to check data is unique in database
func UniqueInDB(fl validator.FieldLevel) bool{
	database, err := db.GetDB()
	if err != nil {
		log.Panic(err)
	}
	param := strings.Split(fl.Param(), `:`)
	tableName := param[0]
	columnName := param[1]
	var count int
	query := database.Table(tableName).Where(columnName+" = ?", fl.Field().String())
	res := query.Count(&count)
	if res.Error != nil {
		errorString := res.Error.Error()
		if strings.Contains(errorString,getUniqInDbTableNotExistsError(tableName)) {
			log.Println("------------------")
			log.Println(res.Error.Error())
			log.Println("------------------")
		}
	}
	if count > 0 {
		return false
	}
	return true
}

func getUniqInDbTableNotExistsError(tableName string) string {
	return "Table '"+env.GetEnvVariable("DB_NAME")+"."+tableName+"' doesn't exist"
}