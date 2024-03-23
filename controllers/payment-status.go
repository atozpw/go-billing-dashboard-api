package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentStatus(c *gin.Context) {

	var payments []struct {
		KpKode string `json:"id"`
		ByrJml int    `json:"count"`
	}

	result := configs.DB.Raw("SELECT c.kp_kode, COUNT(*) AS byr_jml FROM tm_pembayaran a JOIN tm_rekening b ON b.rek_nomor = a.rek_nomor AND a.byr_tgl >= DATE_SUB(NOW(), INTERVAL 1 HOUR) AND a.byr_sts > 0 JOIN tm_pelanggan c ON c.pel_no = b.pel_no GROUP BY c.kp_kode").Scan(&payments)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Penerimaan Status",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Penerimaan Status",
			Data:    payments,
		})
	}

}
