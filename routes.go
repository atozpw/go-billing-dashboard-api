package main

import (
	"github.com/atozpw/go-billing-dashboard-api/controllers"
	"github.com/atozpw/go-billing-dashboard-api/exceptions"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.GET("/access-token", controllers.AccessToken)
		v1.GET("/bill-last-months", controllers.BillLastMonth)
		v1.GET("/bill-release-statuses", controllers.BillReleaseStatus)
		v1.GET("/bill-this-months", controllers.BillThisMonth)
		v1.GET("/branches", controllers.Branch)
		v1.GET("/meter-counts", controllers.MeterCount)
		v1.GET("/payment-effectives", controllers.PaymentEffective)
		v1.GET("/payment-efficients", controllers.PaymentEfficient)
		v1.GET("/payment-statuses", controllers.PaymentStatus)
		v1.GET("/payment-this-months", controllers.PaymentThisMonth)
		v1.GET("/payment-todays", controllers.PaymentToday)
		v1.GET("/reading-this-months", controllers.ReadingThisMonth)
		v1.GET("/water-usage-greater-tens", controllers.WaterUsageGreaterTen)
		v1.GET("/water-usage-last-months", controllers.WaterUsageLastMonth)
		v1.GET("/water-usage-one-to-tens", controllers.WaterUsageOneToTen)
		v1.GET("/water-usage-this-months", controllers.WaterUsageThisMonth)
		v1.GET("/water-usage-zeros", controllers.WaterUsageZero)
	}

	router.NoRoute(exceptions.RouteException)

}
