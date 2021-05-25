package maxtv_companies

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
)

func GetAccount(c *gin.Context) {

	cIdStr := c.Param("company_id")
	var cId int
	var err error
	if cIdStr != "" {
		cId, err = strconv.Atoi(cIdStr)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var account MaxtvCompanie
	DB.Where("id = ?", cId).Find(&account)
	c.JSON(http.StatusOK, account)

}

func GetAccounts(c *gin.Context) {

	var accounts []MaxtvCompanie
	DB.Where("type = ?", "account").Order("created_on desc").Find(&accounts)
	c.JSON(http.StatusOK, accounts)

}

func GetLead(c *gin.Context) {

	var leads []MaxtvCompanie
	DB.Where("type = ?", "lead").Order("created_on desc").Find(&leads)
	c.JSON(http.StatusOK, leads)

}
