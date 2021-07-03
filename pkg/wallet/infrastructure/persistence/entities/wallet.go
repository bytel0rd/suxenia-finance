package entities

import (
	"suxenia-finance/pkg/common/persistence"

	"github.com/google/uuid"
)

type Wallet struct {
	persistence.AuditInfo

	Id string `validate:"required,uuid" db:"id"`

	TotalBalance int `db:"total_balance" validate:"required"`

	AvailableBalance int `db:"available_balance" validate:"required"`

	Version int `db:"version" validate:"required"`

	OwnerId string `db:"owner_id" validate:"required"`
}

func NewWallet(ownerId string, authorName string) Wallet {

	return Wallet{
		Id:        uuid.NewString(),
		OwnerId:   ownerId,
		AuditInfo: persistence.NewAuditInfo(authorName),
	}

}
