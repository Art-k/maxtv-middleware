package db_interface

type MaxtvBuildingDisplayResidentTraffic struct {
	ID                int `json:"id"`
	DisplayId         int `json:"display_id"`
	Traffic           int `json:"traffic"`
	DwellTime         int `json:"dwell_time"`
	ImpressionsPerDay int `json:"impressions_per_day"`
}

func (MaxtvBuildingDisplayResidentTraffic) TableName() string {
	return "maxtv_building_display_resident_traffic"
}
