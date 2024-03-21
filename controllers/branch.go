package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
)

func Branch(c *gin.Context) {

	var branches []struct {
		KpKode string `json:"id"`
		KpKet  string `json:"name"`
	}

	result := configs.DB.Raw("SELECT kp_kode, kp_ket FROM tr_kota_pelayanan").Scan(&branches)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Kota Pelayanan",
			Data:    branches,
		})
	}

}
