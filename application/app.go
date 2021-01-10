package application

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/controllers"
	"github.com/shon-phand/CryptoServer/services"
)

var (
	r = gin.New()
)

func init() {
	fmt.Println("updating database, it will take 18 seconds to sync all data")
	_, err := services.UpdateDatabase()
	if err != nil {
		panic("error in synching data")
	}

	fmt.Println("update complete")
}

func StartApplication() {
	//gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", ShowIndex)
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



func ShowIndex(c *gin.Context) {
	c.JSON(200,gin.H {
		"message":"This is application 2",
	  })
}