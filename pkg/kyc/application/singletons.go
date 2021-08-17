package application

import (
	"suxenia-finance/pkg/kyc/infrastructure/persistence/repos"

	"github.com/google/wire"
)

var BuildSet wire.ProviderSet = wire.NewSet(repos.BuildSet, NewBankingKycApplication)
