package entities

import (
	"database/sql"
	"errors"
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/common/utils"

	"github.com/google/uuid"
)

type BankingKycEntity struct {
	Id string `validate:"required,uuid"`

	Name string `validate:"required"`

	BankAccountName sql.NullString `db:"bank_account_name"`

	BankAccountNumber sql.NullString `db:"bank_account_number"`

	BVN sql.NullString

	BankCode sql.NullString `db:"bank_code"`

	OwnerId string `db:"owner_id" validate:"required"`

	Verified bool `validate:"omitempty"`

	persistence.AuditInfo
}

func (kyc *BankingKycEntity) Validate() (bool, error) {

	if status, validationErrors := utils.Validate(kyc); !status {

		validations := *validationErrors

		return false, errors.New(validations[0].Message)

	}

	return true, nil

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
