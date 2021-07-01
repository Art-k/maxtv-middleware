package db_interface

import "time"

type MaxtvBuilding struct {
	Id                int       `json:"id"`
	Network           string    `json:"network"`
	Name              string    `json:"name"`
	Address           string    `json:"address"`
	CompanyId         int       `json:"company_id"`
	MccId             string    `json:"mcc_id"`
	InstalationDate   time.Time `json:"installation_date"`
	ShowOnMap         bool      `json:"show_on_map"`
	Ratecard          int       `json:"ratecard"`
	RatecardType      int       `json:"ratecard_type"`
	City              string    `json:"city"`
	CorporationNumber string    `json:"corporation_number"`
}
