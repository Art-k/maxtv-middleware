package reports

import (
	"fmt"
	"gorm.io/gorm/clause"
	"math"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"maxtv_middleware/pkg/maxtv_company_campaigns"
	"strconv"
	"strings"
)

const (
	usdCadRate = 1.3

//tax = 0.13
)

var ignoredCompany = [...]int{175608}

type ReportRecA543A struct {
	Currency         string
	Company          string
	Campaign         string
	Status           string
	Order            string
	Invoice          string
	SaleDate         string
	StartDate        string
	EndDate          string
	Length           string
	OrderTotal       string
	DesignFee        string
	Total            string
	PastDays         string
	RemainingDays    string
	PaymentDue       string
	RemainingPayment string
	Deposited        string
	Future           string
	Difference       string
	LinkToOrder      string
	LinkToCampaign   string
}

type ReportA543A struct {
	Header ReportRecA543A   `json:"header"`
	Data   []ReportRecA543A `json:"data"`
}

func initHeader(rep *ReportA543A) {
	rep.Header = ReportRecA543A{
		Currency:         "Currency",
		Company:          "Company",
		Campaign:         "Campaign",
		Status:           "Status",
		Order:            "Order",
		Invoice:          "Invoice",
		SaleDate:         "Sale Date",
		StartDate:        "Start Date",
		EndDate:          "End Date",
		Length:           "Length",
		OrderTotal:       "Order Total",
		DesignFee:        "Design Fee",
		Total:            "Total",
		PastDays:         "Past Days",
		RemainingDays:    "Remaining Days",
		PaymentDue:       "Payment Due",
		RemainingPayment: "Remaining Payment",
		Deposited:        "Deposited",
		Future:           "Future",
		Difference:       "Difference",
		LinkToOrder:      "Link To Order",
		LinkToCampaign:   "Link To Campaign",
	}
}

func splitByCampaigns(ac *MaxtvCompanyCampaign) float64 {

	var campaigns []MaxtvCompanyCampaign
	var sumLength int
	var cLength int

	DB.Where("order_id = ?", ac.OrderId).Find(&campaigns)

	if len(campaigns) == 1 {
		return 1
	} else {
		for _, cmp := range campaigns {
			cmpLength := int(math.Round(cmp.EndDate.Sub(cmp.StartDate).Hours() / 24))
			sumLength += cmpLength
			if cmp.ID == ac.ID {
				cLength = cmpLength
			}
		}
		if sumLength == 0 {
			return float64(1) / float64(len(campaigns))
		} else {
			return float64(cLength) / float64(sumLength)
		}
	}

}

func isIgnoredCampaign(id int) bool {
	for _, val := range ignoredCompany {
		if val == id {
			return true
		}
	}
	return false
}

func dPayments(pmts string) map[string]string {

	//Log.Tracef("the payment string is '%s'", pmts)

	lp := make(map[string]string)
	sp := strings.Index(pmts, "{")
	ep := strings.Index(pmts, "}")
	tmp := pmts[sp+1 : ep]
	tmpList := strings.Split(tmp, ";")

	for ind, val := range tmpList {

		if ind%2 != 0 {
			continue
		}

		if len(tmpList)-3 == ind {
			break
		}

		//fmt.Println("'",val,"'",tmpList[ind+1],"'")

		if len(val) < 2 || len(tmpList[ind+1]) < 2 {
			continue
		}

		var item string
		switch val[:2] {
		case "s:":
			item = strings.ReplaceAll(strings.Split(val, ":")[2], "\"", "")
		case "i:":
			item = strings.ReplaceAll(strings.Split(val, ":")[1], "\"", "")
		}

		var item1 string
		switch tmpList[ind+1][:2] {
		case "s:":
			item1 = strings.ReplaceAll(strings.Split(tmpList[ind+1], ":")[2], "\"", "")
		case "i:":
			item1 = strings.ReplaceAll(strings.Split(tmpList[ind+1], ":")[1], "\"", "")
		}

		lp[item] = item1

	}
	//fmt.Println(lp)
	return lp
}

func getOrderTotalAmount(order *MaxtvCompanyOrder) float64 {

	var total float64
	var payments []MaxtvCompanyPayment
	DB.Where("order_id = ?", order.Id).Find(&payments)

	for _, payment := range payments {
		switch payment.Status {
		case "Deposited":
			total += payment.Amount
		case "Future Payment":
			total += payment.Amount
		case "Declined":
			total += payment.Amount
		case "Refund":
			total += payment.Amount
		case "Cancellations":

		case "Bad Debt":

		case "Collection":

		}
	}

	return total
}

