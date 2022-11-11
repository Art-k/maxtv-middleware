package reports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm/clause"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"maxtv_middleware/pkg/maxtv_company_campaigns"
	"strconv"
	"time"
)

type PercentageRow struct {
	SumBySales float64
	INdex      int
}

type WeeklyReportDetails struct {
	Campaign db_interface.MaxtvCompanyCampaign
	Company  db_interface.MaxtvCompanie
	Order    db_interface.MaxtvCompanyOrder
	User     db_interface.MaxtvUser

	PricePerCampaign  float64
	NumberOfCampaigns int
}

type SaleName string
type ReportMonth time.Time

type WeeklyReportData struct {
	Order     int
	Sales     db_interface.MaxtvUser
	ReportRow map[ReportMonth]float64
}

func PrepareWeeklySaleReport(c *gin.Context) {

	fn := PrepareWeeklySaleReportDo(c.Query("debug"))
	c.File(fn)

}

func PrepareWeeklySaleReportDo(debug string) (reportFile string) {

	now := time.Now()

	beginOfTheMonth := GetBeginOfTheMonth()

	var wr []WeeklyReportDetails

	var campaigns []db_interface.MaxtvCompanyCampaign
	common.DB.
		Preload(clause.Associations).
		Where("end_date >= ?", beginOfTheMonth).
		Where("status = ?", "active").
		Where("order_id <> ?", 0).
		Where("parent_id = ?", 0).
		Where("type = ?", "primary").
		Find(&campaigns)

	tmpCmpNumber := 0

	var users []db_interface.MaxtvUser

	for cind, campaign := range campaigns {

		now1 := time.Now()

		maxtv_company_campaigns.ProcessCampaignData(&campaign, &beginOfTheMonth)

		var company db_interface.MaxtvCompanie
		common.DB.
			Preload(clause.Associations).
			Where("id = ?", campaign.CompanyId).
			Find(&company)

		var order db_interface.MaxtvCompanyOrder
		common.DB.Where("id = ?", campaign.OrderId).Find(&order)

		order.ProcessingOrder()
		fmt.Println("\n--------------------------------------------")
		fmt.Println(cind+1, " of ", len(campaigns))
		fmt.Println("Cmp ID : ", campaign.ID)
		fmt.Println("Cmp Start Date : ", campaign.StartDate.Format("2006-01-02"))
		fmt.Println("Cmp End Date : ", campaign.EndDate.Format("2006-01-02"))
		fmt.Println("Cmp Status : ", campaign.Status)
		fmt.Println("Cmp Order Id : ", campaign.OrderId)
		fmt.Println("Cmp Parent Id : ", campaign.ParentId)
		fmt.Println("Cmp Order Amount : ", order.Details.Amount)
		fmt.Println("Cmp Length : ", campaign.CampaignLength)
		fmt.Println("Cmp Sales Id : ", order.SalePerson)

		var user db_interface.MaxtvUser
		common.DB.Where("id = ?", order.SalePerson).Find(&user)

		uf := false
		for _, u := range users {
			if u.Id == user.Id {
				uf = true
				break
			}
		}
		if !uf {
			users = append(users, user)
		}

		rec := WeeklyReportDetails{Campaign: campaign, Order: order, User: user, Company: company}
		distributePricePerCampaign(&rec, order.Details.Amount)

		fmt.Println("Cmp Amount : ", rec.PricePerCampaign)
		fmt.Println("Cmp Count : ", rec.NumberOfCampaigns)

		wr = append(wr, rec)

		fmt.Println("One campaign took :", time.Now().Sub(now1))

		if debug == "1" {
			tmpCmpNumber += 1
			if tmpCmpNumber > 20 {
				break
			}
		}
	}

	f := excelize.NewFile()
	SheetName := "Sheet1"
	cRowIndex := 1
	f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &[]interface{}{"UB Media Inc."})
	cRowIndex++
	f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &[]interface{}{"Sales by Class Summary"})
	cRowIndex++
	f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &[]interface{}{
		beginOfTheMonth.Format("January 2006") + "-" + beginOfTheMonth.AddDate(0, 12, 0).Format("January 2006"),
	})

	f.MergeCell(SheetName, "A1:O1", "")
	f.MergeCell(SheetName, "A2:O2", "")
	f.MergeCell(SheetName, "A3:O3", "")

	cRowIndex++
	cRowIndex++

	header := []interface{}{""}
	for i := 0; i < 12; i++ {
		m := beginOfTheMonth.AddDate(0, i, 0)
		header = append(header, m.Format("Jan. 2006"))
	}
	header = append(header, "Total")
	header = append(header, "Percentage of sales")
	f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &header)
	cRowIndex += 1
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 12},
	})
	f.SetRowStyle(SheetName, 2, 2, headerStyle)
	f.SetColWidth(SheetName, "A", "A", 30)
	f.SetColWidth(SheetName, "B", "0", 15)

	grandTotal := 0.0
	percentageRowIndexes := []PercentageRow{}
	globalMonthSums := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, user := range users {
		outRow := []interface{}{user.Firstname + " " + user.Lastname}
		salesCampaigns := []WeeklyReportDetails{}
		sumBySales := 0.0
		for i := 0; i < 12; i++ {
			m := beginOfTheMonth.AddDate(0, i, 0)
			sumByMonth := 0.0
			for _, row := range wr {
				if row.User.Id == user.Id {
					if row.Campaign.StartDate.Month() == time.Time(m).Month() && row.Campaign.StartDate.Year() == time.Time(m).Year() {
						t := Date(time.Time(m).Year(), int(time.Time(m).AddDate(0, 1, 0).Month()), 0)
						daysInMonth := t.Day()
						campaignDays := daysInMonth - row.Campaign.StartDate.Day() + 1
						sumByMonth += float64(row.PricePerCampaign) / float64(row.Campaign.CampaignLength) * float64(campaignDays)
						salesCampaigns = append(salesCampaigns, row)
						continue
					}

					if row.Campaign.EndDate.Month() == time.Time(m).Month() && row.Campaign.StartDate.Year() == time.Time(m).Year() {
						campaignDays := row.Campaign.EndDate.Day() + 1
						sumByMonth += float64(row.PricePerCampaign) / float64(row.Campaign.CampaignLength) * float64(campaignDays)
						salesCampaigns = append(salesCampaigns, row)
						continue
					}

					if row.Campaign.StartDate.Before(time.Time(m)) && row.Campaign.EndDate.After(time.Time(m)) {
						t := Date(time.Time(m).Year(), int(time.Time(m).AddDate(0, 1, 0).Month()), 0)
						daysInMonth := t.Day()
						sumByMonth += float64(row.PricePerCampaign) / float64(row.Campaign.CampaignLength) * float64(daysInMonth)
						salesCampaigns = append(salesCampaigns, row)
						continue
					}
				}
			}
			outRow = append(outRow, sumByMonth)
			globalMonthSums[i] += sumByMonth
			sumBySales += sumByMonth
		}
		outRow = append(outRow, sumBySales)
		grandTotal += sumBySales
		f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &outRow)
		percentageRowIndexes = append(percentageRowIndexes, PercentageRow{SumBySales: sumBySales, INdex: cRowIndex})
		cRowIndex += 1

		outRow = []interface{}{"Order Number", "Account Name", "Total Amount", "Sales", "Campaign Length", "Campaign Start Date", "Campaign End Date", "Link To Campaign", "Link To Order"}
		f.SetSheetRow(SheetName, fmt.Sprintf("B%d", cRowIndex), &outRow)
		f.SetRowOutlineLevel(SheetName, cRowIndex, 1)
		st, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Bold: true}})
		f.SetRowStyle(SheetName, cRowIndex, cRowIndex, st)
		cRowIndex += 1
		for _, row := range salesCampaigns {
			outRow = []interface{}{
				row.Order.OrderNumber,
				row.Company.Name,
				row.Order.Details.Amount,
				row.User.Firstname + " " + row.User.Lastname,
				row.Campaign.CampaignLength,
				row.Campaign.StartDate.Format("2006-01-02"),
				row.Campaign.EndDate.Format("2006-01-02"),
				row.Campaign.LinkToCampaign,
				row.Campaign.LinkToOrder,
			}
			f.SetSheetRow(SheetName, fmt.Sprintf("B%d", cRowIndex), &outRow)
			f.SetRowOutlineLevel(SheetName, cRowIndex, 1)
			cRowIndex += 1
		}
	}

	for _, ri := range percentageRowIndexes {
		percentage := ri.SumBySales / grandTotal * 100
		f.SetCellValue(SheetName, fmt.Sprintf("O%d", ri.INdex), percentage)
	}

	reportName := "./data/weekly_sales_reports/WSR_" + time.Now().Format("2006-01-02_15:04:05") + ".xlsx"
	err := f.SaveAs(reportName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The total number of campaigns :", len(campaigns))
	fmt.Println("It took :", time.Now().Sub(now))

	return reportName
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func distributePricePerCampaign(d *WeeklyReportDetails, amount float64) {
	var cmps []db_interface.MaxtvCompanyCampaign
	common.DB.
		Preload(clause.Associations).
		Where("order_id = ?", d.Campaign.OrderId).Find(&cmps)

	overalDisplayCount := 0
	for _, cmp := range cmps {
		overalDisplayCount += len(cmp.Displays)
	}

	d.NumberOfCampaigns = len(cmps)

	switch len(cmps) {
	case 1:
		d.PricePerCampaign = amount
	case 0:
		panic("No campaigns found for order_id = " + strconv.Itoa(d.Campaign.OrderId))
	default:
		d.PricePerCampaign = amount * float64(len(d.Campaign.Displays)) / float64(overalDisplayCount)
	}

}

func GetBeginOfTheMonth() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}