package commons

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ValidateBodyRequest(c *gin.Context, v interface{}) error {
	var err error
	if err = validator.Validate(v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return err
	}
	return nil
}
