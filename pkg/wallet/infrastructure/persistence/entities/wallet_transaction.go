package entities

import (
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/wallet/enums"

	"github.com/google/uuid"
)

type WalletTransaction struct {
	Id string `db:"id" validate:"required,uuid"`

	TransactionType string `db:"transaction_type" validate:"required"`

	TransactionReference string `db:"transaction_reference" validate:"required"`

	Source string `db:"source" validate:"required"`

	Status enums.TransactionStatus `db:"status" validate:"required"`

	Amount int `db:"amount" validate:"required"`

	OpeningBalance int `db:"opening_balance" validate:"required"`

	Platform enums.Platform `db:"platform" validate:"required"`

	OwnerId string `db:"owner_id" validate:"required"`

	Comments string `db:"comments" validate:"required"`

	persistence.AuditInfo
}

func NewWalletTransaction(owner_id string, auditor string) WalletTransaction {

	audit := persistence.NewAuditInfo(auditor)

	return WalletTransaction{
		Id:                   uuid.NewString(),
		TransactionType:      "",
		TransactionReference: "",
		Source:               "",
		Amount:               0,
		OpeningBalance:       0,
		Platform:             "",
		OwnerId:              owner_id,
		Comments:             "",
		AuditInfo:            audit,
	}

}
