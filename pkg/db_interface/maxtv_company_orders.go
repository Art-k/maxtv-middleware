package db_interface

import "time"

type MaxtvCompanyOrder struct {
	Id           int
	Title        string
	OrderNumber  string
	CompanyId    int
	Payments     string
	SaleDate     time.Time
	SalePersonId int    `gorm:"sale_person"`
	AdType       string `gorm:"ad_type"`
	Invoice      string `gorm:"invoice"`
	Network      string `gorm:"network"`
	OrderType    string `gorm:"type"`
}

//
//id                       int auto_increment
//primary key,
//title                    varchar(255)                                not null,
//order_number             varchar(255)                                not null,
//company_id               int                                         not null,
//`order`                  longblob                                    not null,
//type                     varchar(255)                                not null,
//boof_type                varchar(255)                                null,
//date                     datetime                                    not null,
//sale_date                date                                        not null,
//sale_person              int                                         not null,
//telemarketer_person      int                                         null,
//payments                 text                                        not null,
//source                   int                                         not null,
//ad_type                  enum ('S', 'B', 'SB', 'MS') default 'S'     not null,
//collected                int                                         not null,
//charged                  int                                         not null,
//invoice                  varchar(255)                                not null,
//order_type               varchar(255)                                not null,
//to_collection_report     tinyint                     default 0       not null,
//google_order_id          varchar(100)                                not null,
//auto_billing             tinyint                     default 0       not null,
//bambora_transaction_id   varchar(100)                default ''      not null,
//bambora_transaction_date datetime                                    null,
//network                  enum ('maxtv', 'mcc)
