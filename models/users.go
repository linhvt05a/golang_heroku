package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber       string `json:"phonenumber"`
	OTP               string `json:"otp"`
	Password          string `json:"password"`
	IsVerified        bool   `json:"isVerified"`
	Message           string `json:"message"`
	Status            string `json:"status"`
	KeyEncryptDecrypt string `json:"keyEncryptDecrypt"`
	KeyAuthorize      string `json:"keyAuthorize"`
	Type              int    `json:"type"`
	Code              int    `json:"code"`
}

type UserRequestInfo struct {
	PhoneNumber string `json:"phonenumber"`
	OTP         string `json:"otp"`
	Password    string `json:"password" validate:"min=8"`
}

type OTPResponse struct {
	Sid     string `json:"Sid"`
	Message string `json:"message"`
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
	Type       int    `json:"type"`
}

type ErrorsResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}
