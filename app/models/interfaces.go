package models

type modelInterface interface {
	migrate()
	TableName() string
}
