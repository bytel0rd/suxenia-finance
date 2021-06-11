package aggregates

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var wallet WalletAggregate = WalletAggregate{
	id:               2,
	totalBalance:     decimal.NewFromInt(5),
	availableBalance: decimal.NewFromInt(5),
	version:          0,
	ownerId:          "",
	createdBy:        "",
	updatedBy:        "",
	createdAt:        time.Time{},
	updateAt:         time.Time{},
	modified:         false,
}

func ShouldThrowErrorWhileSettingAvailableBalance(t *testing.T) {

	err := wallet.SetAvailableBalance(decimal.NewFromInt(10))

	assert.Error(t, err)

}
