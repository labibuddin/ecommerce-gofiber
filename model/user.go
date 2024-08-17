package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string `gorm:"uniqueIndex;not null" json:"username"`
	Email    	string `gorm:"uniqueIndex;not null" json:"email"`
	Password 	string `gorm:"not null" json:"password"`
	Names    	string `json:"names"`
}

type UserSwagger struct {
	Username	string `json:"username" example:"Alfin Afifi"`
	Email    	string `json:"email" example:"pakvincent@gmail.com"`
	Password 	string `json:"password" example:"12345678"`
	Names    	string `json:"names" example:"Labibuddin"`
}

type UserInputSwagger struct {
	Identity string `json:"identity" example:"Alfin Afifi"`
	Password string `json:"password" example:"12345678"`
}

type UserUpdateSwagger struct{
	Names    	string `json:"names" example:"Udin Labib"`
}
type UserDeleteSwagger struct{
	Password 	string `json:"password" example:"12345678"`
}