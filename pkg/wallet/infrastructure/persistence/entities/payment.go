package entities

import (
	"database/sql"
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/externals/payments"
	"suxenia-finance/pkg/wallet/enums"

	"github.com/google/uuid"
)

type Payment struct {
	Id string `db:"id" validate:"required"`

	ProcessedBy payments.Processor `db:"processed_by" validate:"required"`

	Amount int `db:"amount" validate:"required"`

	OpeningBalance sql.NullInt32 `db:"opening_balance" validate:"omitempty"`

	TransactionReference string `db:"transaction_reference" validate:"required"`

	TransactionSource string `db:"transaction_source" validate:"omitempty"`

	SourceReference sql.NullString `db:"source_reference" validate:"omitempty"`

	Platform enums.Platform `db:"platform" validate:"required"`

	Status enums.TransactionStatus `db:"status" validate:"required"`

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
