package reports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"math"
	. "maxtv_middleware/pkg/common"
	. "maxtv_middleware/pkg/db_interface"
	"maxtv_middleware/pkg/maxtv_company_campaigns"
	"strconv"
	"strings"
	"time"
)

const (
	usdCadRate = 1.3

	//tax = 0.13
)

const dateTimeLayout = "2006-01-02"

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
	CampaignId       string
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
		CampaignId:       "Campaign ID",
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

	var ignoredCompany = [...]int{175608, 127690}

	for _, val := range ignoredCompany {
		if val == id {
			Log.Tracef("The id '%d' is in ignore list '%d'", id, val)
			return true
		}
	}
	Log.Tracef("The id '%d' is not in ignore list", id)
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

func PrepareA543A(c *gin.Context) ReportA543A {

	splitDateStr := c.Query("split_by")

	var report ReportA543A

	var campaigns []MaxtvCompanyCampaign
	DB.
		Where("end_date > ?", "2020-01-01").
		Preload(clause.Associations).
		Order("created_on desc").
		Find(&campaigns)

	initHeader(&report)

	for ind, _ := range campaigns {
		if splitDateStr == "" {
			maxtv_company_campaigns.ProcessCampaignData(&campaigns[ind], nil)
		} else {
			splitDate, err := time.Parse(dateTimeLayout, splitDateStr)
			if err != nil {
				return report
			}
			maxtv_company_campaigns.ProcessCampaignData(&campaigns[ind], &splitDate)
		}
	}

	for ind, campaign := range campaigns {

		var rec ReportRecA543A
		var tax = 0.13

		K := getKSplit(&campaign, &campaigns)

		fmt.Printf("campaign %d of %d", ind, len(campaigns))

		rec.CampaignId = strconv.Itoa(campaign.ID)

		if campaign.OrderId == 0 {
			//Log.Tracef("The order id is 0, skip this company")
			continue
		}

		Log.Tracef("The campaign id : '%d', check if it is in ignore list", campaign.ID)
		if isIgnoredCampaign(campaign.CompanyId) {
			Log.Tracef("Skipped")
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
		order.Details.DesignFee, _ = strconv.ParseFloat(payments["design_fee"], 64)
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

		rec.StartDate = campaign.StartDate.Format(dateTimeLayout)
		rec.EndDate = campaign.EndDate.Format(dateTimeLayout)

		rec.Length = strconv.Itoa(campaign.CampaignLength)
		rec.SaleDate = order.SaleDate.Format(dateTimeLayout)

		//K := splitByCampaigns(&campaign)
		rec.DesignFee = strconv.FormatFloat(order.Details.DesignFee*K, 'f', 2, 64)
		paymentAmount := getOrderTotalAmount(&order)

		paymentAmount = paymentAmount / (tax + 1)

		if paymentAmount < 0 {
			paymentAmount = 0
		}

		if rec.Currency == "USD" {
			paymentAmount = paymentAmount * usdCadRate
		}

		rec.Total = strconv.FormatFloat(paymentAmount*K, 'f', 2, 64)
		rec.OrderTotal = strconv.FormatFloat(paymentAmount*K+order.Details.DesignFee*K, 'f', 2, 64)
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

		aDeposited = aDeposited / (tax + 1)
		aFuture = aFuture / (tax + 1)

		rec.Deposited = strconv.FormatFloat(aDeposited, 'f', 2, 64)
		rec.Future = strconv.FormatFloat(aFuture, 'f', 2, 64)

		if campaign.CampaignLength != 0 {
			rec.Difference = strconv.FormatFloat(paymentRemaining-aFuture, 'f', 2, 64)
		} else {
			rec.Difference = "0"
		}

		rec.LinkToOrder = campaign.LinkToOrder
		rec.LinkToCampaign = campaign.LinkToCampaign

		if campaign.EndDate.Before(time.Now()) {
			continue
		}

		report.Data = append(report.Data, rec)

		fmt.Println(tax)

	}

	return report
}

func getKSplit(camp *MaxtvCompanyCampaign, camps *[]MaxtvCompanyCampaign) float64 {

	if camp.OrderId == 0 {
		return 1
	}

	var campaignsLength int

	for _, rec := range *camps {

		if rec.OrderId == camp.OrderId {
			campaignsLength += rec.CampaignLength
		}

	}

	if camp.CampaignLength == campaignsLength {
		return 1
	}

	return float64(camp.CampaignLength) / float64(campaignsLength)

}
