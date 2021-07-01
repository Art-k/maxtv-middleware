package db_interface

import "time"

type MaxtvCompanyCampaignMedia struct {
	ID                     int       `gorm:"column:id" json:"id"`
	MaxtvCompanyCampaignID int       `gorm:"column:campaign_id" json:"campaign_id"`
	Title                  string    `gorm:"column:title" json:"title"`
	Path                   string    `gorm:"column:path" json:"path"`
	Type                   string    `gorm:"column:type" json:"type"`
	Size                   int       `gorm:"column:size" json:"size"`
	DurationMs             int       `gorm:"column:duration_ms" json:"duration_ms"`
	Date                   time.Time `gorm:"column:date" json:"date"`
	Active                 bool      `gorm:"column:active" json:"active"`
	Preview                bool      `gorm:"column:preview" json:"preview"`
	UploadedBy             int       `gorm:"column:uploaded_by" json:"uploaded_by"`
	Uodated                time.Time `gorm:"column:updated" json:"updated"`
	ActivedOn              time.Time `gorm:"column:actived_on" json:"actived_on"`
	PreviewOn              time.Time `gorm:"column:preview_on" json:"preview_on"`
	ActivedBy              int       `gorm:"column:actived_by" json:"actived_by"`
	PreviewBy              int       `gorm:"column:preview_by" json:"preview_by"`
}

func (MaxtvCompanyCampaignMedia) TableName() string {
	return "maxtv_company_campaign_media"
}