func getActualPayments(orderId int) (float64, float64, float64) {
	var payments []MaxtvCompanyPayment
	DB.Where("order_id = ?", orderId).Find(&payments)

	var amount float64
	var sDeposited float64
	var sFuture float64

	for _, val := range payments {
		amount += val.Amount
		switch val.Status {
		case "Deposited":
			sDeposited += val.Amount
		case "Future Payment":
			sFuture += val.Amount
		}
	}

	return amount, sDeposited, sFuture
}

func PrepareA543A() ReportA543A {

	var report ReportA543A

	var campaigns []MaxtvCompanyCampaign
	DB.
		Where("end_date > ?", "2020-01-01").
		Preload(clause.Associations).
		Find(&campaigns)

	initHeader(&report)

	for ind, _ := range campaigns {
		maxtv_company_campaigns.ProcessCampaignData(&campaigns[ind], nil)
	}

	for ind, campaign := range campaigns {

		var rec ReportRecA543A
		var tax = 0.13

		fmt.Printf("campaign %d of %d", ind, len(campaigns))

		if campaign.OrderId == 0 {
			continue
		}

		if isIgnoredCampaign(campaign.ID) {
			continue
		}

		var company MaxtvCompanie
		DB.Where("id = ?", campaign.CompanyId).Find(&company)

		var order MaxtvCompanyOrder
		DB.Where("id = ?", campaign.OrderId).Find(&order)

		if order.Payments == "" {
			continue
		}
		payments := dPayments(order.Payments)
		_, aDeposited, aFuture := getActualPayments(campaign.OrderId)

		if val, ok := payments["currency"]; ok {
			rec.Currency = val
			if val == "USD" {
				tax = 0
			}
		} else {
			rec.Currency = "CAD*"
		}

		rec.Company = company.Name
		if rec.Company == "Mahnaz Mahourvand" {
			continue
		}

		rec.Campaign = campaign.Title
		rec.Status = campaign.Status
		rec.Order = order.OrderNumber
		rec.Invoice = order.Invoice

		rec.StartDate = campaign.StartDate.Format("2006-01-02")
		rec.EndDate = campaign.EndDate.Format("2006-01-02")

		rec.Length = strconv.Itoa(campaign.CampaignLength)
		rec.SaleDate = order.SaleDate.Format("2006-01-02")

		K := splitByCampaigns(&campaign)
		rec.DesignFee = strconv.FormatFloat(order.Details.DesignFee, 'f', 2, 64)
		paymentAmount := getOrderTotalAmount(&order)
		if paymentAmount < 0 {
			paymentAmount = 0
		}

		if rec.Currency == "USD" {
			paymentAmount = paymentAmount * usdCadRate
		}

		rec.Total = strconv.FormatFloat(paymentAmount, 'f', 2, 64)
		rec.OrderTotal = strconv.FormatFloat(paymentAmount+order.Details.DesignFee, 'f', 2, 64)
		rec.RemainingDays = strconv.Itoa(campaign.RemainingDays)
		rec.PastDays = strconv.Itoa(campaign.PastDays)

		var paymentDue float64
		var paymentRemaining float64

		if campaign.CampaignLength > 0 {
			paymentDue = (paymentAmount / float64(campaign.CampaignLength)) * float64(campaign.PastDays) * K
		}
		if paymentDue < 0 {
			paymentDue = 0
		}
		rec.PaymentDue = strconv.FormatFloat(paymentDue, 'f', 2, 64)

		if campaign.CampaignLength > 0 {
			paymentRemaining = (paymentAmount / float64(campaign.CampaignLength)) * float64(campaign.RemainingDays) * K
		}
		rec.RemainingPayment = strconv.FormatFloat(paymentRemaining, 'f', 2, 64)

		rec.Deposited = strconv.FormatFloat(aDeposited, 'f', 2, 64)
		rec.Future = strconv.FormatFloat(aFuture, 'f', 2, 64)

		if campaign.CampaignLength != 0 {
			rec.Difference = strconv.FormatFloat(paymentRemaining-aFuture, 'f', 2, 64)
		} else {
			rec.Difference = "0"
		}

		rec.LinkToOrder = order.LinkToOrder
		rec.LinkToCampaign = campaign.LinkToCampaign

		report.Data = append(report.Data, rec)

		fmt.Println(tax)

	}

	return report
}
