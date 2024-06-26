package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func WaterUsageThisMonth(c *gin.Context) {

	var waters []struct {
		KpKode     string `json:"id"`
		SmPakai    int    `json:"usage"`
		RekUangair int    `json:"amount"`
	}

	result := configs.DB.Raw("SELECT a.kp_kode, SUM((b.sm_kini - b.sm_lalu)) AS sm_pakai, SUM(getUangAir(a.gol_kode, MONTH(CURDATE()), YEAR(CURDATE()), (b.sm_kini - b.sm_lalu))) AS rek_uangair FROM tm_pelanggan a JOIN tm_stand_meter b ON b.pel_no = a.pel_no AND b.sm_thn = YEAR(CURDATE()) AND b.sm_bln = MONTH(CURDATE()) AND (b.sm_kini - b.sm_lalu) >= 0 AND b.sm_sts = 1 AND a.kps_kode = 0 GROUP BY a.kp_kode").Scan(&waters)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pemakaian Air Bulan Berjalan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pemakaian Air Bulan Berjalan",
			Data:    waters,
		})
	}

}
