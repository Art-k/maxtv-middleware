package api

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/demografics"
	"maxtv_middleware/pkg/maxtv_buildings"
	"maxtv_middleware/pkg/maxtv_companies"
	"maxtv_middleware/pkg/maxtv_company_campaigns"
	"maxtv_middleware/pkg/maxtv_company_orders"
	"maxtv_middleware/pkg/maxtv_company_payments"
	"maxtv_middleware/pkg/maxtv_themes"
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
		auth.GET("/maxtv-building-by-screen", maxtv_buildings.GetMaxTvBuildingByScreen)

		auth.GET("/maxtv-screens", maxtv_buildings.GetMaxTvScreens)

		auth.GET("/maxtv-themes", maxtv_themes.GetMaxTvThemes)

		auth.GET("/building-ratecard/:building_id", demografics.GetBuildingRatecard)

		auth.GET("/building-stats/:building_id", demografics.GetBuildingStat)

		auth.GET("/payments", maxtv_company_payments.GetPayments)

		auth.GET("/orders", maxtv_company_orders.GetOrders)
		auth.GET("/order/:order_id", maxtv_company_orders.GetOrder)

		auth.GET("/accounts", maxtv_companies.GetAccounts)
		auth.GET("/account/:company_id", maxtv_companies.GetAccount)
		auth.GET("/lead", maxtv_companies.GetLead)

		auth.GET("/campaigns", maxtv_company_campaigns.GetCampaign)
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
