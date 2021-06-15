package drivers

import (
	"database/sql"
	"errors"
	"fmt"
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type BankKycDriver struct {
	db *sqlx.DB
}

func (b *BankKycDriver) Create(kyc entities.BankingKycEntity) (*entities.BankingKycEntity, error) {

	if valid, error := kyc.Validate(); !valid {
		return nil, error
	}

	_, err := b.db.NamedExec(

		`INSERT INTO banking_kyc (
			id, name, bank_account_name, bank_account_number, bvn, bank_code, owner_id, verified, created_by, updated_by, created_at, updated_at
		)
		VALUES (
			:id, :name, :bank_account_name, :bank_account_number, :bvn, :bank_code, :owner_id, :verified, :created_by, :updated_by, :created_at, :updated_at
		)`, kyc)

	if err, ok := err.(*pq.Error); ok {
		utils.LoggerInstance.Error(err)
	}

	if err != nil {
		utils.LoggerInstance.Error(err)

		a := err.(*pq.Error)

		utils.LoggerInstance.Error(a)
		return nil, err
	}

	return &kyc, nil
}

func (b BankKycDriver) FindById(id string) (*entities.BankingKycEntity, error) {

	kyc := entities.BankingKycEntity{}

	err := b.db.Get(&kyc, "SELECT * FROM banking_kyc WHERE id = $1", id)

	if err != nil {

		message := err.Error()

		utils.LoggerInstance.Warn(
			message,
		)

		return nil, nil
	}

	return &kyc, nil
}

func (b *BankKycDriver) Update(kyc entities.BankingKycEntity) (*entities.BankingKycEntity, error) {

	if valid, error := kyc.Validate(); !valid {
		return nil, error
	}

	result := entities.BankingKycEntity{
		Id:                "",
		Name:              "",
		BankAccountName:   sql.NullString{},
		BankAccountNumber: sql.NullString{},
		BVN:               sql.NullString{},
		BankCode:          sql.NullString{},
		OwnerId:           "",
		Verified:          false,
		AuditInfo:         persistence.AuditInfo{},
	}

	rows, err := b.db.NamedQuery(
		`UPDATE banking_kyc SET
			name = :name,  
			bank_account_name = :bank_account_name,
			bank_account_number = :bank_account_number,
			bvn = :bvn,
			bank_code = :bank_code,
			verified = :verified,
			updated_by =  :updated_by,
			updated_at = :updated_at
		WHERE
			id = :id
		RETURNING
			id, name, bank_account_name, bank_account_number, bvn, bank_code, owner_id, verified, created_by, updated_by, created_at, updated_at
		`, kyc)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			fmt.Println(err)
		}
	}

	utils.LoggerInstance.Info(result)

	return &result, nil

}

func (b *BankKycDriver) Delete(id string) (bool, error) {

	_, err := b.db.Exec("delete from banking_kyc where id = $1", id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func NewBankycDriver(db *sqlx.DB) (*BankKycDriver, error) {

	if db == nil {
		return nil, errors.New("cannot create banking repo due to invalid connection provided")
	}

	return &BankKycDriver{db}, nil

}
