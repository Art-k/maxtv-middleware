package common

import (
	"strconv"
	"time"

	//"fmt"
	//"github.com/elliotchance/phpserialize"
	"github.com/gin-gonic/gin"
	"strings"
	//php "github.com/kovetskiy/go-php-serialize"
)

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func String2Order(serialised string) (oi TOrderInfo) {

	recs := strings.Split(strings.Split(serialised, "{")[1], ";")
	//fmt.Println(recs)

	for i := 0; i < len(recs)-2; i = i + 2 {
		key := strings.Split(recs[i], ":")
		//fmt.Print(key[2]+" | ")
		val := strings.Split(recs[i+1], ":")
		//fmt.Println(val)

		sw := strings.Replace(key[2], "\"", "", -1)
		switch sw {
		case "order_id":
			oi.OrderId = strings.Replace(val[2], "\"", "", -1)
		case "payments":
			oi.Payments, _ = strconv.Atoi(strings.Replace(val[2], "\"", "", -1))
		case "first_last_payment":
			oi.FirstLastPayment, _ = strconv.Atoi(val[1])
		case "include_design_fee":
			oi.IncludeDesignFee, _ = strconv.Atoi(val[1])
		case "amounts":
			oi.Amount, _ = strconv.ParseFloat(strings.Replace(val[2], "\"", "", -1), 64)
		//case "payments_first":
		//	oi.PaymentFirst = strings.Replace(val[2], "\"", "", -1)
		case "payments_start":
			oi.PaymentStart, _ = time.Parse("02-01-2006", strings.Replace(val[2], "\"", "", -1))
		case "payments_inc":
			oi.PaymentIncrement, _ = strconv.Atoi(strings.Replace(val[2], "\"", "", -1))
		case "payments_inc_type":
			oi.PaymentIncrementType = strings.Replace(val[2], "\"", "", -1)
		case "method":
			oi.Method = strings.Replace(val[2], "\"", "", -1)
		case "design_fee":
			oi.DesignFee, _ = strconv.ParseFloat(strings.Replace(val[2], "\"", "", -1), 64)
		case "currency":
			oi.Currency = strings.Replace(val[2], "\"", "", -1)
		case "tax":
			oi.Tax, _ = strconv.ParseFloat(strings.Replace(val[2], "\"", "", -1), 64)
		case "copied":
			oi.Copied, _ = strconv.Atoi(val[1])
			//case "method_other":
			//	oi.OrderId = strings.Replace(val[2], "\"", "", -1)
		}

	}

	return oi
}

func GetIntParameter(c *gin.Context, param string) (int, error) {

	paramStr := c.Param(param)

	paramInt, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, err
	}

	return paramInt, nil
}
