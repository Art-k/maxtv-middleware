package maxtv_company_orders

import (
	"fmt"
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
		processingOrder(&response.Entities[ind])
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
	processingOrder(&order)
	c.JSON(http.StatusOK, order)

}

func getValue(v []string) string {
	switch v[0] {
	case "s":
		return strings.Replace(v[2], "\"", "", -1)
	case "d":
		return strings.Replace(v[1], "\"", "", -1)
	case "i":
		return strings.Replace(v[1], "\"", "", -1)
	default:
		return "0"
	}
}

func processingOrder(order *MaxtvCompanyOrder) {

	order.LinkToOrder = "https://maxtvmedia.com/cms/?a=211&tab=orders&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(order.CompanyId) +
		"&order_id=" + strconv.Itoa(order.Id)
	order.LinkToCompany = "https://maxtvmedia.com/cms/?a=211&tab=details&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(order.CompanyId)

	tmp := TruncateString(order.Payments, strings.Index(order.Payments, "{"), strings.Index(order.Payments, "}"))
	fmt.Println("\n", tmp)

	list := strings.Split(tmp, ";")

	for i := 0; i < len(list)-2; i = i + 2 {
		key := strings.Split(list[i], ":")
		val := strings.Split(list[i+1], ":")

		sw := strings.Replace(key[2], "\"", "", -1)
		switch sw {
		case "order_id":
			order.Details.OrderId = getValue(val)
		case "payments":
			order.Details.Payments, _ = strconv.Atoi(getValue(val))
		case "first_last_payment":
			order.Details.FirstLastPayment, _ = strconv.Atoi(getValue(val))
		case "include_design_fee":
			order.Details.IncludeDesignFee, _ = strconv.Atoi(getValue(val))
		case "amounts":
			order.Details.Amount, _ = strconv.ParseFloat(getValue(val), 64)
		//case "payments_first":
		//	oi.PaymentFirst = strings.Replace(val[2], "\"", "", -1)
		case "payments_start":
			order.Details.PaymentStart, _ = time.Parse("02-01-2006", strings.Replace(val[2], "\"", "", -1))
		case "payments_inc":
			order.Details.PaymentIncrement, _ = strconv.Atoi(getValue(val))
		case "payments_inc_type":
			order.Details.PaymentIncrementType = strings.Replace(getValue(val), "\"", "", -1)
		case "method":
			order.Details.Method = strings.Replace(getValue(val), "\"", "", -1)
		case "design_fee":
			order.Details.DesignFee, _ = strconv.ParseFloat(getValue(val), 64)
		case "currency":
			order.Details.Currency = strings.Replace(getValue(val), "\"", "", -1)
		case "tax":
			order.Details.Tax, _ = strconv.ParseFloat(getValue(val), 64)
		case "copied":
			order.Details.Copied, _ = strconv.Atoi(getValue(val))
			//case "method_other":
			//	oi.OrderId = strings.Replace(val[2], "\"", "", -1)
		}

	}

}
