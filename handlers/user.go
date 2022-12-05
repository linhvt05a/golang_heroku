package handlers

import (
	"example/kenva-be/commons"
	"example/kenva-be/configs"
	"example/kenva-be/db"
	"example/kenva-be/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	CLIENT = *commons.CreateTwilioAuth()
	DB     = *db.ConnectDB()
)

func CheckUserIsExist(phone string) (bool, int, error) {
	var user models.User
	fmt.Println("phone", phone)
	err := DB.Raw("SELECT phone_number,is_verified FROM users WHERE phone_number = ?", phone).Scan(&user).Error
	if err != nil {
		return false, 0, err
	}
	fmt.Println("user.PhoneNumber", user.PhoneNumber)
	fmt.Println("is_verified", user.IsVerified)
	return bool(user.IsVerified), len(user.PhoneNumber), nil
}

func Register(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.BindJSON(&bodyReq)
	isVerified, phoneLengh, err := CheckUserIsExist(bodyReq.PhoneNumber)
	fmt.Println("isVerified", isVerified)
	fmt.Println("phoneLenght", phoneLengh)
	fmt.Println("err", err)
	if err != nil {
		user := models.ErrorsResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		commons.Errors(c, http.StatusInternalServerError, user)
		return
	}
	if phoneLengh > 0 && !isVerified {
		user := models.User{Type: 0, Status: "Register", PhoneNumber: bodyReq.PhoneNumber}
		DB.Model(&user).Where("phone_number = ?", bodyReq.PhoneNumber).Updates(&user)
		commons.Response(c, http.StatusOK, user)
		return
	}
	if phoneLengh == 0 {
		keyHash := commons.CreateHashKey()
		encrypted := commons.EncryptPassword(bodyReq.Password, keyHash)
		user := models.User{
			Type:              0,
			Status:            strconv.Itoa(http.StatusOK),
			OTP:               bodyReq.OTP,
			Password:          encrypted,
			PhoneNumber:       bodyReq.PhoneNumber,
			KeyEncryptDecrypt: keyHash,
			IsVerified:        isVerified,
		}
		DB.Create(&user)
		commons.Response(c, http.StatusOK, user)
		return
	}
	if phoneLengh > 0 && isVerified {
		user := models.ErrorsResponse{
			StatusCode: http.StatusFound,
			Message:    "User already exists.Please back to login!",
		}
		commons.Errors(c, http.StatusFound, user)
		return
	}
	return
}

func CheckAccountLogin(phone, password string, c *gin.Context) bool {
	var user models.User
	err := DB.Raw("SELECT id, is_verified, phone_number,key_encrypt_decrypt,key_authorize, password FROM users WHERE phone_number = ?", phone).Scan(&user).Error
	if err != nil {
		commons.Errors(c, http.StatusInternalServerError, err.Error())
		return false
	}
	decrypted, _ := commons.DecryptPassword(user.KeyEncryptDecrypt, user.Password)
	fmt.Println("decrypted :", decrypted)
	fmt.Println("is_verified", user.IsVerified)
	if phone == user.PhoneNumber && password == decrypted && user.IsVerified {
		return true
	}
	if user.PhoneNumber != phone {
		errorRes := models.Response{
			Message:    "Tài khoản chưa tồn tại.Vui lòng đăng ký",
			StatusCode: http.StatusNotFound,
		}
		commons.Errors(c, http.StatusNotFound, &errorRes)
		return false
	}
	if phone == user.PhoneNumber && password != decrypted {
		errorRes := models.Response{
			Message:    "Sai tài khoản hoặc mật khẩu.",
			StatusCode: http.StatusBadRequest,
		}
		commons.Errors(c, http.StatusBadRequest, &errorRes)
		return false
	}
	return true
}

func Login(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.ShouldBindJSON(&bodyReq)
	fmt.Println("loggg phone", bodyReq.PhoneNumber)
	fmt.Println("loggg passs", bodyReq.Password)
	isUserExisted := CheckAccountLogin(bodyReq.PhoneNumber, bodyReq.Password, c)
	fmt.Println("isUserExisted", isUserExisted)
	if isUserExisted {
		user := models.User{Type: 2, Status: strconv.Itoa(http.StatusOK), PhoneNumber: bodyReq.PhoneNumber}
		DB.Model(&user).Where("phone_number = ?", bodyReq.PhoneNumber).Updates(&user)
		commons.Response(c, http.StatusOK, &user)
		return
	}
	return
}

func SendOTP(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.ShouldBindJSON(&bodyReq)
	params := &openapi.CreateVerificationParams{}
	params.SetTo(bodyReq.PhoneNumber)
	params.SetChannel("sms")
	configs, _ := configs.LoadConfig(".")
	fmt.Println("configs", configs.TwilioAccountSid)
	resp, err := CLIENT.VerifyV2.CreateVerification(configs.TwilioVerifySid, params)
	if err != nil {
		fmt.Println("Eroorrrr sent otp", err.Error())
		errorRes := models.Response{
			Message:    "InternalServerError",
			StatusCode: http.StatusInternalServerError,
		}
		commons.Errors(c, http.StatusInternalServerError, &errorRes)
		return
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		otpRes := models.OTPResponse{
			Sid:     *resp.Sid,
			Message: "OTP was sent!",
		}
		commons.Response(c, http.StatusOK, &otpRes)
		return
	}
}

