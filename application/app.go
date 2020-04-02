package application

import (
	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/controllers"
)

var (
	r = gin.Default()
)

func StartApplication() {
	gin.SetMode(gin.ReleaseMode)
	r.GET("/currency/", controllers.GetAllCurrency())
	r.GET("/currency/:symbol", controllers.GetCurrency())
	r.GET("/internal/currency/UpdateDb", controllers.UpdateDB())
	//go services.UpdateDatabase()
	gin.SetMode(gin.ReleaseMode)
	r.Run(":8080")
}
