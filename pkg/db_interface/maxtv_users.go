package db_interface

type MaxtvUser struct {
	Id        int    `gorm:"id"`
	FirstName string `gorm:"firstname"`
	LastName  string `gorm:"lastname"`
	Email     string `gorm:"email"`
	Active    bool   `gorm:"active"`
}
