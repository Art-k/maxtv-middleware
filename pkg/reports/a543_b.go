package reports

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/common"
)

type ReportRecA543B struct {
	Currency                 string
	Company                  string
	Address                  string
	SalesPerson              string
	CommissionRate           string
	Order                    string
	OrdersCount              string
	SaleDate                 string
	StartDate                string
	EndDate                  string
	Length                   string
	Screens                  string
	OrderTotal               string
	EquipmentAndInstallation string
	Maintenance              string
	TotalCost                string
	MarkUp                   string
	ScreenAllocation         string
	ServiceAllocation        string
	PastDays                 string
	RemainingDays            string
	PaymentDue               string
	RemainingPayment         string
	Deposited                string
	Future                   string
	Difference               string
	CompanyLink              string
	OrderLink                string
}

type ReportA543B struct {
	Header ReportRecA543B   `json:"header"`
	Data   []ReportRecA543B `json:"data"`
}

func initHeaderA543B(rep *ReportA543B) {
	rep.Header = ReportRecA543B{
		Currency:                 "Currency",
		Company:                  "Company",
		Address:                  "Address",
		SalesPerson:              "Sales Person",
		CommissionRate:           "Commission Rate",
		Order:                    "Order",
		OrdersCount:              "Orders Count",
		SaleDate:                 "Sale Date",
		StartDate:                "Start Date",
		EndDate:                  "End Date",
		Length:                   "Length",
		Screens:                  "Screens",
		OrderTotal:               "Order Total",
		EquipmentAndInstallation: "Equipment & Installation",
		Maintenance:              "Maintenance",
		TotalCost:                "Total Cost",
		MarkUp:                   "Mark Up",
		ScreenAllocation:         "Screen Allocation",
		ServiceAllocation:        "Service Allocation",
		PastDays:                 "Past Days",
		RemainingDays:            "Remaining Days",
		PaymentDue:               "Payment Due",
		RemainingPayment:         "Remaining Payment",
		Deposited:                "Deposited",
		Future:                   "Future",
		Difference:               "Difference",
		CompanyLink:              "Company Link",
		OrderLink:                "Order Link",
	}
}

func PrepareA543B(c *gin.Context) ReportA543B {

	splitDateStr := c.Query("split_by")
	common.Log.Infof("Split by date : '%s'", splitDateStr)

	var report ReportA543B

	initHeaderA543B(&report)

	return report
}
