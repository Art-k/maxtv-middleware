package maxtv_companies

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetAccounts(c *gin.Context) {

	var accounts []MaxtvCompanie
	DB.Where("type = ?", "account").Find(&accounts)
	c.JSON(http.StatusOK, accounts)

}
