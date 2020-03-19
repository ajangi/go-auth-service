package config

import "github.com/ajangi/gAuthService/app/utils/env"
// DbConfig is the config of database connection in gorm
var DbConfig  = func() string{
	var databaseHost       = env.GetEnvVariable("DB_Host")
	var databasePort       = env.GetEnvVariable("DB_PORT")
	var databaseName       = env.GetEnvVariable("DB_NAME")
	var databaseUser       = env.GetEnvVariable("DB_USER")
	var databasePassword   = env.GetEnvVariable("DB_PASSWORD")
	return databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+ databaseName +"?charset=utf8&parseTime=True&loc=Local"
}
