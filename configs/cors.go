package configs

import "github.com/gin-contrib/cors"

func Cors() cors.Config {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5173"}

	return config

}
