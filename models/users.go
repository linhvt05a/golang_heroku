package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string `json:"phonenumber"`
	OTP         string `json:"otp"`
	Password    string `json:"password"`
	IsVerified  bool   `json:"isVerified"`
	Message     string `json:"message"`
	Status      string `json:"code"`
}

type SendOTP struct {
	PhoneNumber string `json:"phonenumber"`
	OTP         string `json:"otp"`
	Password    string `json:"password" validate:"min=1,max=16"`
}

type OTPResponse struct {
	Sid     string `json:"Sid"`
	Message string `json:"message"`
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}
