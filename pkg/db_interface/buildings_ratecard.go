package db_interface

import "time"

type BuildingRatecard struct {
	Id         int       `json:"id"`
	BuildingId int       `json:"building_id"`
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	ChangeDate time.Time `json:"change_date"`
}

type Tabler interface {
	TableName() string
}

func (BuildingRatecard) TableName() string {
	return "buildings_ratecard"
}
