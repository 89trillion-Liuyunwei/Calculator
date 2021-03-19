package route

import (
	"github.com/gin-gonic/gin"
	"max/api"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/calculation", api.Calculation)
	return r
}
