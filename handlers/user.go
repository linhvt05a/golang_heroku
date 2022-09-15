package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"example/kenva-be/commons"
	"example/kenva-be/configs"
	"example/kenva-be/db"
	"example/kenva-be/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	CLIENT             = *configs.CreateTwilioAuth()
	DB                 = *db.ConnectDB()
	VERIFY_SERVICE_SID = os.Getenv("VERIFY_SERVICE_SID")
)

func CheckUserIsExist(phone string) bool {
	var user models.User
	var exists bool
	err := DB.Model(&user).Select("count(*) > 0").Where("phone_number = ? AND is_verified = ?", phone, true).Find(&exists).Error
	if err != nil {
		return false
	}
	return exists
}
func Register(c *gin.Context) {
	bodyReq := models.UserRequestInfo{}
	c.ShouldBindJSON(&bodyReq)
	fmt.Printf("log phone '%s'\n", bodyReq.PhoneNumber)
	fmt.Printf("log pass '%s'\n", bodyReq.Password)
	isUserExist := CheckUserIsExist(bodyReq.PhoneNumber)
	if isUserExist {
		c.JSON(http.StatusFound, gin.H{"message": "User does existed. Please back to login!"})
		return
	} else {
		user := models.User{
			Type:        0,
			Password:    bodyReq.Password,
			PhoneNumber: bodyReq.PhoneNumber,
			IsVerified:  false,
			Message:     "Register",
		}
		c.JSON(http.StatusContinue, gin.H{"data": &user})
		return
	}
}

func SendOTP(c *gin.Context) {
	bodyReq := models.UserRequestInfo{}
	c.ShouldBindJSON(&bodyReq)
	params := &openapi.CreateVerificationParams{}
	params.SetTo(bodyReq.PhoneNumber)
	params.SetChannel("sms")
	resp, err := CLIENT.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		fmt.Println("Eroorrrr sent otp", err.Error())
		errorRes := models.Response{
			Message:    "InternalServerError",
			StatusCode: http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, gin.H{"errors": &errorRes})
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

func CheckAccountLogin(phone, password string, c *gin.Context) bool {
	var user models.User
	fmt.Println("loggggg ohine", phone)
	fmt.Println("loggggg pass", password)
	// Use GORM API build SQ
	body := models.UserRequestInfo{PhoneNumber: phone, Password: password}
	commons.ValidateBodyRequest(c, body)
	DB.Raw("SELECT id, phone_number,key_encrypt_decrypt,key_authorize, password FROM users WHERE phone_number = ?", phone).Scan(&user)

	decrypted := commons.DecryptPassword(user.KeyEncryptDecrypt, user.Password)
	fmt.Println("decrypted :", decrypted)
	if user.Password == decrypted && user.PhoneNumber != phone {
		errorRes := models.Response{
			Message:    "Sai tài khoản hoặc mật khẩu.",
			StatusCode: http.StatusNotFound,
		}
		c.JSON(http.StatusNotFound, gin.H{"errors": &errorRes})
		return false
	}
	if user.PhoneNumber != phone && user.Password != decrypted {
		errorRes := models.Response{
			Message:    "Tài khoản chưa tồn tại.Vui lòng đăng ký",
			StatusCode: http.StatusAccepted,
		}
		c.JSON(http.StatusAccepted, gin.H{"errors": &errorRes})
		return false
	}

	return password == decrypted && phone == user.PhoneNumber
}

func Login(c *gin.Context) {
	bodyReq := models.UserRequestInfo{}
	c.ShouldBindJSON(&bodyReq)
	fmt.Printf("log phone '%s'\n", bodyReq.PhoneNumber)
	fmt.Printf("log pass '%s'\n", bodyReq.Password)
	isUserExisted := CheckAccountLogin(bodyReq.PhoneNumber, bodyReq.Password, c)
	fmt.Println(isUserExisted)
	if isUserExisted {
		user := models.User{Type: 2, Message: "Loggin Success", Status: "200"}
		DB.Model(&user).Where("phone_number = ?", bodyReq.PhoneNumber).Updates(&user)
		c.JSON(http.StatusOK, gin.H{"data": &user})
		return
	} else {
		errorRes := models.Response{
			Message:    "Sai mật khẩu.Vui lòng thử lại.",
			StatusCode: http.StatusAccepted,
		}
		c.JSON(http.StatusAccepted, gin.H{"errors": &errorRes})
		return
	}
}

func VerifyOTP(c *gin.Context) {
	var verifyOTP models.UserRequestInfo
	c.BindJSON(&verifyOTP)
	fmt.Println("PHONE", verifyOTP.PhoneNumber)
	fmt.Println("OTP", verifyOTP.OTP)
	fmt.Println("PASSWORD", verifyOTP.Password)
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(verifyOTP.PhoneNumber)
	params.SetCode(verifyOTP.OTP)

	resp, err := CLIENT.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

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
		isUserExisted := CheckUserIsExist(verifyOTP.PhoneNumber)
		if isUserExisted {
			c.JSON(http.StatusFound, gin.H{"message": "User is existed !"})
			return
		} else {
			bytes := make([]byte, 16) //generate a random 32 byte key for AES-256
			if _, err := rand.Read(bytes); err != nil {
				panic(err.Error())
			}

			keyHass := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
			hashedPassword := commons.EncryptPassword(verifyOTP.Password, keyHass)
			fmt.Printf("key to encrypt/decrypt : %s\n", keyHass)
			fmt.Println("hashedPassword", hashedPassword)
			user := models.User{
				Status:            *resp.Status,
				Message:           "Tạo tài khoản thành công !",
				PhoneNumber:       verifyOTP.PhoneNumber,
				Password:          hashedPassword,
				OTP:               verifyOTP.OTP,
				IsVerified:        *resp.Status == "approved",
				KeyEncryptDecrypt: keyHass,
			}
			DB.Create(&user)
			c.JSON(http.StatusCreated, gin.H{"data": &user})
			return
		}
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
