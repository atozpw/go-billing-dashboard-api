package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentThisMonth(c *gin.Context) {

	var payments []struct {
		ByrTgl   string `json:"date"`
		ByrTotal int    `json:"amount"`
	}

	result := configs.DB.Raw("SELECT DATE_FORMAT(byr_tgl, '%Y-%m-%d') AS byr_tgl, SUM(byr_total) AS byr_total FROM tm_pembayaran WHERE byr_tgl >= CONCAT(DATE_FORMAT(CURDATE(), '%Y-%m'), '-01 00:00:00') AND byr_sts > 0 GROUP BY DATE_FORMAT(byr_tgl, '%Y-%m-%d')").Scan(&payments)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Penerimaan Bulanan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Penerimaan Bulanan",
			Data:    payments,
		})
	}

}
