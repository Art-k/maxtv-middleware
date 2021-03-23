package common

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log   *logrus.Logger
	DB    *gorm.DB
	DBErr error
)

type OrderInfo struct {
	OrderId              string
	Payments             string
	FirstLastPayment     int
	IncludeDesignFee     int
	Amount               string
	PaymentFirst         int
	PaymentStart         string
	PaymentIncrement     string
	PaymentIncrementType string
	Method               string
	DesignFee            string
	Currency             string
	Tax                  string
	Copied               int
	MethodOther          string
}
