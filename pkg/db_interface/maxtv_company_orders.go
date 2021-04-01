package db_interface

import "time"

type MaxtvCompanyOrder struct {
	Id          int
	Title       string
	OrderNumber string
	CompanyId   int
	Payments    string
	SaleDate    time.Time
}
