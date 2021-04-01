package common

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

var (
	Log   *logrus.Logger
	DB    *gorm.DB
	DBErr error
)

type TOrderInfo struct {
	OrderId              string
	Payments             int
	FirstLastPayment     int
	IncludeDesignFee     int
	Amount               float64
	PaymentFirst         int
	PaymentStart         time.Time
	PaymentIncrement     int
	PaymentIncrementType string
	Method               string
	DesignFee            float64
	Currency             string
	Tax                  float64
	Copied               int
	MethodOther          string
}
