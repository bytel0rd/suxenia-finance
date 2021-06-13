package dtos

import "suxenia-finance/pkg/common/persistence"

type BankingKycViewModel struct {
	Id string `json:"id"`

	Name string `json:"name"`

	BankAccountName *string `json:"bankAccountName"`

	BankAccountNumber *string `json:"bankAccountNumber"`

	BVN *string `json:"bvn"`

	BankCode *string `json:"bankCode"`

	OwnerId string `json:"ownerId"`

	Verified bool `json:"verified"`

	persistence.AuditInfo
}

type CreateBankKycDTO struct {
	Name string `json:"name"`

	BankAccountName *string `json:"bankAccountName"`

	BankAccountNumber *string `json:"bankAccountNumber"`

	BVN *string `json:"bvn"`

	BankCode *string `json:"bankCode"`

	OwnerId *string `json:"ownerId"`
}

type UpdateBankKycDTO struct {
	Id string `json:"id"`

	Name string `json:"name"`

	BankAccountName *string `json:"bankAccountName"`

	BankAccountNumber *string `json:"bankAccountNumber"`

	BVN *string `json:"bvn"`

	BankCode *string `json:"bankCode"`
}
