package apiserver

import (
	"github.com/gin-gonic/gin"
)

var getHealthz = func(c *gin.Context) {
	if !initialized {
		c.JSON(503, map[string]string{"status": "initializing"})
	} else if applicationIsHealthy {
		c.JSON(200, map[string]string{"status": "up"})
	} else {
		c.JSON(500, map[string]string{"status": "down"})
	}
}

var getMetrics = func(c *gin.Context) {
	c.String(200, "namespace.value 0\n")
}
