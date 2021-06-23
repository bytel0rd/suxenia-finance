package aggregates

import (
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var wallet WalletAggregate = WalletAggregate{
	id:               2,
	totalBalance:     decimal.NewFromInt(5),
	availableBalance: decimal.NewFromInt(5),
	version:          0,
	ownerId:          "",
	modified:         false,
	AuditData:        objects.AuditData{},
}

func ShouldThrowErrorWhileSettingAvailableBalance(t *testing.T) {

	err := wallet.SetAvailableBalance(decimal.NewFromInt(10))

	assert.Error(t, err)

}
