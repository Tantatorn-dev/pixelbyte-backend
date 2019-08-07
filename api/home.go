package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home home for the api
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a Pixelbyte API",
	})
}
