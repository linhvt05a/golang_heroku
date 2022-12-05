package routers

import (
	"example/kenva-be/handlers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies([]string{"192.168.10.1"})
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/send-otp", handlers.SendOTP)
	router.POST("/verify-otp", handlers.VerifyOTP)
	router.POST("/create-password", handlers.UpdatePassword)
	router.POST("/forgot-password", handlers.ForgotPassword)
	router.Run("localhost:8088")
}
