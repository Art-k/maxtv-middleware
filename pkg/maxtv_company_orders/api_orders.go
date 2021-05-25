package maxtv_company_orders

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
)

func GetOrders(c *gin.Context) {

}

func GetOrder(c *gin.Context) {

	orderIdStr := c.Param("order_id")
	var orderId int
	var err error
	if orderIdStr != "" {
		orderId, err = strconv.Atoi(orderIdStr)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var order MaxtvCompanyOrder
	common.DB.Where("id = ?", orderId).Find(&order)
	c.JSON(http.StatusOK, order)

}
