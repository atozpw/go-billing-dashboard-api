package controllers

import (
	"net/http"
	"os"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func BillLastMonth(c *gin.Context) {

	var bills []struct {
		KpKode    string `json:"id"`
		RekLembar int    `json:"count"`
		RekTotal  int    `json:"amount"`
	}

	period := os.Getenv("BILL_PERIOD")

	result := configs.DB.Raw("SELECT a.kp_kode, COUNT(*) AS rek_lembar, SUM(b.rek_total) AS rek_total FROM tm_pelanggan a JOIN tm_rekening b ON b.pel_no = a.pel_no AND b.rek_thn = YEAR(DATE_SUB(CURDATE(), INTERVAL 1 + ? MONTH)) AND b.rek_bln = MONTH(DATE_SUB(CURDATE(), INTERVAL 1 + ? MONTH)) AND b.rek_sts = 1 GROUP BY a.kp_kode", period, period).Scan(&bills)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Tagihan Bulan Lalu",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Tagihan Bulan Lalu",
			Data:    bills,
		})
	}

}
