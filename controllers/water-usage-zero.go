package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func WaterUsageZero(c *gin.Context) {

	var waters []struct {
		KpKode     string `json:"id"`
		SmJml      int    `json:"count"`
		RekUangair int    `json:"amount"`
	}

	result := configs.DB.Raw("SELECT a.kp_kode, COUNT(*) AS sm_jml, SUM(getUangAir(a.gol_kode, MONTH(CURDATE()), YEAR(CURDATE()), 0)) AS rek_uangair FROM tm_pelanggan a JOIN tm_stand_meter b ON b.pel_no = a.pel_no AND b.sm_thn = YEAR(CURDATE()) AND b.sm_bln = MONTH(CURDATE()) AND b.sm_kini = b.sm_lalu AND b.sm_sts = 1 AND a.kps_kode = 0 GROUP BY a.kp_kode").Scan(&waters)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pemakaian Air 0",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pemakaian Air 0",
			Data:    waters,
		})
	}

}
