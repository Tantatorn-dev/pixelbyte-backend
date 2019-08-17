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

	// status of a Pixelbyte
	r.GET("/status", api.GetStatus)

	// Websocket endpoint
	r.GET("/ws", api.WSEndpoint)

	return r
}
