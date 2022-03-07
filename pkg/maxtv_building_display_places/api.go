package maxtv_building_display_places

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetMaxTvDisplayPlace(c *gin.Context) {
	id := c.Param("id")
	var place db_interface.MaxtvBuildingDisplayPlace
	DB.Where("id = ?", id).Find(&place)
	if place.Id == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, place)
}

func GetMaxTvDisplayPlaces(c *gin.Context) {
	var place []db_interface.MaxtvBuildingDisplayPlace
	DB.Find(&place)
	c.JSON(http.StatusOK, place)
}
