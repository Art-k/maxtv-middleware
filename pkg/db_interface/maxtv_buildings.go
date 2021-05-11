package db_interface

import "time"

type MaxtvBuilding struct {
	Id                int
	Network           string
	Name              string
	Address           string
	CompanyId         int
	MccId             string
	InstalationDate   time.Time
	ShowOnMap         bool
	Ratecard          int
	RatecardType      int
	City              string
	CorporationNumber string
}
