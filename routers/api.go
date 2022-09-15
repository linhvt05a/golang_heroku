package routers

import (
	"example/kenva-be/handlers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/send-otp", handlers.SendOTP)
	router.POST("/verify-otp", handlers.VerifyOTP)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)
	router.Run("localhost:8088")
}
