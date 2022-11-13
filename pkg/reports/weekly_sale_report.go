package reports

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"maxtv_middleware/pkg/common"
	"maxtv_middleware/pkg/db_interface"
	"maxtv_middleware/pkg/maxtv_company_campaigns"
	"strconv"
	"strings"
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

	NumberOfCampaigns int
	PricePerCampaign  float64
	PricePerDay       float64
	NumberOfDays      int
}

type SaleName string
type ReportMonth time.Time

type WeeklyReportData struct {
	Order     int
	Sales     db_interface.MaxtvUser
	ReportRow map[ReportMonth]float64
}

func PrepareWeeklySaleReport(c *gin.Context) {

	fn := PrepareWeeklySaleReportDo(c.Query("debug"), c.Query("year_mode"), c.Query("cache"), c.Query("year"))

	tmp := strings.Split(fn, "/")
	file := tmp[len(tmp)-1]

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+file)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	c.File(fn)

}

func PrepareWeeklySaleReportDo(debug, yearMode, cacheMode, year string) (reportFile string) {

	now := time.Now()

	var beginOfTheMonth time.Time
	if yearMode == "1" {
		beginOfTheMonth = GetBeginOfTheYear(year)
	} else {
		beginOfTheMonth = GetBeginOfTheMonth()
	}

	var wr []WeeklyReportDetails

	var campaigns []db_interface.MaxtvCompanyCampaign
	var orders []db_interface.MaxtvCompanyOrder
	var companies []db_interface.MaxtvCompanie
	var users []db_interface.MaxtvUser

	var ids []int
	var id_orders []int
	var id_users []int
	if cacheMode != "1" {
		common.DB.
			Model(&db_interface.MaxtvCompanyCampaign{}).
			Where("end_date >= ?", beginOfTheMonth).
			Where("status = ?", "active").
			Where("order_id <> ?", 0).
			Where("parent_id = ?", 0).
			Where("type = ?", "primary").
			Pluck("order_id", &ids)
		common.DB.
			Preload(clause.Associations).
			Where("id IN ?", ids).
			Find(&orders)
		common.DB.
			Model(&db_interface.MaxtvCompanyOrder{}).
			Where("id IN ?", ids).
			Pluck("id", &id_orders)
		common.DB.
			Model(&db_interface.MaxtvCompanyOrder{}).
			Where("id IN ?", id_orders).
			Pluck("sale_person", &id_users)
		file, _ := json.MarshalIndent(orders, "", " ")
		_ = ioutil.WriteFile("orders.json", file, 0644)
	} else {
		file, _ := ioutil.ReadFile("orders.json")
		_ = json.Unmarshal([]byte(file), &orders)
	}

	if cacheMode != "1" {
		common.DB.
			Preload(clause.Associations).
			Where("order_id IN ?", id_orders).
			Find(&campaigns)

		common.DB.
			Model(&db_interface.MaxtvCompanyCampaign{}).
			Where("order_id IN ?", id_orders).
			Pluck("company_id", &ids)
		file, _ := json.MarshalIndent(campaigns, "", " ")
		_ = ioutil.WriteFile("campaigns.json", file, 0644)
	} else {
		file, _ := ioutil.ReadFile("campaigns.json")
		_ = json.Unmarshal([]byte(file), &campaigns)
	}

	if cacheMode != "1" {
		common.DB.
			Where("id IN ?", ids).
			Find(&companies)
		file, _ := json.MarshalIndent(companies, "", " ")
		_ = ioutil.WriteFile("companies.json", file, 0644)
	} else {
		file, _ := ioutil.ReadFile("companies.json")
		_ = json.Unmarshal([]byte(file), &companies)
	}

	if cacheMode != "1" {
		common.DB.
			Where("id IN ?", id_users).
			Find(&users)
		file, _ := json.MarshalIndent(users, "", " ")
		_ = ioutil.WriteFile("users.json", file, 0644)
	} else {
		file, _ := ioutil.ReadFile("users.json")
		_ = json.Unmarshal([]byte(file), &users)
	}

	fmt.Println("Found campaigns: ", len(campaigns))
	fmt.Println("Found companies: ", len(companies))
	fmt.Println("Found orders: ", len(orders))
	fmt.Println("Found users: ", len(users))

	tmpCmpNumber := 0

	for cind, campaign := range campaigns {

		now1 := time.Now()

		maxtv_company_campaigns.ProcessCampaignData(&campaign, &beginOfTheMonth)

		company := getCompanyById(companies, campaign.CompanyId)

		order := getOrderById(orders, campaign.OrderId)
		order.ProcessingOrder(&company)

		user := getUserById(users, order.SalePerson)

		fmt.Println("\n--------------------------------------------")
		fmt.Println(cind+1, " of ", len(campaigns))
		fmt.Println("Finds took :", time.Now().Sub(now1))
		fmt.Println("Cmp ID : ", campaign.ID)
		fmt.Println("Cmp Start Date : ", campaign.StartDate.Format("2006-01-02"))
		fmt.Println("Cmp End Date : ", campaign.EndDate.Format("2006-01-02"))
		fmt.Println("Cmp Status : ", campaign.Status)
		fmt.Println("Cmp Order Id : ", campaign.OrderId)
		fmt.Println("Cmp Parent Id : ", campaign.ParentId)
		fmt.Println("Cmp Sales Id : ", order.SalePerson)
		fmt.Println("Cmp Link : ", campaign.LinkToCampaign)
		fmt.Println("Cmp Order Link : ", campaign.LinkToOrder)

		rec := WeeklyReportDetails{Campaign: campaign, Order: order, User: user, Company: company}
		distributePricePerCampaign(&rec, order.RealAmount, campaigns)

		//fmt.Println("Cmp Amount Per Campaign : ", rec.PricePerCampaign)
		//fmt.Println("Cmp Amount Per Campaign : ", rec.PricePerCampaign)
		fmt.Println("Cmp Count : ", rec.NumberOfCampaigns)
		fmt.Println("Cmp Length : ", campaign.CampaignLength)
		fmt.Println("Cmp Order Amount : ", order.RealAmount)
		fmt.Println("Cmp Number of Days : ", rec.NumberOfDays)
		fmt.Println("Cmp Price Per Days : ", rec.PricePerDay)

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
		beginOfTheMonth.Format("January 2006") + "-" + beginOfTheMonth.AddDate(0, 11, 0).Format("January 2006"),
	})

	numbersStyle, _ := f.NewStyle(&excelize.Style{
		NumFmt: 2,
	})
	border := excelize.Border{
		Type:  "1",
		Color: "000000",
		Style: 6,
	}
	var borders []excelize.Border
	borders = append(borders, border)
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Border:    borders,
		Font:      &excelize.Font{Bold: true, Size: 14},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		NumFmt:    2,
	})

	f.SetRowStyle(SheetName, 1, 1, titleStyle)
	f.SetRowStyle(SheetName, 2, 2, titleStyle)
	f.SetRowStyle(SheetName, 3, 3, titleStyle)

	f.MergeCell(SheetName, "A1:O1", "")
	f.MergeCell(SheetName, "A2:O2", "")
	f.MergeCell(SheetName, "A3:O3", "")

	cRowIndex++
	cRowIndex++

	header := []interface{}{"Sales Person"}
	for i := 0; i < 12; i++ {
		m := beginOfTheMonth.AddDate(0, i, 0)
		header = append(header, m.Format("Jan. 2006"))
	}
	header = append(header, "Total")
	header = append(header, "Percentage of sales")
	f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &header)
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 14}, Alignment: &excelize.Alignment{
			Horizontal: "center",
			WrapText:   true,
		},
	})
	f.SetRowStyle(SheetName, cRowIndex, cRowIndex, headerStyle)
	f.SetColWidth(SheetName, "A", "A", 30)
	f.SetColWidth(SheetName, "B", "O", 12)
	cRowIndex += 1

	grandTotal := 0.0
	percentageRowIndexes := []PercentageRow{}
	globalMonthSums := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, user := range users {

		sumBySales := 0.0
		salesWr := getSalesWr(wr, user)
		uname := user.Firstname + " " + user.Lastname
		//uname := user.Firstname + " " + user.Lastname + " (" + strconv.Itoa(user.Id) + ")" + " [" + strconv.Itoa(len(salesWr)) + "]"
		outRow := []interface{}{uname}
		for i := 0; i < 12; i++ {

			sumByMonth := 0.0
			for _, row := range salesWr {
				if row.Order.Details.Amount <= 1 {
					continue
				}
				mb := beginOfTheMonth.AddDate(0, i, 0)
				daysInMonth := Date(mb.Year(), int(mb.AddDate(0, 1, 0).Month()), 0).Day()
				me := mb.AddDate(0, 0, daysInMonth).Add(-1 * time.Second)

				sd := row.Campaign.StartDate
				ed := row.Campaign.EndDate

				if sd.Unix() <= mb.Unix() && ed.Unix() >= me.Unix() {
					sumByMonth += row.PricePerDay * float64(daysInMonth)
					continue
				}

				if sd.Unix() >= mb.Unix() && sd.Unix() <= me.Unix() {
					campaignDays := daysInMonth - sd.Day() + 1
					sumByMonth += row.PricePerDay * float64(campaignDays)
					continue
				}

				if ed.Unix() >= mb.Unix() && ed.Unix() <= me.Unix() {
					campaignDays := sd.Day()
					sumByMonth += row.PricePerDay * float64(campaignDays)
					continue
				}
				if sd.Unix() >= mb.Unix() && sd.Unix() <= me.Unix() && ed.Unix() >= mb.Unix() && ed.Unix() <= me.Unix() {
					campaignDays := ed.Day() - sd.Day() + 1
					sumByMonth += row.PricePerDay * float64(campaignDays)
					continue
				}

				//if sd.Equal(mb) && sd.Before(me) && ed.After(mb) && ed.Equal(me) {
				//	campaignDays := ed.Day() - sd.Day() + 1
				//	sumByMonth += row.PricePerDay * float64(campaignDays)
				//	continue
				//}

			}
			outRow = append(outRow, sumByMonth)
			globalMonthSums[i] += sumByMonth
			sumBySales += sumByMonth
		}

		if sumBySales == 0 {
			continue
		}
		outRow = append(outRow, sumBySales)
		grandTotal += sumBySales
		f.SetSheetRow(SheetName, fmt.Sprintf("A%d", cRowIndex), &outRow)

		f.SetCellStyle(SheetName, fmt.Sprintf("B%d", cRowIndex), fmt.Sprintf("O%d", cRowIndex), numbersStyle)
		percentageRowIndexes = append(percentageRowIndexes, PercentageRow{SumBySales: sumBySales, INdex: cRowIndex})
		cRowIndex += 1

		outRow = []interface{}{"Order Number", "Account Name", "Number of Campaign", "Total Amount", "Price per day", "Sales", "Campaign Length", "Campaign Start Date", "Campaign End Date", "Link To Campaign", "Link To Order"}
		f.SetSheetRow(SheetName, fmt.Sprintf("B%d", cRowIndex), &outRow)
		f.SetRowOutlineLevel(SheetName, cRowIndex, 1)
		st, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Bold: true}})
		f.SetRowStyle(SheetName, cRowIndex, cRowIndex, st)
		cRowIndex += 1
		for _, row := range salesWr {
			outRow = []interface{}{
				row.Order.OrderNumber,
				row.Company.Name,
				row.NumberOfCampaigns,
				row.Order.RealAmount,
				row.PricePerDay,
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

	var tgts float64
	for _, gts := range globalMonthSums {
		tgts += gts
	}
	globalMonthSums = append(globalMonthSums, tgts)
	f.SetSheetRow(SheetName, fmt.Sprintf("B%d", cRowIndex), &globalMonthSums)
	f.SetCellStyle(SheetName, fmt.Sprintf("B%d", cRowIndex), fmt.Sprintf("O%d", cRowIndex), titleStyle)

	for _, ri := range percentageRowIndexes {
		percentage := ri.SumBySales / grandTotal
		f.SetCellValue(SheetName, fmt.Sprintf("O%d", ri.INdex), percentage)
		st, _ := f.NewStyle(&excelize.Style{
			NumFmt: 10,
		})
		f.SetCellStyle(SheetName, fmt.Sprintf("O%d", ri.INdex), fmt.Sprintf("O%d", ri.INdex), st)
	}

	reportName := "./data/weekly_sales_reports/WSR_ym_" + yearMode + "_" + year + "_" + time.Now().Format("2006-01-02_15:04:05") + ".xlsx"
	err := f.SaveAs(reportName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The total number of campaigns :", len(campaigns))
	fmt.Println("It took :", time.Now().Sub(now))

	return reportName
}

func getSalesWr(wr []WeeklyReportDetails, user db_interface.MaxtvUser) []WeeklyReportDetails {
	result := []WeeklyReportDetails{}
	for _, row := range wr {
		if row.User.Id == user.Id {
			result = append(result, row)
		}
	}
	return result
}

func getUserById(users []db_interface.MaxtvUser, person int) db_interface.MaxtvUser {
	for _, user := range users {
		if user.Id == person {
			return user
		}
	}
	return db_interface.MaxtvUser{}
}

func getOrderById(orders []db_interface.MaxtvCompanyOrder, id int) db_interface.MaxtvCompanyOrder {
	for _, order := range orders {
		if order.Id == id {
			return order
		}
	}
	return db_interface.MaxtvCompanyOrder{}
}

func getCompanyById(companies []db_interface.MaxtvCompanie, id int) db_interface.MaxtvCompanie {
	for _, c := range companies {
		if c.Id == id {
			return c
		}
	}
	return db_interface.MaxtvCompanie{}
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func distributePricePerCampaign(d *WeeklyReportDetails, amount float64, campaigns []db_interface.MaxtvCompanyCampaign) {

	var cmps []db_interface.MaxtvCompanyCampaign

	for _, cmp := range campaigns {
		if cmp.OrderId == d.Campaign.OrderId {
			cmps = append(cmps, cmp)
		}
	}
	daysCount := 0
	overallDisplayCount := 0
	for _, cmp := range cmps {
		maxtv_company_campaigns.ProcessCampaignData(&cmp, nil)
		//fmt.Println("> Link to campaign :", cmp.LinkToCampaign)
		overallDisplayCount += len(cmp.Displays)
		daysCount += cmp.CampaignLength
	}

	d.NumberOfCampaigns = len(cmps)
	d.NumberOfDays = daysCount
	//fmt.Println("Campaign count", len(cmps))

	switch len(cmps) {
	case 1:
		d.PricePerCampaign = amount
		d.NumberOfCampaigns = len(cmps)
		d.PricePerDay = float64(amount) / float64(daysCount)
	case 0:
		panic("No campaigns found for order_id = " + strconv.Itoa(d.Campaign.OrderId))
	default:
		if daysCount == 0 {
			panic("Days count is 0")
		}
		d.PricePerCampaign = float64(amount) * float64(d.Campaign.CampaignLength) / float64(daysCount)
		//d.PricePerCampaign = float64(amount) * float64(len(d.Campaign.Displays)) / float64(overallDisplayCount)
		d.NumberOfCampaigns = len(cmps)
		d.PricePerDay = float64(amount) / float64(daysCount)
	}

}

func GetBeginOfTheMonth() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func GetBeginOfTheYear(year string) time.Time {
	var y int
	var err error
	if year != "" {
		y, err = strconv.Atoi(year)
		if err != nil {
			y = time.Now().Year()
		}
	} else {
		y = time.Now().Year()
	}

	return time.Date(y, 1, 1, 0, 0, 0, 0, time.Now().Location())
}
