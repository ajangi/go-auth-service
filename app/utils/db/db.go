package db

import (
	"github.com/ajangi/gAuthService/app/config"
	"github.com/jinzhu/gorm"
)

type fn func()
// GetDB is a method to initiate a database connection
func GetDB() (db *gorm.DB, err error) {
	return gorm.Open("mysql", config.DbConfig())
}

// MigrateTable is a function to migrate a table
func MigrateTable(tableName string, migrate fn)  {
	database, err := GetDB()
	if err != nil {
		return
	}
	if !database.HasTable(tableName) {
		migrate()
	}
}