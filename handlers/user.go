package handlers

import (
	"example/kenva-be/commons"
	"example/kenva-be/db"
	"example/kenva-be/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"
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
	phone := models.SendOTP{}
	c.ShouldBindJSON(&phone)
	fmt.Printf("log phone '%s'\n", phone.PhoneNumber)
	fmt.Printf("log pass '%s'\n", phone.Password)
	client := CreateTwilioAuth()
	params := &openapi.CreateVerificationParams{}
	formatPhone := phone.PhoneNumber
	params.SetTo(formatPhone)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)

	errValidate := commons.ValidateBodyRequest(c, phone)

	if errValidate == nil {
		if err != nil {
			fmt.Println("Eroorrrr sent otp", err.Error())
			errorRes := models.Response{
				Message:    "Lỗi gửi mã xác nhận.Xin vui lòng thử lại.",
				StatusCode: http.StatusBadRequest,
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": &errorRes})
			return
		} else {
			fmt.Printf("Sent verification '%s'\n", *resp.Sid)
			otpRes := models.OTPResponse{
				Sid:     *resp.Sid,
				Message: "OTP was sent!",
			}
			c.JSON(http.StatusOK, gin.H{"data": &otpRes})
			return
		}
	}

}

func VerifyOTP(c *gin.Context) {
	client := CreateTwilioAuth()
	db := db.ConnectDB()
	var verifyOTP models.SendOTP
	c.BindJSON(&verifyOTP)
	fmt.Println("PHONE", verifyOTP.PhoneNumber)
	fmt.Println("OTP", verifyOTP.OTP)

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(verifyOTP.PhoneNumber)
	params.SetCode(verifyOTP.OTP)

	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println(err.Error())
		errorRes := models.Response{
			Message:    "Lỗi gửi mã xác nhận.Xin vui lòng thử lại.",
			StatusCode: http.StatusAccepted,
		}
		c.JSON(http.StatusAccepted, gin.H{"data": &errorRes})
		return
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		password := []byte(verifyOTP.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(hashedPassword))

		user := models.User{
			Status:      *resp.Status,
			Message:     "Xác thực thành công!",
			PhoneNumber: verifyOTP.PhoneNumber,
			Password:    string(hashedPassword),
			OTP:         verifyOTP.OTP,
			IsVerified:  *resp.Status == "approved",
		}

		db.Create(&user)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": &user})
		return
	} else {
		fmt.Println("Incorrect!")
		errorRes := models.Response{
			Message:    "Mã xác nhận không đúng.Xin vui lòng thử lại.",
			StatusCode: http.StatusFound,
		}
		c.JSON(http.StatusFound, gin.H{"errors": &errorRes})
		return
	}
}
