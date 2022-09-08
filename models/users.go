package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type SendOTP struct {
	PhoneNumber string `json:"phonenumber"`
	OTP         string `json:"otp"`
}
