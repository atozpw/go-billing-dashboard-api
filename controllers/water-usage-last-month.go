package controllers

import (
	"net/http"
	"os"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func WaterUsageLastMonth(c *gin.Context) {

	var waters []struct {
		KpKode     string `json:"id"`
		RekPakai   int    `json:"usage"`
		RekUangair int    `json:"amount"`
	}

	period := os.Getenv("BILL_PERIOD")

	result := configs.DB.Raw("SELECT a.kp_kode, SUM((b.rek_stankini - b.rek_stanlalu)) AS rek_pakai, SUM(b.rek_uangair) AS rek_uangair FROM tm_pelanggan a JOIN tm_rekening b ON b.pel_no = a.pel_no AND b.rek_thn = YEAR(DATE_SUB(CURDATE(), INTERVAL ? MONTH)) AND b.rek_bln = MONTH(DATE_SUB(CURDATE(), INTERVAL ? MONTH)) AND b.rek_sts = 1 GROUP BY a.kp_kode", period, period).Scan(&waters)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pemakaian Air Bulan Lalu",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pemakaian Air Bulan Lalu",
			Data:    waters,
		})
	}

}
