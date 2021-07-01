package db_interface

import "time"

type MaxtvCompanyPayment struct {
	Id          int       `json:"id"`
	CompanyId   int       `json:"company_id"`
	OrderId     int       `json:"order_id"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Depositedon time.Time `json:"deposited_on"`
	Status      string    `json:"status"`
	Charged     int       `json:"charged"`
	Collected   int       `json:"collected"`
}
