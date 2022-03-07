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

type ResponseHeader struct {
	Total          int64   `json:"total"`
	ResponseTook   float64 `json:"response-took-seconds"`
	ProcessingTook float64 `json:"processing-took-seconds"`
}

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
