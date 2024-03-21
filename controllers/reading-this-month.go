package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func ReadingThisMonth(c *gin.Context) {

	var readings []struct {
		SmTgl   string `json:"date"`
		SmTotal int    `json:"count"`
	}

	result := configs.DB.Raw("SELECT DATE_FORMAT(sm_tgl, '%Y-%m-%d') AS sm_tgl, COUNT(*) AS sm_total FROM tm_stand_meter WHERE sm_tgl >= CONCAT(DATE_FORMAT(CURDATE(), '%Y-%m'), '-01 00:00:00') AND sm_sts = 1 GROUP BY DATE_FORMAT(sm_tgl, '%Y-%m-%d')").Scan(&readings)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Baca Meter Bulanan",
			Data:    readings,
		})
	}

}
