package drivers

import (
	"database/sql"
	"errors"
	"fmt"
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type BankKycDriver struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewBankycDriver(db *sqlx.DB, logger *zap.SugaredLogger) (*BankKycDriver, error) {

	if db == nil {
		return nil, errors.New("cannot create banking repo due to invalid connection provided")
	}

	if logger == nil {
		return nil, errors.New("cannot create banking repo due missing logger instance")

	}

	return &BankKycDriver{db, logger}, nil

}

func (b *BankKycDriver) Create(kyc entities.BankingKycEntity) (*entities.BankingKycEntity, *structs.DBException) {

	if valid, error := kyc.Validate(); !valid {

		exception := structs.NewDBException(error, true)

		return nil, &exception
	}

	_, err := b.db.NamedExec(

		`INSERT INTO banking_kyc (
			id, name, bank_account_name, bank_account_number, bvn, bank_code, owner_id, verified, created_by, updated_by, created_at, updated_at
		)
		VALUES (
			:id, :name, :bank_account_name, :bank_account_number, :bvn, :bank_code, :owner_id, :verified, :created_by, :updated_by, :created_at, :updated_at
		)`, kyc)

	if err, ok := err.(*pq.Error); ok {
		b.logger.Error(err)
	}

	if err != nil {

		pgError := err.(*pq.Error)
		b.logger.Error(pgError)

		exception := structs.NewDBException(pgError, false)
		return nil, &exception
	}

	return &kyc, nil
}

func (b BankKycDriver) FindById(id string) (*entities.BankingKycEntity, *structs.DBException) {

	kyc := entities.BankingKycEntity{}

	err := b.db.Get(&kyc, "SELECT * FROM banking_kyc WHERE id = $1", id)

	if err != nil {

		message := err.Error()

		b.logger.Warn(
			message,
		)

		return nil, nil
	}

	return &kyc, nil
}

func (b *BankKycDriver) Update(kyc entities.BankingKycEntity) (*entities.BankingKycEntity, *structs.DBException) {

	if valid, error := kyc.Validate(); !valid {
		exception := structs.NewDBException(error, true)
		return nil, &exception
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
		exception := structs.NewDBException(err, true)

		return nil, &exception
	}

	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			fmt.Println(err)
		}
	}

	b.logger.Info(result)

	return &result, nil

}

func (b *BankKycDriver) Delete(id string) (bool, *structs.DBException) {

	_, err := b.db.Exec("delete from banking_kyc where id = $1", id)

	if err != nil {
		b.logger.Error(err)

		exception := structs.NewDBException(err, true)
		return false, &exception
	}

	return true, nil
}
