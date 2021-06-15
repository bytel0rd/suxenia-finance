package repos

import kycAggregates "suxenia-finance/pkg/kyc/domain/aggregates"

type IBankingKycRepo interface {
	RetrieveById(id string) (*kycAggregates.BankingKYC, bool, error)

	Create(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, error)

	Update(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, error)

	Delete(id string) (bool, error)
}
