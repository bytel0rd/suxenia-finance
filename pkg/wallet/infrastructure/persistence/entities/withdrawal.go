package entities

import (
	"database/sql"
	"suxenia-finance/pkg/common/persistence"

	"github.com/google/uuid"
)

type Withdrawal struct {
	Id string `db:"id" validate:"required,uuid"`

	ProcessedBy string `db:"processed_by" validate:"required"`

	Amount int `db:"amount" validate:"required"`

	OpeningBalance int `db:"opening_balance" validate:"required"`

	TransactionReference string `db:"transaction_reference" validate:"required"`

	TransactionSource string `db:"transaction_source" validate:"omitempty"`

	SourceReference string `db:"source_reference" validate:"omitempty"`

	Platform string `db:"platform" validate:"required"`

	ApprovedBy sql.NullString `db:"approved_by" validate:"omitempty"`

	Status string `db:"status" validate:"required"`

	Comments string `db:"comments" validate:"omitempty"`

	OwnerId string `db:"owner_id" validate:"required"`

	persistence.AuditInfo
}

func NewWithdrawal(ownerId string, auditor string) Withdrawal {
	return Withdrawal{
		Id:        uuid.NewString(),
		OwnerId:   ownerId,
		AuditInfo: persistence.NewAuditInfo(auditor),
	}
}
