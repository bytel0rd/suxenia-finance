package entities

import (
	"database/sql"
	"suxenia-finance/pkg/common/persistence"

	"github.com/google/uuid"
)

type Payment struct {
	Id string `db:"id" validate:"required"`

	ProcessedBy string `db:"processed_by" validate:"required"`

	Amount int `db:"amount" validate:"required"`

	OpeningBalance sql.NullInt32 `db:"opening_balance" validate:"omitempty"`

	TransactionReference string `db:"transaction_reference" validate:"required"`

	TransactionSource string `db:"transaction_source" validate:"omitempty"`

	SourceReference sql.NullString `db:"source_reference" validate:"omitempty"`

	Platform string `db:"platform" validate:"required"`

	Status string `db:"status" validate:"required"`

	Comments string `db:"comments" validate:"required"`

	OwnerId string `db:"owner_id" validate:"required"`

	persistence.AuditInfo
}

func NewPayment(ownerId string, auditor string) Payment {

	return Payment{
		Id:        uuid.NewString(),
		OwnerId:   ownerId,
		AuditInfo: persistence.NewAuditInfo(auditor),
	}
}
