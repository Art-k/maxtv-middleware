package maxtv_users

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
	"time"
)

type responseType struct {
	ResponseHeader
	Entities []MaxtvUser `json:"entities"`
}

func GetUser(c *gin.Context) {

	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user MaxtvUser
	DB.Where("id = ?", userId).Find(&user)

	if user.Id == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetUsers(c *gin.Context) {

	st := time.Now()

	var users []MaxtvUser

	db := DB.Model(&MaxtvUser{})

	db = db.Where("access_level in ?", []int{1, 200, 300, 500, 400, 600})
	db = db.Where("active = ?", true)

	db.Find(&users)

	var response responseType
	db.Count(&response.Total)

	//st1 := time.Now()

	response.Entities = users
	response.ResponseTook = time.Now().Sub(st).Seconds()
	response.ProcessingTook = -1

	c.JSON(http.StatusOK, response)

}
