package aggregates

import "suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

type AccountValidator func(account *BankingKYC) (*string, bool, error)

type BVNValidator func(account *BankingKYC, bvn string) (*string, bool, error)

type VirtualAccountProvider func(bankKyc *BankingKYC) (*entities.VirtualAccountEntity, error)
