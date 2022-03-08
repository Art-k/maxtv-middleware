package db_interface

import "time"

type MaxtvBuilding struct {
	Id                int       `json:"id"`
	Network           string    `json:"network"`
	Name              string    `json:"name"`
	Address           string    `json:"address"`
	PostalCode        string    `json:"postal_code"`
	CompanyId         int       `json:"company_id"`
	MccId             string    `json:"mcc_id"`
	InstallationDate  time.Time `json:"installation_date" gorm:"column:instalation_date"`
	ShowOnMap         bool      `json:"show_on_map"`
	Ratecard          int       `json:"ratecard"`
	RatecardType      int       `json:"ratecard_type"`
	City              string    `json:"city"`
	State             string    `json:"state"`
	Country           string    `json:"country"`
	CorporationNumber string    `json:"corporation_number"`
	GeoLocation       string    `json:"geo_location"`
	Visitors          int       `json:"visitors"`
	Residents         int       `json:"residents"`
	Units             int       `json:"units"`
}
