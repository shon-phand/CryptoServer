package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/CryptoServer/services"
	"github.com/shon-phand/CryptoServer/utils/errors"
)

func GetCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {
		//time.Sleep(6 * time.Second)
		symbol := strings.ToUpper(c.Param("symbol"))
		ok := ValidateSymbol(symbol)

		if !ok {
			c.JSON(http.StatusBadRequest, errors.StatusBadRequestError("invalid symbol, symbol should conatain letters only"))
			return
		}

		curr, err := services.CurrencyService.GetCurrency(symbol)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, curr)

	}
}

func GetAllCurrency() gin.HandlerFunc {

	return func(c *gin.Context) {

		res, err := services.CurrencyService.GetAllCurrency()
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, res)

	}
}

func UpdateDB() gin.HandlerFunc {
	return func(c *gin.Context) {

		tt, err := services.UpdateDatabase()

		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusCreated, tt)

	}

}

func ValidateSymbol(s string) bool {
	var reg = regexp.MustCompile(`^[A-Z]+$`).MatchString
	res := reg(s)
	return res
}
