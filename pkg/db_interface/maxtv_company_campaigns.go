package db_interface

import "time"

type MaxtvCompanyCampaign struct {
	ID           int       `gorm:"id" json:"id"`
	CompanyId    int       `gorm:"company_id" json:"company_id"`
	Status       string    `gorm:"status" json:"status"`
	StartDate    time.Time `gorm:"start_date" json:"start_date"`
	EndDate      time.Time `gorm:"end_date" json:"end_date"`
	CampaignType string    `gorm:"column:type" json:"campaign_type"` // 'primary', 'secondary'
	ParentId     int       `gorm:"column:parent_id" json:"parent_id"`
	Title        string    `gorm:"title" json:"title"`
	OrderId      int       `gorm:"order_id" json:"order_id"`
	CreatedOn    time.Time `gorm:"created_on" json:"created_on"`
	AdType       string    `gorm:"ad_type" json:"ad_type"`
	ShortUrl     string    `gorm:"short_url" json:"short_url"`

	LinkToCampaign         string `gorm:"-" json:"link_to_campaign"`
	LinkToCompany          string `gorm:"-" json:"link_to_company"`
	LinkToOrder            string `gorm:"-" json:"link_to_order"`
	LinkToImpressionReport string `gorm:"-" json:"link_to_impression_report"`
	LinkToStatJson         string `gorm:"-" json:"link_to_stat_json"`

	CampaignLength int `gorm:"-" json:"campaign_length"`
	PastDays       int `gorm:"-" json:"past_days"`
	RemainingDays  int `gorm:"-" json:"remaining_days"`

	Media    []MaxtvCompanyCampaignMedia    `gorm:"foreignKey:MaxtvCompanyCampaignID" json:"media"`
	Displays []MaxtvCompanyCampaignDisplay  `gorm:"foreignKey:CampaignID" json:"displays"`
	Schedule []MaxtvCompanyCampaignSchedule `gorm:"foreignKey:CampaignID" json:"schedule"`

	EmailPsdAdDraft     time.Time `gorm:"email_psd_ad_draft" json:"email_psd_ad_draft"`       //datetime                              not null,
	EmailAdDraft        time.Time `gorm:"email_ad_draft" json:"email_ad_draft"`               //datetime                              not null,
	EmailYourAdIsUp     time.Time `gorm:"email_your_ad_is_up" json:"email_your_ad_is_up"`     //datetime                              not null,
	EmailUpdateDesign   time.Time `gorm:"email_update_design" json:"email_update_design"`     //datetime                              not null,
	EmailPauseAd        time.Time `gorm:"email_pause_ad" json:"email_pause_ad"`               //datetime                              not null,
	EmailResumeAd       time.Time `gorm:"email_resume_ad" json:"email_resume_ad"`             //datetime                              not null,
	EmailUpdateBuilding time.Time `gorm:"email_update_building" json:"email_update_building"` //datetime                              not null,
	EmailReport         time.Time `gorm:"email_report" json:"email_report"`                   //datetime                              not null,
	EmailManualEmail    time.Time `gorm:"email_manual_email" json:"email_manual_email"`       //datetime                              not null,
	PsdShortUrl         string    `gorm:"psd_short_url" json:"psd_short_url"`                 //varchar(255)                          not null,
	ActiveMedia         string    `gorm:"active_media" json:"active_media"`                   //varchar(255)                          not null,

	//primary key,
	//network                  varchar(50)           default 'maxtv' not null,
	//designer_id              int                                   not null,
	//contract_type            int                                   not null,
	//exclusive                int                                   not null,
	//campaign_length          int                                   not null,
	//campaign_length_type     enum ('months', 'weeks', 'days')      not null,
	//campaign_lenth_manual    enum ('0', '1')                       not null,
	//spot                     int                                   not null,
	//start_date               datetime                              not null,
	//end_date                 datetime                              not null,
	//manual_dates             int                                   not null,
	//note                     text                                  not null,
	//note_info                text                                  not null,
	//note_info_history        text                                  not null,
	//email_art_request        datetime                              not null,
	//parent_id                int                                   not null,
	//ActiveFrom time.Time `gorm:"active_from"`
	//ActiveTo   time.Time `gorm:"active_to"`
	//animation                int                                   not null,
	//artwork_type             varchar(100)          default ''      not null,
	//number_of_changes        int                                   not null,
	//watermark                int                                   not null,
	//tracking                 int                                   not null,
	//designer_1               datetime                              not null,
	//designer_2               datetime                              not null,
	//designer_3               datetime                              not null,
	//art_agent                int                                   null,
	//workflow_priority        int                                   not null,
	//created_on               datetime                              not null,
	//estimated_start_date     datetime                              null,
	//category_id              int                   default 0       not null,
	//service_workflow_status  int                   default 0       not null,
	//representative_notify    tinyint               default 0       not null,
	//design_template_id       int                                   null,
	//google_lineitem_id       varchar(100)                          not null,
	//banner_url               varchar(255)                          not null,
	//targeting                text                                  not null,
	//fill_loop                tinyint(1)            default 0       not null,
	//building_list_is_changed tinyint(1)
}
