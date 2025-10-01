package main

import "time"

type User struct {
	FirstName   string    `form:"first_name" json:"first_name" binding:"required,min=2,max=16"`
	LastName    string    `form:"last_name" json:"last_name" binding:"required,min=2,max=16"`
	Username    string    `form:"username" json:"username" binding:"required,alphanum,min=8,max=32"`
	Email       string    `form:"email" json:"email" binding:"required,email,min=8,max=32"`
	PhoneNumber string    `form:"phone_number" json:"phone_number" binding:"required,number,min=11,max=11,startswith=09"`
	BirthDate   time.Time `form:"birth_date" json:"birth_date" time_format:"2006/01/02" binding:"required,valid_birth_date"`
	NationalID  string    `form:"national_id" json:"national_id" binding:"required,number,min=10,max=10,valid_national_id"`
}
