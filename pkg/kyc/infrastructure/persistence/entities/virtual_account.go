package entities

import (
	"suxenia-finance/pkg/common/persistence"

	"github.com/google/uuid"
)

type VirtualAccountEntity struct {
	Id string `db:"id"`

	AccountName string `db:"account_name"`

	AccountNumber string `db:"account_number"`

	BankName string `db:"bank_name"`

	Provider string `db:"provider"`

	Reference string `db:"reference"`

	OwnerId string `db:"owner_id"`

	persistence.AuditInfo
}

func NewVirtualAccountEntity() VirtualAccountEntity {

	return VirtualAccountEntity{
		Id:            uuid.NewString(),
		AccountName:   "",
		AccountNumber: "",
		BankName:      "",
		Provider:      "",
		Reference:     "",
		OwnerId:       "",
		AuditInfo:     persistence.AuditInfo{},
	}

}
