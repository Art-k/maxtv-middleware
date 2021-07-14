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

func GetMaxTvBuildingByScreen(c *gin.Context) {

	screen := c.Query("display_sysid")

	var dbScreen MaxtvBuildingDisplay
	DB.Where("sysid = ?", screen).Find(&dbScreen)
	if dbScreen.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var dbBuilding MaxtvBuilding
	DB.Where("id = ?", dbScreen.BuildingId).Find(&dbBuilding)
	if dbBuilding.Id == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, dbBuilding)

}

func GetMaxTvBuilding(c *gin.Context) {

	Id := c.Param("id")
	var building MaxtvBuilding
	DB.
		Where("id = ?", Id).
		Find(&building)

	if building.Id == 0 {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, building)

}

func GetMaxTvBuildings(c *gin.Context) {
	// TODO we need to add ability to download csv so we have some issue here
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
