package db_interface

import "time"

type OrderDetails struct {
	OrderId              string    `json:"order_id"`
	Payments             int       `json:"payments"`
	FirstLastPayment     int       `json:"first_last_payment"`
	IncludeDesignFee     int       `json:"include_design_fee"`
	Amount               float64   `json:"amount"`
	PaymentStart         time.Time `json:"payment_start"`
	PaymentIncrement     int       `json:"payment_increment"`
	PaymentIncrementType string    `json:"payment_increment_type"`
	Method               string    `json:"method"`
	DesignFee            float64   `json:"design_fee"`
	Currency             string    `json:"currency"`
	Tax                  float64   `json:"tax"`
	Copied               int       `json:"copied"`
}

type MaxtvCompanyOrder struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	OrderNumber    string    `json:"order_number"`
	CompanyId      int       `json:"company_id"`
	Payments       string    `json:"payments"`
	SaleDate       time.Time `json:"sale_date"`
	SalePersonId   int       `gorm:"sale_person" json:"sales_person_id"`
	AdType         string    `gorm:"ad_type" json:"ad_type"`
	Invoice        string    `gorm:"invoice" json:"invoice"`
	Network        string    `gorm:"network" json:"network"`
	OrderType      string    `gorm:"type" json:"order_type"`
	TermLength     int       `gorm:"term_length" json:"term_length"`
	TermLengthType string    `gorm:"term_length_type" json:"term_length_type"`

	Details       OrderDetails `gorm:"-" json:"details"`
	LinkToCompany string       `gorm:"-" json:"link_to_company"`
	LinkToOrder   string       `gorm:"-" json:"link_to_order"`
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
