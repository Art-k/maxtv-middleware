package pythonReporter

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
	"time"
)

type responseType struct {
	MccId       string
	IsPaid      bool
	SaleDate    time.Time
	ActiveOrder TOrderInfo
	CompanyID   int
	OrderID     int
	Message     string
}

func responseFalse(mccId string, c *gin.Context, msg string) {
	response := responseType{
		MccId:   mccId,
		IsPaid:  false,
		Message: msg,
	}
	c.JSON(http.StatusOK, response)
}

func IsPaidMaxtvBuilding(c *gin.Context) {

	mccId := c.Query("mcc_id")
	var dbMaxTvBuilding db_interface.MaxtvBuilding
	err := DB.Where("mcc_id = ?", mccId).Find(&dbMaxTvBuilding).Error
	if err != nil {
		Log.Error(err)
		RespondWithError(c, http.StatusInternalServerError, err)
	}

	if dbMaxTvBuilding.Id == 0 {
		responseFalse(mccId, c, "There is no building linked to this MCC building")
		return
	}

	var dbMaxTvCompany db_interface.MaxtvCompanie
	err = DB.Where("building_id = ?", dbMaxTvBuilding.Id).Find(&dbMaxTvCompany).Error
	if err != nil {
		Log.Error(err)
		RespondWithError(c, http.StatusInternalServerError, err)
	}

	if dbMaxTvCompany.Id == 0 {
		responseFalse(mccId, c, "There is no Building Account connected to this MCC building")
		return
	}

	var dbMaxTvCompanyOrders []db_interface.MaxtvCompanyOrder
	err = DB.Where("company_id = ?", dbMaxTvCompany.Id).Find(&dbMaxTvCompanyOrders).Error
	if err != nil {
		Log.Error(err)
		RespondWithError(c, http.StatusInternalServerError, err)
	}

	if len(dbMaxTvCompanyOrders) == 0 {
		responseFalse(mccId, c, "There is no Orders")
		return
	}

	var response responseType
	response.MccId = mccId

	for _, order := range dbMaxTvCompanyOrders {

		if time.Now().Before(order.SaleDate.AddDate(5, 0, 0)) {
			response.ActiveOrder = String2Order(order.Payments)
		}

		if response.ActiveOrder.Amount > 100 {
			response.IsPaid = true
			response.SaleDate = order.SaleDate
			response.CompanyID = dbMaxTvCompany.Id
			response.OrderID = order.Id
		}

	}

	c.JSON(http.StatusOK, response)

}
