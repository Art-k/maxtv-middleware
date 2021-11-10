package db_interface

import "time"

type UserCosts struct {
	Id              int       `json:"id"`
	UserId          int       `json:"user_id"`
	Salary          float64   `json:"salary"`
	RateOffice      float64   `json:"rate_office"`
	RateSelf        float64   `json:"rate_self"`
	BuildingFeeRate float64   `json:"building_fee_rate"`
	DateFrom        time.Time `json:"date_from"`
	CreatedOn       time.Time `json:"created_on"`
	CreatedBy       int       `json:"created_by"`
	Deleted         int       `json:"deleted"`
	BRateOffice     float64   `json:"b_rate_office"`
	BRateSelf       float64   `json:"b_rate_self"`
	Cost            float64   `json:"cost"`
	RateHourly      float64   `json:"rate_hourly"`
	SalaryType      string    `json:"salary_type"`
}

func (UserCosts) TableName() string {
	return "maxtv_users_cost"
}

//id                int unsigned auto_increment,
//user_id           int unsigned                                               not null,
//salary            float unsigned                                             not null,
//rate_office       float unsigned                                             not null,
//rate_self         float unsigned                                             not null,
//building_fee_rate float unsigned                                             not null,
//date_from         datetime                                                   not null,
//created_on        datetime                         default CURRENT_TIMESTAMP null,
//created_by        int unsigned                                               not null,
//deleted           smallint unsigned                                          not null,
//b_rate_office     float unsigned                                             not null,
//b_rate_self       float unsigned                                             not null,
//cost              float unsigned                                             not null,
//rate_hourly       float unsigned                                             not null,
//salary_type       enum ('advertising', 'building') default 'advertising'     not null,
