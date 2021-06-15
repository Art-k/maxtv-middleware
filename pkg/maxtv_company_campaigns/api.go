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
	ResponseHeader
	Entities []MaxtvCompanyCampaign `json:"entities"`
}

func GetCampaign(c *gin.Context) {

	db := DB.Model(&MaxtvCompanyCampaign{})

	campaignIdStr := c.Param("campaign_id")
	if campaignIdStr != "" {
		campaignId, err := strconv.Atoi(campaignIdStr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		db.Where("id = ?", campaignId)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var response MaxtvCompanyCampaign
	err := db.Find(&response).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if response.Id == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	processCampaignData(&response)
	c.JSON(http.StatusOK, response)
}

func GetCampaigns(c *gin.Context) {

	st := time.Now()

	status := strings.ToLower(c.Query("status"))
	orderIdStr := strings.ToLower(c.Query("order-id"))
	campaignType := strings.ToUpper(c.Query("campaign-type"))
	companyIdStr := strings.ToUpper(c.Query("company-id"))
	orderBy := c.Query("order-by")

	db := DB.Model(&MaxtvCompanyCampaign{})

	if campaignType != "" {
		db = db.Where("ad_type = ?", campaignType)
	}

	if orderBy != "" {
		db = db.Order(strings.ReplaceAll(orderBy, "|", " "))
	} else {
		db = db.Order("created_on desc")
	}

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

	if companyIdStr != "" {
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		db = db.Where("company_id = ?", companyId)
	}

	var response responseType
	db.Count(&response.Total)

	var campaigns []MaxtvCompanyCampaign
	db.Find(&campaigns)

	st1 := time.Now()

	for ind, _ := range campaigns {
		processCampaignData(&campaigns[ind])
	}

	response.Entities = campaigns
	response.ResponseTook = time.Now().Sub(st).Seconds()
	response.ProcessingTook = time.Now().Sub(st1).Seconds()
	c.JSON(http.StatusOK, response)

}

func processCampaignData(camp *MaxtvCompanyCampaign) {

	camp.LinkToOrder = "https://maxtvmedia.com/cms/?a=211&tab=orders&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(camp.CompanyId) +
		"&order_id=" + strconv.Itoa(camp.OrderId)
	camp.LinkToCompany = "https://maxtvmedia.com/cms/?a=211&tab=details&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(camp.CompanyId)
	camp.LinkToCampaign = "https://maxtvmedia.com/cms/?a=211&tab=campaigns&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(camp.CompanyId) +
		"&campaign_id=" + strconv.Itoa(camp.Id)

	now := time.Now()
	if now.After(camp.StartDate) && now.Before(camp.EndDate) {
		camp.PastDays = int(now.Sub(camp.StartDate).Hours() / 24)
		camp.RemainingDays = int(camp.EndDate.Sub(now).Hours()/24) + 1
	}

	camp.CampaignLength = int(camp.EndDate.Sub(camp.StartDate).Hours() / 24)

	if now.After(camp.StartDate) && now.After(camp.EndDate) {
		camp.PastDays = camp.CampaignLength
	}

	if now.Before(camp.StartDate) && now.Before(camp.EndDate) {
		camp.RemainingDays = camp.CampaignLength
	}

}
