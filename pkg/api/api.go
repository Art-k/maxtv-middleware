package api

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/demografics"
	"maxtv_middleware/pkg/maxtv_building_display_places"
	"maxtv_middleware/pkg/maxtv_building_display_resident_traffic"
	"maxtv_middleware/pkg/maxtv_building_display_sizes"
	"maxtv_middleware/pkg/maxtv_buildings"
	"maxtv_middleware/pkg/maxtv_companies"
	"maxtv_middleware/pkg/maxtv_company_campaigns"
	"maxtv_middleware/pkg/maxtv_company_orders"
	"maxtv_middleware/pkg/maxtv_company_payments"
	"maxtv_middleware/pkg/maxtv_db"
	"maxtv_middleware/pkg/maxtv_themes"
	"maxtv_middleware/pkg/maxtv_users"
	"maxtv_middleware/pkg/pythonReporter"
	"maxtv_middleware/pkg/reports"
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
		auth.GET("/maxtv-building/:id", maxtv_buildings.GetMaxTvBuilding)
		auth.GET("/maxtv-building-by-screen", maxtv_buildings.GetMaxTvBuildingByScreen)

		auth.GET("/maxtv-screen/:id", maxtv_buildings.GetMaxTvScreen)
		auth.GET("/maxtv-screen-traffic/:id", maxtv_building_display_resident_traffic.GetMaxTvScreenTraffic)
		auth.GET("/maxtv-screens", maxtv_buildings.GetMaxTvScreens)

		auth.GET("/maxtv-screen-place/:id", maxtv_building_display_places.GetMaxTvDisplayPlace)
		auth.GET("/maxtv-screen-places", maxtv_building_display_places.GetMaxTvDisplayPlaces)

		auth.GET("/maxtv-screen-size/:id", maxtv_building_display_sizes.GetMaxTvScreenSize)
		auth.GET("/maxtv-screen-sizes", maxtv_building_display_sizes.GetMaxTvScreenSizes)

		auth.GET("/maxtv-themes", maxtv_themes.GetMaxTvThemes)
		auth.GET("/maxtv-theme/:id", maxtv_themes.GetMaxTvTheme)

		auth.GET("/building-ratecard/:building_id", demografics.GetBuildingRatecard)
		//auth.GET("/building-ratecard/:building_id/household-income", demografics.GetBuildingRatecardHI)

		auth.GET("/building-stats/:building_id", demografics.GetBuildingStat)

		auth.GET("/payments", maxtv_company_payments.GetPayments)

		auth.GET("/orders", maxtv_company_orders.GetOrders)
		auth.GET("/order/:order_id", maxtv_company_orders.GetOrder)

		auth.GET("/accounts", maxtv_companies.GetAccounts)
		auth.GET("/account/:company_id", maxtv_companies.GetAccount)

		auth.GET("/lead", maxtv_companies.GetLead)

		auth.GET("/campaigns", maxtv_company_campaigns.GetCampaigns)
		auth.GET("/campaign/:campaign_id", maxtv_company_campaigns.GetCampaign)
		auth.GET("/campaign/:campaign_id/buildings", maxtv_company_campaigns.GetCampaignBuildings)
		auth.GET("/campaign/:campaign_id/media", maxtv_company_campaigns.GetCampaignMedia)

		auth.GET("/users", maxtv_users.GetUsers)
		auth.GET("/user/:user_id", maxtv_users.GetUser)

		auth.GET("/reports", reports.ReportList)
		auth.GET("/report/:report_name", reports.ReportGet)

		auth.GET("/db/describe/:table_name", maxtv_db.Describe)
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
