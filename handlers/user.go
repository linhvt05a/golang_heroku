package handlers

import (
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
	GIN                *gin.Context
)

func CheckUserIsExist(phone string) bool {
	var user models.User
	fmt.Println("phone", phone)
	DB.Raw("SELECT phone_number FROM users WHERE phone_number = ?", phone).Scan(&user)
	fmt.Println("user.PhoneNumber", user.PhoneNumber)
	return len(user.PhoneNumber) > 0
}

func Register(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.BindJSON(&bodyReq)
	isUserExist := CheckUserIsExist(bodyReq.PhoneNumber)
	fmt.Println("check user existed", isUserExist)

	if isUserExist {
		user := models.ErrorsResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "User does existed. Please login",
		}
		commons.Errors(c, http.StatusBadRequest, user)
	} else {
		keyHash := commons.CreateHashKey()
		encrypted := commons.EncryptPassword(bodyReq.Password, keyHash)
		user := models.User{
			Type:              1,
			Message:           "Register",
			OTP:               bodyReq.OTP,
			Password:          encrypted,
			PhoneNumber:       bodyReq.PhoneNumber,
			KeyEncryptDecrypt: keyHash,
		}
		DB.Create(&user)
		commons.Response(c, http.StatusOK, user)
	}
	return
}

func CheckAccountLogin(phone, password string, c *gin.Context) bool {
	var user models.User
	// body := models.UserRequestInfo{PhoneNumber: phone, Password: password}
	// commons.ValidateBodyRequest(c, body)

	err := DB.Raw("SELECT id, phone_number,key_encrypt_decrypt,key_authorize, password FROM users WHERE phone_number = ?", phone).Scan(&user).Error
	if err != nil {
		commons.Errors(c, http.StatusInternalServerError, err.Error())
	}
	if phone == user.PhoneNumber {
		decrypted, err := commons.DecryptPassword(user.KeyEncryptDecrypt, user.Password)
		if err != nil {
			commons.Errors(c, http.StatusBadRequest, err.Error())
		}
		fmt.Println("decrypted :", decrypted)

		if password != decrypted && user.PhoneNumber == phone {
			errorRes := models.Response{
				Message:    "Sai tài khoản hoặc mật khẩu.",
				StatusCode: http.StatusBadRequest,
			}
			commons.Errors(c, http.StatusBadRequest, &errorRes)
			return false
		}
		if user.PhoneNumber != phone && user.Password != decrypted {
			errorRes := models.Response{
				Message:    "Tài khoản chưa tồn tại.Vui lòng đăng ký",
				StatusCode: http.StatusNotFound,
			}
			commons.Errors(c, http.StatusNotFound, &errorRes)
			return false
		}

		if phone == user.PhoneNumber && password == decrypted {
			return true
		}
	}

	return false
}

func Login(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.ShouldBindJSON(&bodyReq)
	fmt.Println("loggg phone", bodyReq.PhoneNumber)
	fmt.Println("loggg passs", bodyReq.Password)
	isUserExisted := CheckAccountLogin(bodyReq.PhoneNumber, bodyReq.Password, c)
	fmt.Println("isUserExisted", isUserExisted)
	if isUserExisted {
		user := models.User{Type: 2, Message: "Loggin Success", Status: "200"}
		DB.Model(&user).Where("phone_number = ?", bodyReq.PhoneNumber).Updates(&user)
		commons.Response(c, http.StatusFound, &user)
	} else {
		errorRes := models.Response{
			Message:    "Tài khoản chưa tồn tại trên hệ thống.",
			StatusCode: http.StatusBadRequest,
		}
		commons.Errors(c, http.StatusBadRequest, &errorRes)
	}
}

func SendOTP(c *gin.Context) {
	var bodyReq models.UserRequestInfo
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
		commons.Errors(c, http.StatusInternalServerError, &errorRes)
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		otpRes := models.OTPResponse{
			Sid:     *resp.Sid,
			Message: "OTP was sent!",
		}
		commons.Response(c, http.StatusOK, &otpRes)
	}
}

