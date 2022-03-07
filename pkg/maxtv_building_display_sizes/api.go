package maxtv_building_display_sizes

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"

	"maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetMaxTvScreenSize(c *gin.Context) {

	id := c.Param("id")
	size := db_interface.MaxtvBuildingDisplaySize{}
	DB.Where("id = ?", id).Find(&size)
	if size.Id == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, size)

}

func GetMaxTvScreenSizes(c *gin.Context) {

	var sizes []db_interface.MaxtvBuildingDisplaySize
	DB.Find(&sizes)
	c.JSON(http.StatusOK, sizes)

}
