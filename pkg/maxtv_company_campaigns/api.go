package maxtv_company_campaigns

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type responseType struct {
	Total    int64                  `json:"total"`
	Entities []MaxtvCompanyCampaign `json:"entities"`
}

func GetCampaign(c *gin.Context) {

	status := strings.ToLower(c.Query("status"))
	orderIdStr := strings.ToLower(c.Query("order-id"))

	db := DB.Model(&MaxtvCompanyCampaign{})

	if status != "" {
		db = db.Where("status = ?", status)
		switch status {
		case "active":
			db = db.Where("end_date > ?", time.Now()).Where("start_date < ?", time.Now())
		}
	}

	if orderIdStr != "" {
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		db = db.Where("order_id = ?", orderId)
	}

	var response responseType
	db.Count(&response.Total)

	var campaigns []MaxtvCompanyCampaign
	db.Find(&campaigns)
	response.Entities = campaigns
	c.JSON(http.StatusOK, response)

}