func VerifyOTP(c *gin.Context) {
	var verifyOTP models.UserRequestInfo
	c.BindJSON(&verifyOTP)
	var user models.User
	params := &openapi.CreateVerificationCheckParams{}
	configs, _ := configs.LoadConfig(".")
	params.SetTo(verifyOTP.PhoneNumber)
	params.SetCode(verifyOTP.OTP)
	isVerified, phoneLengh, errCheckUser := CheckUserIsExist(verifyOTP.PhoneNumber)
	fmt.Println("check user exist", isVerified)
	fmt.Println("phoneLengh", phoneLengh)
	fmt.Println("errCheckUser", errCheckUser)
	resp, errVerifyOTP := CLIENT.VerifyV2.CreateVerificationCheck(configs.TwilioVerifySid, params)
	keyHash := commons.CreateHashKey()
	hashedPassword := commons.EncryptPassword(verifyOTP.Password, keyHash)
	errType := DB.Raw("SELECT id,type FROM users WHERE phone_number = ?", verifyOTP.PhoneNumber).Scan(&user).Error
	fmt.Print("access_token", user.ID)

	if errCheckUser != nil || errType != nil {
		errRes := models.ErrorsResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error !",
		}
		commons.Errors(c, http.StatusInternalServerError, errRes)
		return
	}

	if phoneLengh > 0 && user.Type == 3 && isVerified && *resp.Status == "approved" {
		res := models.User{
			Type:        3,
			IsVerified:  isVerified,
			Code:        http.StatusOK,
			PhoneNumber: verifyOTP.PhoneNumber,
			Message:     "User forgot password",
		}
		DB.Model(&res).Where("phone_number = ?", verifyOTP.PhoneNumber).Updates(&res)
		commons.Response(c, http.StatusOK, res)
		return
	}

	if phoneLengh > 0 && *resp.Status == "approved" && user.Type == 2 {
		fmt.Println("user.ID", user.ID)
		fmt.Println("configs.AccessTokenExpiresIn", configs.AccessTokenExpiresIn)
		fmt.Println("configs.AccessTokenPrivateKey", configs.AccessTokenPrivateKey)
		access_token, _ := commons.CreateToken(configs.AccessTokenExpiresIn, user, configs.AccessTokenPrivateKey)
		fmt.Println("access_token", access_token)
		res := models.User{
			Type:              2,
			IsVerified:        *resp.Status == "approved",
			Code:              http.StatusOK,
			PhoneNumber:       verifyOTP.PhoneNumber,
			Password:          hashedPassword,
			KeyEncryptDecrypt: keyHash,
			AccessToken:       access_token,
		}

		saveUser := models.User{
			Type:        2,
			IsVerified:  *resp.Status == "approved",
			Code:        http.StatusOK,
			PhoneNumber: verifyOTP.PhoneNumber,
			AccessToken: access_token,
		}
		commons.Response(c, http.StatusOK, saveUser)
		DB.Model(&res).Where("phone_number = ?", verifyOTP.PhoneNumber).Updates(&res)
		return
	}

	if phoneLengh == 0 {
		errorRes := models.ErrorsResponse{
			Message:    "User does not exist.Please back to register!",
			StatusCode: http.StatusBadRequest,
		}
		commons.Errors(c, http.StatusBadRequest, &errorRes)
		return
	}
	if errVerifyOTP != nil && *resp.Status != "approved" {
		errorRes := models.ErrorsResponse{
			Message:    "OTP invalid.Please try again! ",
			StatusCode: http.StatusBadRequest,
		}
		commons.Errors(c, http.StatusBadRequest, &errorRes)
		return
	}
}

func ForgotPassword(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	c.BindJSON(&bodyReq)
	fmt.Print("loggg", bodyReq.PhoneNumber)

	is_verified, phoneLengh, err := CheckUserIsExist(bodyReq.PhoneNumber)
	if err != nil {
		errorRes := models.ErrorsResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		commons.Errors(c, http.StatusInternalServerError, errorRes)
		return
	}
	if is_verified && phoneLengh > 0 {
		res := models.User{
			Type:    3,
			Message: "Update password",
			Code:    http.StatusOK,
		}
		DB.Model(&res).Where("phone_number = ?", bodyReq.PhoneNumber).Updates(&res)
		commons.Response(c, http.StatusOK, res)
	}
	return

}

func UpdatePassword(c *gin.Context) {
	var bodyReq models.UserRequestInfo
	var user models.User
	c.BindJSON(&bodyReq)
	fmt.Print("loggg", bodyReq.PhoneNumber)
	is_verified, phoneLengh, err := CheckUserIsExist(bodyReq.PhoneNumber)
	errGetPass := DB.Raw("SELECT id, is_verified, phone_number,key_encrypt_decrypt,key_authorize, password FROM users WHERE phone_number = ?", bodyReq.PhoneNumber).Scan(&user).Error
	decrypted, _ := commons.DecryptPassword(user.KeyEncryptDecrypt, user.Password)
	fmt.Println("decrypted :", decrypted)
	fmt.Println("is_verified", is_verified)
	fmt.Println("phoneLengh", phoneLengh)
	fmt.Println("err", err)
	if errGetPass != nil {
		user := models.ErrorsResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server",
		}
		commons.Errors(c, http.StatusInternalServerError, user)
		return
	}

	if is_verified {
		if bodyReq.OldPass == decrypted {
			key := commons.CreateHashKey()
			newHash := commons.EncryptPassword(bodyReq.NewPass, key)
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
				commons.Errors(c, http.StatusInternalServerError, &errorRes)
				return
			}
			commons.Response(c, http.StatusOK, &user)
			return
		}
		if bodyReq.NewPass == decrypted {
			errorRes := models.ErrorsResponse{
				Message:    "The new password must not be the same as the old password !",
				StatusCode: http.StatusNotFound,
			}
			commons.Errors(c, http.StatusNotFound, &errorRes)
			return
		}

	} else {
		errorRes := models.ErrorsResponse{
			Message:    "Old password is incorrect!",
			StatusCode: http.StatusNotFound,
		}
		commons.Errors(c, http.StatusNotFound, &errorRes)
		return
	}

}
