package models

import (
	"fmt"

	"github.com/GemaSatya/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(){

	connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", env.ReadEnv("DB_USERNAME"), env.ReadEnv("DB_PASSWORD"), env.ReadEnv("DB_HOST"), env.ReadEnv("DB_PORT"), env.ReadEnv("DB_NAME"))

	database, err := gorm.Open(mysql.Open(connectionUrl))
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&User{}, &Post{}, &Login{})

	DB = database

}