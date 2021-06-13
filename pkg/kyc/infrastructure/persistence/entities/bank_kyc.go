package entities

import (
	"database/sql"
	"suxenia-finance/pkg/common/persistence"

	"github.com/google/uuid"
)

type BankingKycEntity struct {
	Id string

	Name string

	BankAccountName sql.NullString `db:"bank_account_name"`

	BankAccountNumber sql.NullString `db:"bank_account_number"`

	BVN sql.NullString

	BankCode sql.NullString `db:"bank_code"`

	OwnerId string `db:"owner_id"`

	Verified bool

	persistence.AuditInfo
}

func (kyc *BankingKycEntity) Validate() (error, bool) {

	return nil, true

}

func NewBankingKycEntity(owner_name, owner_id string) BankingKycEntity {
	return BankingKycEntity{
		Id:                uuid.NewString(),
		Name:              owner_name,
		BankAccountName:   sql.NullString{},
		BankAccountNumber: sql.NullString{},
		BankCode:          sql.NullString{},
		BVN:               sql.NullString{},
		OwnerId:           owner_id,
		Verified:          false,
		AuditInfo:         persistence.NewAuditInfo(owner_name),
	}
}
