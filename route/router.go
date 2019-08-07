package route

import (
	"pixelbyte-backend/api"

	"github.com/gin-gonic/gin"
)

// Init init routes for the API
func Init() *gin.Engine {
	r := gin.Default()
	// home
	r.GET("/", api.Home)

	return r
}
