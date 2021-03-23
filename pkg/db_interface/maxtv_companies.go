package db_interface

type MaxtvCompanie struct {
	Id          int
	ParentId    int
	BuildingId  int
	IsBuilding  int
	AccountType string `gorm:"column:type"`
	Name        string
	Address     string
}
