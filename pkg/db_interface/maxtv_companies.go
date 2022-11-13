package db_interface

type MaxtvCompanie struct {
	Id                int     `json:"id"`
	ParentId          int     `json:"parent_id"`
	BuildingId        int     `json:"building_id"`
	IsBuilding        int     `json:"is_building"`
	Tax               float64 `gorm:"tax" json:"tax"`
	AccountType       string  `gorm:"column:type" json:"account_type"`
	Name              string  `json:"name"`
	Address           string  `json:"address"`
	Owner             string  `gorm:"owner" json:"owner"`
	Email             string  `gorm:"email" json:"email"`
	EmailCustom       string  `gorm:"email_custom" json:"email_custom"`
	Phone             string  `gorm:"phone" json:"phone"`
	PhoneMobile       string  `gorm:"phone_mobile" json:"phone_mobile"`
	Assistant         string  `gorm:"assistant" json:"assistant"` //assistant              varchar(255)                                      not null,
	Manager           string  `gorm:"manager" json:"manager"`     //manager                varchar(255)
	CreatedOn         string  `gorm:"created_on" json:"created_on"`
	ExcludeFromReport bool    `gorm:"column:exclude_report" json:"exclude_from_report"`
}

//id                     int auto_increment
//primary key,
//parent_id              int                                 default 0     not null,
//currency               varchar(4)                          default 'CAD' not null,
//tax                    float                               default 13    not null,
//subtype                int                                 default 0     not null,
//building_id            int                                 default 0     null,
//name                   varchar(255)                                      not null,
//address                varchar(255)                                      not null,
//geo                    varchar(255)                                      not null,
//geo_lat                varchar(255)                                      not null,
//geo_lng                varchar(255)                                      not null,
//geo_skip               varchar(255)                                      not null,
//address_custom         text                                              not null,
//short_url              varchar(6)                                        not null,
//city                   varchar(255)                                      not null,
//phone                  text                                              not null,
//head_office_phone      varchar(255)                                      not null,
//security_desk_phone    varchar(255)                                      not null,
//phone_mobile           text                                              not null,
//email                  varchar(255)                                      not null,
//email_custom           text                                              not null,
//zip                    varchar(20)                                       not null,
//website                varchar(255)                                      not null,
//category               int                                               not null,
//type                   varchar(255)                                      not null,
//is_building            int                                               not null,
//call_in                tinyint                                           not null,
//call_in_source         varchar(50)                                       not null,
//status                 varchar(255)                                      not null,
//renewal_status         varchar(255)                                      not null,
//callback               datetime                                          not null,
//callback_note          text                                              not null,
//sn_street              varchar(255)                                      not null,
//ew_street              varchar(255)                                      not null,
//priority               int                                               not null,
//owner                  varchar(255)                                      not null,
//assistant              varchar(255)                                      not null,
//manager                varchar(255)                                      not null,
//assigned_to            int                                               not null,
//notes                  text                                              not null,
//editedon               datetime                                          not null,
//canceled               tinyint(1)                                        not null,
//payments               text                                              not null,
//payments_amount_tax    varchar(255)                                      not null,
//latest_sale_date       date                                              not null,
//src_address            varchar(255)                                      not null,
//lead_source            int                                               not null,
//request_lead           enum ('0', '1')                                   not null,
//request_lead_by        int                                               not null,
//source                 int                                               not null,
//campaign_active        int                                 default 1     not null,
//`use`                  varchar(255)                                      not null,
//phone_ext              varchar(255)                                      not null,
//units                  varchar(255)                                      not null,
//pm_company             varchar(255)                                      not null,
//zones                  int                                               not null,
//state                  varchar(255)                                      not null,
//country                varchar(255)                                      not null,
//fax                    varchar(255)                                      not null,
//property_management_id int                                               not null,
//code                   varchar(50) collate utf8_unicode_ci default ''    not null,
//company_code           varchar(255)                                      not null,
//qualified              tinyint(1)                          default 0     not null,
//created_on             datetime                                          null,
//created_by             int                                               null,
//is_bed_client          tinyint(1)                          default 0     not null,
//company_rating         tinyint(1)                          default 0     not null,
//rental                 tinyint(1)                          default 0     null,
//next_action_date       datetime                                          not null,
//google_advertiser_id   varchar(100)
