package maxtv_company_campaigns

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetCampaign(c *gin.Context) {

	var campaigns []MaxtvCompanyCampaign
	DB.Find(&campaigns)
	c.JSON(http.StatusOK, campaigns)

}
