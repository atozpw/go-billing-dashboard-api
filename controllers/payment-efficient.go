package controllers

import (
	"net/http"
	"os"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentEfficient(c *gin.Context) {

	var payments []struct {
		KpKode   string `json:"id"`
		RekTotal int    `json:"amount"`
	}

	period := os.Getenv("BILL_PERIOD")

	result := configs.DB.Raw("SELECT a.kp_kode, SUM(b.rek_total) AS rek_total FROM tm_pelanggan a JOIN tm_rekening b ON b.pel_no = a.pel_no AND b.rek_thn = YEAR(DATE_SUB(CURDATE(), INTERVAL ? MONTH)) AND b.rek_bln = MONTH(DATE_SUB(CURDATE(), INTERVAL ? MONTH)) AND b.rek_sts = 1 AND b.rek_byr_sts > 0 GROUP BY a.kp_kode", period, period).Scan(&payments)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Efisiensi Penerimaan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Efisiensi Penerimaan",
			Data:    payments,
		})
	}

}