func VerifyOTP(c *gin.Context) {
	var verifyOTP models.UserRequestInfo
	var user models.User
	c.BindJSON(&verifyOTP)
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(verifyOTP.PhoneNumber)
	params.SetCode(verifyOTP.OTP)

	resp, err := CLIENT.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)
	keyHash := commons.CreateHashKey()
	hashedPassword := commons.EncryptPassword(verifyOTP.Password, keyHash)
	if err != nil {
		fmt.Println(err.Error())
		errorRes := models.Response{
			Message:    "Lỗi gửi mã xác nhận.Xin vui lòng thử lại.",
			StatusCode: http.StatusAccepted,
		}
		commons.Errors(c, http.StatusAccepted, &errorRes)
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		isUserExisted := CheckUserIsExist(verifyOTP.PhoneNumber)
		if isUserExisted {
			userRes := models.User{
				Message:           *resp.Status,
				Type:              2,
				IsVerified:        *resp.Status == "approved",
				PhoneNumber:       verifyOTP.PhoneNumber,
				Password:          hashedPassword,
				KeyEncryptDecrypt: keyHash,
				OTP:               verifyOTP.OTP,
			}
			err := DB.Model(&user).Where("phone_number = ?", verifyOTP.PhoneNumber).Updates(&userRes).Error
			if err != nil {
				errorRes := models.ErrorsResponse{
					Message:    "Internal server",
					StatusCode: http.StatusInternalServerError,
				}
				commons.Errors(c, http.StatusOK, &errorRes)
				return
			}
			commons.Response(c, http.StatusOK, &userRes)
		} else {
			fmt.Printf("key to encrypt/decrypt : %s\n", keyHash)
			fmt.Println("hashedPassword", hashedPassword)
			user := models.User{
				Status:            *resp.Status,
				Message:           "Tạo tài khoản thành công !",
				PhoneNumber:       verifyOTP.PhoneNumber,
				Password:          hashedPassword,
				OTP:               verifyOTP.OTP,
				IsVerified:        *resp.Status == "approved",
				KeyEncryptDecrypt: keyHash,
				Type:              2,
			}
			DB.Create(&user)
			commons.Response(c, http.StatusCreated, &user)
		}
	} else {
		fmt.Println("Incorrect!")
		errorRes := models.Response{
			Message:    "Mã xác nhận không đúng.Xin vui lòng thử lại.",
			StatusCode: http.StatusFound,
		}
		commons.Errors(c, http.StatusOK, &errorRes)
	}
}

func ForgotPassword(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	var user models.User
	c.BindJSON(&bodyReq)
	fmt.Print("loggg", bodyReq.PhoneNumber)

	isUserExist := CheckUserIsExist(user.PhoneNumber)
	if isUserExist {
		user := models.User{
			Type:    3,
			Message: "Forgot password",
		}
		commons.Response(c, http.StatusOK, &user)
	}
}

func UpdatePassword(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.BindJSON(&bodyReq)
	fmt.Print("loggg", bodyReq.PhoneNumber)

	isUserExist := CheckUserIsExist(bodyReq.PhoneNumber)
	fmt.Println("isUserExist", isUserExist)
	if isUserExist {
		key := commons.CreateHashKey()
		newHash := commons.EncryptPassword(bodyReq.Password, key)
		user := models.User{
			Message:           "Cập nhật mật khẩu thành công !",
			PhoneNumber:       bodyReq.PhoneNumber,
			Password:          newHash,
			KeyEncryptDecrypt: key,
		}
		err := DB.Model(&user).Where("phone_number = ?", bodyReq.PhoneNumber).Updates(&user).Error
		if err != nil {
			errorRes := models.ErrorsResponse{
				Message:    "Internal server",
				StatusCode: http.StatusInternalServerError,
			}
			commons.Errors(c, http.StatusOK, &errorRes)
			return
		}
		commons.Response(c, http.StatusOK, &user)
	} else {
		errorRes := models.ErrorsResponse{
			Message:    "Số điện thoại không tồn tại.Vui lòng đăng ký để sử dụng.",
			StatusCode: http.StatusNotFound,
		}
		commons.Errors(c, http.StatusNotFound, &errorRes)

	}
	return

}
