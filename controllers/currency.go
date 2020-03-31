package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/services"
)

func GetCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {

		curr := c.Param("symbol")

		result := services.CurrencyService.GetCurrency(curr)

		c.JSON(http.StatusNotImplemented, "implement me")

	}
}

func GetAllCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusNotImplemented, "implement me")

	}
}
