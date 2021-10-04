package db_interface

type MaxtvUser struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Active      bool   `json:"active"`
	AccessLevel int    `json:"access_level"`
	ApiToken    string `json:"-" gorm:"column:api_token"`
}
