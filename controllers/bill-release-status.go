package controllers

import (
	"net/http"
	"os"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func BillReleaseStatus(c *gin.Context) {

	var statuses []struct {
		KpKode string `json:"id"`
		RekJml int    `json:"count"`
	}

	period := os.Getenv("BILL_PERIOD")

	result := configs.DB.Raw("SELECT a.kp_kode, COUNT(*) AS rek_jml FROM tm_pelanggan a JOIN tm_drd_awal b ON b.pel_no = a.pel_no AND b.rek_thn = YEAR(DATE_ADD(CURDATE(), INTERVAL 1 - ? MONTH)) AND b.rek_bln = MONTH(DATE_ADD(CURDATE(), INTERVAL 1 - ? MONTH)) AND b.rek_sts = 1 AND a.kps_kode = 0 GROUP BY a.kp_kode", period, period).Scan(&statuses)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Status Penerbitan",
			Data:    statuses,
		})
	}

}
