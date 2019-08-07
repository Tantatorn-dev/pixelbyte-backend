package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status a status of each pixelbrite
type Status struct {
	ID       uint8 `json:"id"`
	IsOnline bool  `json:"is_online"`
}

// GetStatus Get a pixelbrite status
func GetStatus(c *gin.Context) {
	//test only
	var testStatus []Status
	testStatus = append(testStatus, Status{
		ID:       1,
		IsOnline: true,
	})
	testStatus = append(testStatus, Status{
		ID:       2,
		IsOnline: true,
	})

	c.JSON(http.StatusOK, testStatus)
}
