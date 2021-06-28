package utils

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestIsValidString(t *testing.T) {
	emptyString := ""
	validString := "Random string"

	assert.False(t, IsValidString(""))
	assert.True(t, IsValidString("Tayo Adekunle"))

	assert.False(t, IsValidStringPointer(&emptyString))
	assert.False(t, IsValidStringPointer(new(string)))

	assert.True(t, IsValidStringPointer(&validString))

}

func TestIntToDecimal(t *testing.T) {

	x_converted := decimal.NewFromInt(1).Div(decimal.NewFromInt(6))
	//IntToDecimal(1099)
	// amountInMajor := x_converted.StringFixedBank(2)

	// assert.Equal(t, amountInMajor, "100.00")

	y_converted := IntToDecimal(1099)
	// amountInMajor = y_converted.StringFixedBank(2)

	sum := x_converted.Add(y_converted)
	bad2, _ := x_converted.BigFloat().Float64()

	var bad float32 = 0.1666666666666667

	fmt.Printf("raw: %v  \nrounded: %v \n", bad, bad2)
	fmt.Printf("raw: %s  \nrounded: %s \n", x_converted.String(), x_converted.StringFixedBank(2))
	fmt.Printf("raw: %s  \nrounded: %s \n", sum.String(), sum.StringFixedBank(2))

	// assert.Equal(t, amountInMajor, "1000.10")
}
