package db_interface

type MaxtvUser struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	FirstName   string `gorm:"columns:firstname" json:"first_name"`
	LastName    string `gorm:"columns:lastname" json:"last_name"`
	Active      bool   `json:"active"`
	AccessLevel int    `json:"access_level"`
	ApiToken    string `json:"-" gorm:"column:api_token"`
}
