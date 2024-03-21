package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentToday(c *gin.Context) {

	var payments []struct {
		KpKode    string `json:"id"`
		RekLembar int    `json:"count"`
		RekTotal  int    `json:"amount"`
	}

	result := configs.DB.Raw("SELECT a.kp_kode, COUNT(*) AS rek_lembar, SUM(b.rek_total) AS rek_total FROM tm_pelanggan a JOIN tm_rekening b ON b.pel_no = a.pel_no AND b.rek_sts = 1 AND b.rek_byr_sts > 0 JOIN tm_pembayaran c ON c.rek_nomor = b.rek_nomor AND c.byr_sts > 0 AND c.byr_tgl >= CONCAT(CURDATE(), ' 00:00:00') GROUP BY a.kp_kode").Scan(&payments)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Penerimaan Harian",
			Data:    payments,
		})
	}

}
