package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/configs"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		cors.New(cors.Config{
			AllowOrigins:     configs.Configurations.CorsConfig.AllowOrigins,
			AllowMethods:     configs.Configurations.CorsConfig.AllowMethods,
			AllowHeaders:     configs.Configurations.CorsConfig.AllowHeaders,
			AllowCredentials: configs.Configurations.CorsConfig.AllowCredentials,
		})(c)

		c.Next()
	}
}
