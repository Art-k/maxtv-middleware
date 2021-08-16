package maxtv_building_display_resident_traffic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
)

type ScreenImpression struct {
	db_interface.MaxtvBuildingDisplayResidentTraffic
	ImpressionPerPlay float64 `json:"impression_per_play"`
}

func GetMaxTvScreenTraffic(c *gin.Context) {

	screenIdStr := c.Param("id")
	screenId, err := strconv.Atoi(screenIdStr)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	var traffic db_interface.MaxtvBuildingDisplayResidentTraffic
	common.DB.Where("display_id = ?", screenId).Find(&traffic)
	if traffic.ID == 0 {
		err := fmt.Errorf("Screen ID %s is not found ", screenIdStr)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	fmt.Println(traffic.DwellTime * traffic.Traffic)
	fmt.Println(float64(traffic.DwellTime) * float64(traffic.Traffic) / 86400.00)

	response := ScreenImpression{
		MaxtvBuildingDisplayResidentTraffic: traffic,
		ImpressionPerPlay:                   float64(traffic.DwellTime) * float64(traffic.Traffic) / 86400.0,
	}

	c.JSON(http.StatusOK, response)

}
