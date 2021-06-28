package entities

import (
	"database/sql"
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/externals/payments"
	"suxenia-finance/pkg/wallet/enums"

	"github.com/google/uuid"
)

type Withdrawal struct {
	Id string `db:"id" validate:"required,uuid"`

	ProcessedBy payments.Processor `db:"processed_by" validate:"required"`

	Amount int `db:"amount" validate:"required"`

	OpeningBalance int `db:"opening_balance" validate:"required"`

	TransactionReference string `db:"transaction_reference" validate:"required"`

	TransactionSource string `db:"transaction_source" validate:"omitempty"`

	SourceReference string `db:"source_reference" validate:"omitempty"`

	Platform enums.Platform `db:"platform" validate:"required"`

	ApprovedBy sql.NullString `db:"approved_by" validate:"omitempty"`

	Status enums.TransactionStatus `db:"status" validate:"required"`

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
