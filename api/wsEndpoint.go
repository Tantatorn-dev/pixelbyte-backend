package api

import (
	"pixelbyte-backend/ws"

	"github.com/gin-gonic/gin"
)

// WSEndpoint api endpoint for Websocket
func WSEndpoint(c *gin.Context) {
	ws.Handler(c.Writer, c.Request)
}
