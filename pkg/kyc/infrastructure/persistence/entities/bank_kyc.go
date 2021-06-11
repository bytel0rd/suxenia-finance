package entities

import (
	"suxenia-finance/pkg/common/persistence"
)

type BankingKycEntity struct {
	Id string

	Name string

	BankAccountName string

	BankAccountNumber string

	BankCode string

	BVN string

	OwnerId string

	verified bool

	persistence.AuditInfo
}
