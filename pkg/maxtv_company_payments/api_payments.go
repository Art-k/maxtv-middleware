package maxtv_company_payments

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"net/http"
	"strconv"
)

func GetPayments(c *gin.Context) {

	orderIdStr := c.Query("order_id")
	if orderIdStr == "" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var payments []MaxtvCompanyPayment
	DB.Where("order_id = ?", orderId).Find(&payments)
	c.JSON(http.StatusOK, payments)

}
