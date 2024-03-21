package main

import (
	"github.com/atozpw/go-billing-dashboard-api/controllers"
	"github.com/atozpw/go-billing-dashboard-api/exceptions"
	"github.com/atozpw/go-billing-dashboard-api/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/access-token", controllers.AccessToken)
		v1.GET("/bill-last-months", middlewares.Auth, controllers.BillLastMonth)
		v1.GET("/bill-release-statuses", middlewares.Auth, controllers.BillReleaseStatus)
		v1.GET("/bill-this-months", middlewares.Auth, controllers.BillThisMonth)
		v1.GET("/branches", middlewares.Auth, controllers.Branch)
		v1.GET("/meter-counts", middlewares.Auth, controllers.MeterCount)
		v1.GET("/payment-effectives", middlewares.Auth, controllers.PaymentEffective)
		v1.GET("/payment-efficients", middlewares.Auth, controllers.PaymentEfficient)
		v1.GET("/payment-statuses", middlewares.Auth, controllers.PaymentStatus)
		v1.GET("/payment-this-months", middlewares.Auth, controllers.PaymentThisMonth)
		v1.GET("/payment-todays", middlewares.Auth, controllers.PaymentToday)
		v1.GET("/reading-this-months", middlewares.Auth, controllers.ReadingThisMonth)
		v1.GET("/water-usage-greater-tens", middlewares.Auth, controllers.WaterUsageGreaterTen)
		v1.GET("/water-usage-last-months", middlewares.Auth, controllers.WaterUsageLastMonth)
		v1.GET("/water-usage-one-to-tens", middlewares.Auth, controllers.WaterUsageOneToTen)
		v1.GET("/water-usage-this-months", middlewares.Auth, controllers.WaterUsageThisMonth)
		v1.GET("/water-usage-zeros", middlewares.Auth, controllers.WaterUsageZero)
	}

	router.NoRoute(exceptions.RouteException)

}
