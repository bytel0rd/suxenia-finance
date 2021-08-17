package drivers

import (
	"suxenia-finance/pkg/common/infrastructure/logs"

	"github.com/google/wire"
)

var BuildSet wire.ProviderSet = wire.NewSet(
	NewWalletDriver,
	NewPaymentDriver,
	NewWalletTransactionDriver,
	logs.NewLogger,
	NewWithdrawalDriver,
)
