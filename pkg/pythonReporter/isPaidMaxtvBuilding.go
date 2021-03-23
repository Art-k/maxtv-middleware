package pythonReporter

import (
	"github.com/gin-gonic/gin"
	. "maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"net/http"
)

func IsPaidMaxtvBuilding(c *gin.Context) {

	mccId := c.Query("mcc_id")
	var dbMaxTvBuilding db_interface.MaxtvBuilding
	err := DB.Where("mcc_id = ?", mccId).Find(&dbMaxTvBuilding).Error
	if err != nil {
		Log.Error(err)
		RespondWithError(c, http.StatusInternalServerError, err)
	}

	var dbMaxTvCompany db_interface.MaxtvCompanie
	err = DB.Where("building_id = ?", dbMaxTvBuilding.Id).Find(&dbMaxTvCompany).Error
	if err != nil {
		Log.Error(err)
		RespondWithError(c, http.StatusInternalServerError, err)
	}

	var dbMaxTvCompanyOrders []db_interface.MaxtvCompanyOrder
	err = DB.Where("company_id = ?", dbMaxTvCompany.Id).Find(&dbMaxTvCompanyOrders).Error
	if err != nil {
		Log.Error(err)
		RespondWithError(c, http.StatusInternalServerError, err)
	}

	for _, order := range dbMaxTvCompanyOrders {
		var dbPayments []db_interface.MaxtvCompanyPayment
		err = DB.
			Where("company_id = ?", dbMaxTvCompany.Id).
			Where("order_id = ?", order.Id).
			Find(&dbPayments).
			Error
		if err != nil {
			Log.Error(err)
			RespondWithError(c, http.StatusInternalServerError, err)
		}
	}

}
