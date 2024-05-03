package models

import "gorm.io/gorm"

type users struct {
	gorm.Model
	role     string `gorm:"not null ; default:user"`
	username string
	email    string
	password string
}
type errors struct {
	NameError     string
	EmailError    string
	PasswordError string
	RoleError     string
	CommonError   string
}
type compare struct {
	Password string
	Role     string
	UserName string
}
type UserDetails struct {
	UserName string
	Email    string
}
