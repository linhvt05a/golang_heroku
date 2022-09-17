package commons

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": &data})
	return
}
func Errors(c *gin.Context, statusCode int, error interface{}) {
	c.JSON(statusCode, gin.H{"errors": &error})
	return
}
