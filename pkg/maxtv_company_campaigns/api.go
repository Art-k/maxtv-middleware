package maxtv_company_campaigns

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	var splitByDate *time.Time
	splitByDateStr := c.Query("split_by_date")
	if splitByDateStr != "" {
		sbd, err := time.Parse("2006-01-02", splitByDateStr)
		if err != nil {
			splitByDate = &sbd
		} else {
			splitByDate = nil
		}
	} else {
		splitByDate = nil
	}

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
	err := db.
		Preload(clause.Associations).
		Find(&response).Error

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if response.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	ProcessCampaignData(&response, splitByDate)
	c.JSON(http.StatusOK, response)
}

func GetCampaignBuildings(c *gin.Context) {

	campaignId := c.Param("campaign_id")

	var displays []MaxtvCompanyCampaignDisplay

	DB.
		//Preload(clause.Associations).
		Where("campaign_id = ?", campaignId).
		Find(&displays)

	c.JSON(http.StatusOK, displays)

}

func GetCampaignMedia(c *gin.Context) {

	campaignId := c.Param("campaign_id")

	var displays []MaxtvCompanyCampaignMedia

	DB.
		//Preload(clause.Associations).
		Where("campaign_id = ?", campaignId).
		Find(&displays)

	c.JSON(http.StatusOK, displays)

}

func GetCampaigns(c *gin.Context) {

	st := time.Now()

	status := strings.ToLower(c.Query("status"))
	orderIdStr := strings.ToLower(c.Query("order-id"))
	campaignType := strings.ToUpper(c.Query("campaign-type"))
	companyIdStr := strings.ToUpper(c.Query("company-id"))
	startDate := c.Query("start-date")
	endDate := c.Query("end-date")
	orderBy := c.Query("order-by")
	createdOn := c.Query("created-on")
	pageStr := c.Query("page")
	perPageStr := c.Query("per-page")

	var err error
	var page int
	page = 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}

	var perPage int
	perPage = 15
	if perPageStr != "" {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil {
			perPage = 15
		}
	}

	var splitByDate *time.Time
	splitByDateStr := c.Query("split_by_date")
	if splitByDateStr != "" {
		sbd, err := time.Parse("2006-01-02", splitByDateStr)
		if err == nil {
			splitByDate = &sbd
		} else {
			splitByDate = nil
		}
	} else {
		splitByDate = nil
	}

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

	DateQuery("start_date", startDate, db)
	DateQuery("end_date", endDate, db)
	DateQuery("created_on", createdOn, db)

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
	db.
		Preload(clause.Associations).
		Limit(perPage).
		Offset(perPage * (page - 1)).
		Find(&campaigns)

	st1 := time.Now()

	for ind, _ := range campaigns {
		ProcessCampaignData(&campaigns[ind], splitByDate)
	}

	response.Entities = campaigns
	response.ResponseTook = time.Now().Sub(st).Seconds()
	response.ProcessingTook = time.Now().Sub(st1).Seconds()
	c.JSON(http.StatusOK, response)

}

func ProcessCampaignData(camp *MaxtvCompanyCampaign, splitByDate *time.Time) {

	camp.LinkToOrder = "https://maxtvmedia.com/cms/?a=211&tab=orders&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(camp.CompanyId) +
		"&order_id=" + strconv.Itoa(camp.OrderId)
	camp.LinkToCompany = "https://maxtvmedia.com/cms/?a=211&tab=details&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(camp.CompanyId)
	camp.LinkToCampaign = "https://maxtvmedia.com/cms/?a=211&tab=campaigns&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(camp.CompanyId) +
		"&campaign_id=" + strconv.Itoa(camp.ID)

	camp.LinkToImpressionReport = "https://campaign-report.maxtvmedia.com/analytics/" + camp.ShortUrl

	camp.LinkToStatJson = "https://proposal-api.maxtvmedia.com/campaign/" + camp.ShortUrl + "/stat"

	var now time.Time
	if splitByDate == nil {
		now = time.Now()
	} else {
		now = *splitByDate
	}

	if now.After(camp.StartDate) && now.Before(camp.EndDate) {
		camp.PastDays = int(now.Sub(camp.StartDate).Hours()/24) + 1
		camp.RemainingDays = int(camp.EndDate.Sub(now).Hours() / 24)
	}

	camp.CampaignLength = int(camp.EndDate.Sub(camp.StartDate).Hours()/24) + 1

	if now.After(camp.StartDate) && now.After(camp.EndDate) {
		camp.PastDays = camp.CampaignLength
	}

	if now.Before(camp.StartDate) && now.Before(camp.EndDate) {
		camp.RemainingDays = camp.CampaignLength
	}

}
