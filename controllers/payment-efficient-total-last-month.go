package controllers

import (
	"net/http"
	"os"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentEfficientTotalLastMonth(c *gin.Context) {

	var payments struct {
		RekTotal int `json:"amount"`
	}

	period := os.Getenv("BILL_PERIOD")

	result := configs.DB.Raw("SELECT SUM(a.rek_total) AS rek_total FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor AND a.rek_sts = 1 AND a.rek_byr_sts > 0 AND b.byr_sts > 0 AND b.byr_tgl >= CONCAT(DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 1 MONTH) , '%Y-%m'), '-01 00:00:00') AND b.byr_tgl < CONCAT(DATE_FORMAT(CURDATE(), '%Y-%m'), '-01 00:00:00') AND a.rek_thn = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 + ? MONTH)) AND a.rek_bln = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 + ? MONTH))", period, period).Scan(&payments)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Total Efisiensi Penerimaan Bulan Lalu",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Total Efisiensi Penerimaan Bulan Lalu",
			Data:    payments,
		})
	}

}
