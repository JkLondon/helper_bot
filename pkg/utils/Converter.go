package utils

import (
	"github.com/shopspring/decimal"
)

func StrToInt64(x string) (number int64, err error) {
	numberDecimal, err := decimal.NewFromString(x)
	if err != nil {
		return 0, err
	}
	number = numberDecimal.IntPart()
	return number, nil
}
