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
	Name string `json:"name" validate:"required"`

	BankAccountName *string `json:"bankAccountName" validate:"omitempty"`

	BankAccountNumber *string `json:"bankAccountNumber" validate:"omitempty,len=10"`

	BVN *string `json:"bvn" validate:"omitempty,len=11"`

	BankCode *string `json:"bankCode" validate:"omitempty"`

	OwnerId *string `json:"ownerId" validate:"required"`
}

type UpdateBankKycDTO struct {
	Id string `json:"id" validate:"required"`

	Name *string `json:"name" validate:"omitempty"`

	BankAccountName *string `json:"bankAccountName" validate:"omitempty"`

	BankAccountNumber *string `json:"bankAccountNumber" validate:"omitempty,len=10"`

	BVN *string `json:"bvn" validate:"omitempty,len=11"`

	BankCode *string `json:"bankCode" validate:"omitempty"`
}

type DeleteBankKycDTO struct {
	Id string `json:"id" validate:"required"`
}
