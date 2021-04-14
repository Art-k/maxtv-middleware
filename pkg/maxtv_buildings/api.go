package maxtv_buildings

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
)

func GetMaxTvScreens(c *gin.Context) {

	var displays []MaxtvBuildingDisplay

	buildingIdStr := c.Query("BuildingId")

	db := DB.Model(&MaxtvBuildingDisplay{})
	if buildingIdStr != "" {
		bldId, err := strconv.Atoi(buildingIdStr)
		if err == nil {
			db = db.Where("building_id = ?", bldId)
		}
	}

	db.Find(&displays)

	c.JSON(http.StatusOK, displays)

}

func GetMaxTvBuildings(c *gin.Context) {
	// TODO we need to add ability to download csv
	var buildings []MaxtvBuilding

	network := c.Query("Network")
	showOnMap := c.Query("ShowOnMap")

	db := DB.Model(&MaxtvBuilding{})

	if network != "" {
		db.Where("network = ?", network)
	}

	if showOnMap != "" {
		som := true
		if showOnMap != "1" {
			som = false
		}
		db.Where("show_on_map = ?", som)
	}

	db.Find(&buildings)

	c.JSON(http.StatusOK, buildings)
}
