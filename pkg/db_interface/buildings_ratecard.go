package db_interface

import "time"

type BuildingRatecard struct {
	Id         int
	BuildingId int
	Name       string
	Value      string
	ChangeDate time.Time
}

type Tabler interface {
	TableName() string
}

func (BuildingRatecard) TableName() string {
	return "buildings_ratecard"
}
