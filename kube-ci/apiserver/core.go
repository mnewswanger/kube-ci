package apiserver

import (
	"github.com/gin-gonic/gin"
)

var getHealthz = func(c *gin.Context) {
	if applicationIsHealthy {
		c.String(200, `{"status": "up"}`)
	} else {
		c.String(500, `{"status": "down"}`)
	}
}
