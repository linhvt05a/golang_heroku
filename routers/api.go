package routers

import (
	"example/kenva-be/handlers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/sendOTP", handlers.SendOTP)
	router.POST("/verifyOTP", handlers.VerifyOTP)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
