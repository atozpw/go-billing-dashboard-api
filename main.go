package main

import (
	"os"

	"github.com/atozpw/go-billing-dashboard-api/configs"
	"github.com/atozpw/go-billing-dashboard-api/helpers"
	"github.com/atozpw/go-billing-dashboard-api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	configs.LoadEnvironment()
	configs.ConnectDatabase()
	configs.Logging()
	configs.Mode()

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(helpers.LoggerFormatter))
	router.Use(gin.Recovery())
	router.Use(middlewares.Timeout())
	router.Use(cors.New(configs.Cors()))
	Routes(router)
	router.Run(os.Getenv("APP_URL"))

}
