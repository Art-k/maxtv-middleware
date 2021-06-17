package maxtv_db

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	"net/http"
	//. "maxtv_middleware/pkg/db_interface"
)

type columns struct {
	Field string
	Type  string
}

func Describe(c *gin.Context) {

	tableName := c.Param("table_name")
	var columns []columns
	DB.Raw("DESCRIBE " + tableName).Scan(&columns)

	c.JSON(http.StatusOK, columns)

}
