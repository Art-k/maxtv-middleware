package reports

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportList(c *gin.Context) {

	type Reports struct {
		Name        string
		Description string
		EndPoint    string
	}

	var reports []Reports

	reports = append(reports, Reports{
		Name:        "Active Campaigns",
		Description: "Moshes report, placed in google sheets",
		EndPoint:    "/report/a543_a",
	})

	c.JSON(http.StatusOK, reports)

}

func ReportGet(c *gin.Context) {

	report := c.Param("report_name")

	switch report {
	case "a543_a":
		resp := PrepareA543A(c)
		c.JSON(http.StatusOK, resp)
	case "a543_b":
		resp := PrepareA543B(c)
		c.JSON(http.StatusOK, resp)
	case "weekly_sale_report":
		PrepareWeeklySaleReport(c)
		c.AbortWithStatus(http.StatusOK)
	default:

	}

}
