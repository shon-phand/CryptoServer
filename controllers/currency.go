package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/services"
)

func GetCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {

		curr := c.Param("symbol")

		result, err := services.CurrencyService.GetCurrency(curr)

		if err != nil {
			c.JSON(err.Status, err)
		}
		fmt.Println("res in controller", result)
		c.JSON(http.StatusOK, result)

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
