package handlers

import (
	"example/kenva-be/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

func CreateTwilioAuth() *twilio.RestClient {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return client
}

func SendOTP(c *gin.Context) {
	var phone models.SendOTP
	c.BindJSON(&phone)
	fmt.Printf("log phone '%s'\n", phone.PhoneNumber)
	client := CreateTwilioAuth()
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone.PhoneNumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println("Eroorrrr sent otp", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})

	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "OTP was sent!", "vsid": *resp.Sid})
	}
}

func VerifyOTP(c *gin.Context) {
	client := CreateTwilioAuth()
	var verifyOTP models.SendOTP

	c.BindJSON(&verifyOTP)
	fmt.Println("PHONE", verifyOTP.PhoneNumber)
	fmt.Println("OTP", verifyOTP.OTP)
	// fmt.Scanln(&verifyOTP.OTP)

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(verifyOTP.PhoneNumber)
	params.SetCode(verifyOTP.OTP)

	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": *resp.Status})
	} else {
		fmt.Println("Incorrect!")
		c.IndentedJSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Incorrect OTP !"})
	}
}
