package db_interface

type MaxtvCompanyCampaignDisplay struct {
	Id         int `gorm:"id" json:"id"`
	CampaignID int `gorm:"campaign_id" json:"campaign_id"`
	DisplayID  int `gorm:"display_id" json:"display_id"`
	BuildingID int `gorm:"building_id" json:"building_id"`
}
