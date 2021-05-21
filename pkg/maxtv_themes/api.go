package maxtv_themes

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
)

func GetMaxTvThemes(c *gin.Context) {

	var dbThemes []db_interface.MaxtvTheme
	common.DB.Find(&dbThemes)
	c.JSON(http.StatusOK, dbThemes)

}
