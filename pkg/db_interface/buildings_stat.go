package db_interface

import "time"

type BuildingStat struct {
	Id                        int       `json:"id"`
	BuildingId                int       `json:"building_id"`
	MedianAge                 string    `json:"median_age"`
	Newcomers                 string    `json:"new_comers"`
	AvgHouseholdSize          string    `json:"avg_house_hold_size"`
	RenovatedInPastYear       string    `json:"renovated_in_past_year"`
	AvgHouseholdIncome        string    `json:"avg_house_hold_income"`
	WorkingResidents          string    `json:"working_residents"`
	HouseholdsWithChildren    string    `json:"house_hold_with_children"`
	TurnoverRate              string    `json:"turnover_rate"`
	CondoPetFriendly          string    `json:"condo_pet_friendly"`
	CarsPerHousehold          string    `json:"cars_per_house_hold"`
	PopulationSize            string    `json:"population_size"`
	NumberOfHouseholds        string    `json:"number_of_households"`
	HouseholdsWithoutChildren string    `json:"households_without_children"`
	NotInLabourForce          string    `json:"not_in_labour_force"`
	ChangeDate                time.Time `json:"change_date"`
}

func (BuildingStat) TableName() string {
	return "buildings_stat"
}
