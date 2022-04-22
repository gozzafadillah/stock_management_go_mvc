package config

import (
	"gozzafadillah/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	InitDB()
}

var DB *gorm.DB
var username = "root"
var password = ""

var host = "host.docker.internal"

// var host = "127.0.0.1"
var port = "3306"
var dbName = "db_gudang"
var dbTest = "db_gudang_test"

func InitDB() {
	var err error
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitDBTest() {
	var err error
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbTest + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrateTest()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, models.Product{}, models.Product_Status{})
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{}, models.Product{}, models.Product_Status{})
	DB.AutoMigrate(&models.User{}, models.Product{}, models.Product_Status{})
}
