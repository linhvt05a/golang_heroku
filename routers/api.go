package routers

import (
	"example/kenva-be/handlers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/albums", handlers.CreateUser)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
