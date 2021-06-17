package maxtv_users

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
	"time"
)

type responseType struct {
	common.ResponseHeader
	Entities []db_interface.MaxtvUser `json:"entities"`
}

func GetUser(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func GetUsers(c *gin.Context) {

	st := time.Now()

	var users []db_interface.MaxtvUser

	db := common.DB.Model(&db_interface.MaxtvUser{})

	db.Find(&users)

	var response responseType
	db.Count(&response.Total)

	st1 := time.Now()
	response.Entities = users
	response.ResponseTook = time.Now().Sub(st).Seconds()
	response.ProcessingTook = time.Now().Sub(st1).Seconds()
	c.JSON(http.StatusOK, response)

}
