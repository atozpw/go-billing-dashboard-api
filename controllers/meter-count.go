package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func MeterCount(c *gin.Context) {

	var meters []struct {
		KpKode  string `json:"id"`
		SmTotal int    `json:"total"`
		SmDiisi int    `json:"fill"`
	}

	result := configs.DB.Raw("SELECT a.kp_kode, COUNT(*) AS sm_total, SUM(IF(b.sm_sts = 1, 1, 0)) AS sm_diisi FROM tm_pelanggan a JOIN tm_stand_meter b ON b.pel_no = a.pel_no AND b.sm_thn = YEAR(CURDATE()) AND b.sm_bln = MONTH(CURDATE()) AND b.sm_sts > 0 AND b.sm_sts < 3 AND a.kps_kode = 0 GROUP BY a.kp_kode").Scan(&meters)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Stand Meter",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Stand Meter",
			Data:    meters,
		})
	}

}
