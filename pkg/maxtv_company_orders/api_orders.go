package maxtv_company_orders

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type responseType struct {
	ResponseHeader
	Entities []MaxtvCompanyOrder `json:"entities"`
}

func GetOrders(c *gin.Context) {

	st := time.Now()

	companyIdStr := c.Query("company_id")
	orderBy := c.Query("order-by")

	db := DB.Model(&MaxtvCompanyOrder{})

	if companyIdStr != "" {
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		db = db.Where("company_id = ?", companyId)
	}

	if orderBy != "" {
		db.Order(strings.ReplaceAll(orderBy, "|", " "))
	} else {
		db.Order("sale_date desc")
	}

	var response responseType
	db.Count(&response.Total)
	db.Find(&response.Entities)

	st1 := time.Now()
	for ind, _ := range response.Entities {
		response.Entities[ind].ProcessingOrder()
	}
	response.ResponseTook = time.Now().Sub(st).Seconds()
	response.ProcessingTook = time.Now().Sub(st1).Seconds()

	c.JSON(http.StatusOK, response)

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
	DB.Where("id = ?", orderId).Find(&order)
	order.ProcessingOrder()
	c.JSON(http.StatusOK, order)

}
