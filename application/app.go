package application

import (
	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/controllers"
)

var (
	r = gin.Default()
)

func StartApplication() {

	r.GET("/currency/:symbol", controllers.GetCurrency())
	r.GET("/currency", controllers.GetAllCurrency())

	r.Run(":8080")
}
