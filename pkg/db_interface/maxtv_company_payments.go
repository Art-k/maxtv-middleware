package db_interface

import "time"

type MaxtvCompanyPayment struct {
	Id          int
	CompanyId   int
	OrderId     int
	Amount      float64
	Date        time.Time
	Depositedon time.Time
	Status      string
	Charged     int
	Collected   int
}
