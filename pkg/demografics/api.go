package demografics

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetBuildingRatecard(c *gin.Context) {

	buildingId, err := common.GetIntParameter(c, "building_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var dbRateCard []db_interface.BuildingRatecard
	common.DB.Where("building_id = ?", buildingId).Find(&dbRateCard)

	c.JSON(http.StatusOK, dbRateCard)

}

func GetBuildingStat(c *gin.Context) {

	buildingId, err := common.GetIntParameter(c, "building_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var dbBuildingStat db_interface.BuildingStat
	common.DB.Where("building_id = ?", buildingId).Find(&dbBuildingStat)

	c.JSON(http.StatusOK, dbBuildingStat)

}
