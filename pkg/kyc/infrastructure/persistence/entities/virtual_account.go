package entities

import (
	"suxenia-finance/pkg/common/persistence"

	"github.com/google/uuid"
)

type VirtualAccountEntity struct {
	Id string

	AccountName string

	AccountNumber string

	BankName string

	Provider string

	Reference string

	OwnerId string

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
