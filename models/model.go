package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Name string
	Password string
	Login Login `gorm:"foreignKey:SessionId"`
	Posts []Post `gorm:"foreignKey:UserRefer"`
}

type Post struct {
	gorm.Model
	Title string
	Description string
	UserRefer uint
}

type Login struct{
	HashedPassword string
	SessionToken string
	CSRFToken string
	SessionId uint
}