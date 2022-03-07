package db_interface

type MaxtvCompanyCampaignSchedule struct {
	ID         int    `gorm:"column:id" json:"id"`
	CampaignID int    `gorm:"column:campaign_id" json:"campaign_id"`
	Hours      string `gorm:"column:hours" json:"hours"`
	Shows      string `gorm:"column:shows" json:"shows"`
}

func (MaxtvCompanyCampaignSchedule) TableName() string {
	return "maxtv_company_campaign_schedule"
}
