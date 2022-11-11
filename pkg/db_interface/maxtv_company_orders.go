package db_interface

import (
	"fmt"
	"maxtv_middleware/pkg/common"
	"strconv"
	"strings"
	"time"
)

type OrderDetails struct {
	OrderId              string    `json:"order_id"`
	Payments             int       `json:"payments"`
	FirstLastPayment     int       `json:"first_last_payment"`
	IncludeDesignFee     int       `json:"include_design_fee"`
	Amount               float64   `json:"amount"`
	PaymentStart         time.Time `json:"payment_start"`
	PaymentIncrement     int       `json:"payment_increment"`
	PaymentIncrementType string    `json:"payment_increment_type"`
	Method               string    `json:"method"`
	DesignFee            float64   `json:"design_fee"`
	Currency             string    `json:"currency"`
	Tax                  float64   `json:"tax"`
	Copied               int       `json:"copied"`
}

type MaxtvCompanyOrder struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	OrderNumber    string    `json:"order_number"`
	CompanyId      int       `json:"company_id"`
	Payments       string    `json:"payments"`
	SaleDate       time.Time `json:"sale_date"`
	SalePerson     int       `json:"sales_person_id"`
	AdType         string    `gorm:"ad_type" json:"ad_type"`
	Invoice        string    `gorm:"invoice" json:"invoice"`
	Network        string    `gorm:"network" json:"network"`
	OrderType      string    `gorm:"type" json:"order_type"`
	TermLength     int       `gorm:"term_length" json:"term_length"`
	TermLengthType string    `gorm:"term_length_type" json:"term_length_type"`
	Source         int       `json:"source"`

	Details       OrderDetails `gorm:"-" json:"details"`
	LinkToCompany string       `gorm:"-" json:"link_to_company"`
	LinkToOrder   string       `gorm:"-" json:"link_to_order"`
}

func (order *MaxtvCompanyOrder) ProcessingOrder() {

	now := time.Now()

	order.LinkToOrder = "https://maxtvmedia.com/cms/?a=211&tab=orders&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(order.CompanyId) +
		"&order_id=" + strconv.Itoa(order.Id)
	order.LinkToCompany = "https://maxtvmedia.com/cms/?a=211&tab=details&type=account&fullview=1" +
		"&company_id=" + strconv.Itoa(order.CompanyId)

	tmp := common.TruncateString(order.Payments, strings.Index(order.Payments, "{"), strings.Index(order.Payments, "}"))
	//fmt.Println("\n", tmp)

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

	fmt.Println("ProcessingOrder took :", time.Now().Sub(now))

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

//
//id                       int auto_increment
//primary key,
//title                    varchar(255)                                not null,
//order_number             varchar(255)                                not null,
//company_id               int                                         not null,
//`order`                  longblob                                    not null,
//type                     varchar(255)                                not null,
//boof_type                varchar(255)                                null,
//date                     datetime                                    not null,
//sale_date                date                                        not null,
//sale_person              int                                         not null,
//telemarketer_person      int                                         null,
//payments                 text                                        not null,
//source                   int                                         not null,
//ad_type                  enum ('S', 'B', 'SB', 'MS') default 'S'     not null,
//collected                int                                         not null,
//charged                  int                                         not null,
//invoice                  varchar(255)                                not null,
//order_type               varchar(255)                                not null,
//to_collection_report     tinyint                     default 0       not null,
//google_order_id          varchar(100)                                not null,
//auto_billing             tinyint                     default 0       not null,
//bambora_transaction_id   varchar(100)                default ''      not null,
//bambora_transaction_date datetime                                    null,
//network                  enum ('maxtv', 'mcc)
