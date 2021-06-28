package drivers

import (
	"errors"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
)

type WalletTransactionDriver struct {
	db *sqlx.DB
}

func NewWalletTransactionDriver(db *sqlx.DB) (*WalletTransactionDriver, error) {

	if db == nil {
		return nil, errors.New("empty db instance while creating wallet transaction db driver")
	}

	driver := WalletTransactionDriver{db: db}

	return &driver, nil
}

func (w *WalletTransactionDriver) FindById(transactionId string) (*entities.WalletTransaction, *structs.DBException) {

	transaction := entities.WalletTransaction{}

	error := w.db.Get(&transaction, `Select * from wallet_transactions where id = $1`, transactionId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return nil, &exception

	}

	return &transaction, nil

}

func (w *WalletTransactionDriver) Create(transaction entities.WalletTransaction) (*entities.WalletTransaction, *structs.DBException) {

	result := entities.WalletTransaction{}

	rows, error := w.db.NamedQuery(
		`INSERT INTO wallet_transactions 
			(id, transaction_type, transaction_reference, source, status, amount, opening_balance, platform, owner_id, comments, created_by, created_at, updated_by, updated_at) 
		VALUES 
			(:id, :transaction_type, :transaction_reference, :source, :status, :amount, :opening_balance, :platform, :owner_id, :comments, :created_by, :created_at, :updated_by, :updated_at) 
		RETURNING *`, transaction)

	if error != nil {

		utils.LoggerInstance.Error(error)
		exception := structs.NewDBException(error, true)

		return nil, &exception
	}

	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			utils.LoggerInstance.Error(err)
		}
	}

	return &result, nil

}

func (w *WalletTransactionDriver) Update(wallet entities.WalletTransaction) (*entities.WalletTransaction, *structs.DBException) {

	result := entities.WalletTransaction{}

	rows, error := w.db.NamedQuery(
		`UPDATE
			wallet_transactions
		SET 
			transaction_type = :transaction_type, 
			transaction_reference = :transaction_reference,
			source = :source, 
			status = :status, 
			amount = :amount, 
			opening_balance = :opening_balance, 
			platform = :platform, 
			owner_id = :owner_id, 
			comments = :comments, 
			created_by = :created_by, 
			created_at = :created_at, 
			updated_by = :updated_by, 
			updated_at = :updated_at 
		WHERE 
			id = :id
		RETURNING *`, wallet)

	if error != nil {

		utils.LoggerInstance.Error(error)
		exception := structs.NewDBException(error, true)

		return nil, &exception
	}

	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			utils.LoggerInstance.Error(err)
		}
	}

	return &result, nil

}

func (w *WalletTransactionDriver) Delete(transactionId string) (bool, *structs.DBException) {

	_, error := w.db.Exec(
		`
		DELETE FROM
			wallet_transactions
		WHERE
			id = $1
	`,
		transactionId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return false, &exception

	}

	return true, nil

}
