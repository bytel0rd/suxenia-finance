package application

import (
	"suxenia-finance/pkg/common/infrastructure/cache"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/repos"

	"github.com/google/wire"
)

var BuildSet wire.ProviderSet = wire.NewSet(
	NewPaymentApplication,
	NewWalletApplication,
	drivers.BuildSet,
	repos.NewWalletRepo,
	cache.NewRedisCache,
)
