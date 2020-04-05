package application

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/controllers"
)

var (
	r = gin.New()
)

func StartApplication() {
	//gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/currency/", controllers.GetAllCurrency())
	r.GET("/currency/:symbol", controllers.GetCurrency())
	r.GET("/internal/currency/UpdateDb", controllers.UpdateDB())
	//go services.UpdateDatabase()
	gin.SetMode(gin.ReleaseMode)
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 20 * time.Second,
		ErrorLog:     nil,
	}

	if err := server1.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
