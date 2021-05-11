package db_interface

import "time"

type BuildingStat struct {
	Id                        int
	BuildingId                int
	MedianAge                 string
	Newcomers                 string
	AvgHouseholdSize          string
	RenovatedInPastYear       string
	AvgHouseholdIncome        string
	WorkingResidents          string
	HouseholdsWithChildren    string
	TurnoverRate              string
	CondoPetFriendly          string
	CarsPerHousehold          string
	PopulationSize            string
	NumberOfHouseholds        string
	HouseholdsWithoutChildren string
	NotInLabourForce          string
	ChangeDate                time.Time
}

func (BuildingStat) TableName() string {
	return "buildings_stat"
}
