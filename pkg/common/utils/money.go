package utils

import "github.com/shopspring/decimal"

func IntToDecimal(amount int) decimal.Decimal {
	return decimal.NewFromInt(int64(amount)).Div(decimal.NewFromInt(100))
}
