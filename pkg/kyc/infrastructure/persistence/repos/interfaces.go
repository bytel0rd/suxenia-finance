package repos

import (
	"suxenia-finance/pkg/common/structs"
	kycAggregates "suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/drivers"

	"github.com/google/wire"
)

type IBankingKycRepo interface {
	RetrieveById(id string) (*kycAggregates.BankingKYC, bool, *structs.DBException)

	Create(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, *structs.DBException)

	Update(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, *structs.DBException)

	Delete(id string) (bool, *structs.DBException)
}

var BuildSet wire.ProviderSet = wire.NewSet(NewBankycRepo, drivers.BuildSet)
