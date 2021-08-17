package drivers

import (
	"errors"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type WithdrawalDriver struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewWithdrawalDriver(db *sqlx.DB, logger *zap.SugaredLogger) (*WithdrawalDriver, error) {

	if db == nil {
		return nil, errors.New("empty db instance while creating withdrawal db driver")
	}

	if logger == nil {
		return nil, errors.New("missing logger instance while creating withdrawal db driver")
	}

	driver := WithdrawalDriver{db, logger}

	return &driver, nil
}

func (w *WithdrawalDriver) FindById(withdrawalId string) (*entities.Withdrawal, *structs.DBException) {

	withdrawal := entities.Withdrawal{}

	error := w.db.Get(&withdrawal, `Select * from withdrawals where id = $1`, withdrawalId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return nil, &exception

	}

	return &withdrawal, nil

}

func (w *WithdrawalDriver) Create(withdrawal entities.Withdrawal) (*entities.Withdrawal, *structs.DBException) {

	result := entities.Withdrawal{}

	rows, error := w.db.NamedQuery(
		`INSERT INTO withdrawals 
			(id, processed_by, amount, opening_balance, owner_id, transaction_reference, transaction_source, source_reference, platform, approved_by, status, comments, created_by, created_at, updated_by, updated_at) 
		VALUES 
			(:id, :processed_by, :amount, :opening_balance, :owner_id, :transaction_reference, :transaction_source, :source_reference, :platform, :approved_by, :status, :comments, :created_by, :created_at, :updated_by, :updated_at) 
		RETURNING *`, withdrawal)

	if error != nil {

		w.logger.Error(error)
		exception := structs.NewDBException(error, true)

		return nil, &exception
	}

	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			w.logger.Error(err)
		}
	}

	return &result, nil

}

func (w *WithdrawalDriver) Update(withdrawal entities.Withdrawal) (*entities.Withdrawal, *structs.DBException) {

	result := entities.Withdrawal{}

	rows, error := w.db.NamedQuery(
		`UPDATE
			Withdrawals
		SET 
			processed_by = :processed_by, 
			amount = :amount, 
			opening_balance = :opening_balance, 
			transaction_reference = :transaction_reference, 
			transaction_source = :transaction_source, 
			source_reference = :source_reference, 
			platform = :platform, 
			owner_id = :owner_id,
			approved_by = :approved_by, 
			status = :status, 
			comments = :comments, 
	        updated_by = :updated_by, 
			updated_at = :updated_at
		WHERE 
			id = :id
		RETURNING *`, withdrawal)

	if error != nil {

		w.logger.Error(error)
		exception := structs.NewDBException(error, true)

		return nil, &exception
	}

	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			w.logger.Error(err)
		}
	}

	return &result, nil

}

func (w *WithdrawalDriver) Delete(withdrawalId string) (bool, *structs.DBException) {

	_, error := w.db.Exec(
		`
		DELETE FROM
			withdrawals
		WHERE
			id = $1
	`,
		withdrawalId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return false, &exception

	}

	return true, nil

}
