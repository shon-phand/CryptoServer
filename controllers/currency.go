package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/services"
)

func GetCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {

		symbol := strings.ToUpper(c.Param("symbol"))

		curr, err := services.CurrencyService.GetCurrency(symbol)
		if err != nil {
			c.JSON(err.Status, err)
		}
		c.JSON(http.StatusOK, curr)

	}
}

func GetAllCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {
		res, err := services.CurrencyService.GetAllCurrency()
		if err != nil {
			c.JSON(err.Status, err)
		}
		c.JSON(http.StatusOK, res)

	}
}

func UpdateDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("started updating DB")
		go services.UpdateDatabase()
		fmt.Println("statement executed")
		c.JSON(http.StatusCreated, "it will take time to update the DB")
	}

}
