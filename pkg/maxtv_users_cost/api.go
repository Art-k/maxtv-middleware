package maxtv_users_cost

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
)

type responseType struct {
	ResponseHeader
	Entities []UserCosts `json:"entities"`
}

func GetUsersCost(c *gin.Context) {

	var response responseType
	DB.Find(&response.Entities)
	c.JSON(http.StatusOK, response)

}

func GetUserCost(c *gin.Context) {

	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	var response UserCosts
	DB.Where("user_id = ?", userId).Where("deleted = ?", 0).Last(&response)

	c.JSON(http.StatusOK, response)

}

func GetUserCostHistory(c *gin.Context) {

	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var response responseType

	DB.Where("user_id = ?", userId).Find(&response.Entities)
	DB.Where("user_id = ?", userId).Count(&response.Total)

	c.JSON(http.StatusOK, response)

}
