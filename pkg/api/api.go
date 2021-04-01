package api

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/maxtv_buildings"
	"maxtv_middleware/pkg/pythonReporter"
	"os"
	"strings"
)

func Processing() {

	r := gin.Default()

	auth := r.Group("/")
	auth.Use(TokenAuthMiddleware())
	{
		auth.GET("/python-reporter/is-paid-maxtv-building", pythonReporter.IsPaidMaxtvBuilding)
		auth.GET("/maxtv-buildings", maxtv_buildings.GetMaxTvBuildings)
	}

	r.Run(":" + os.Getenv("PORT"))
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("AUTH_TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		common.Log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		//token := c.Request.FormValue("api_token")

		if len(c.Request.Header["Authorization"]) == 0 {
			common.RespondWithError(c, 401, "Auth header required")
			return
		}

		if c.Request.Header["Authorization"][0] == "" {
			common.RespondWithError(c, 401, "API token required")
			return
		}

		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

		if token == "" {
			common.RespondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			common.RespondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}
