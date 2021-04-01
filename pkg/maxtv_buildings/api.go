package maxtv_buildings

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetMaxTvBuildings(c *gin.Context) {

	var buildings []MaxtvBuilding

	network := c.Query("Network")

	db := DB.Model(&MaxtvBuilding{})

	if network != "" {
		db.Where("network = ?", network)
	}

	db.Find(&buildings)

	c.JSON(http.StatusOK, buildings)
}
