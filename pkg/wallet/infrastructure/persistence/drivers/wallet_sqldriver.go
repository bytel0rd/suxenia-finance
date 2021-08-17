package drivers

import (
	"errors"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type WalletDriver struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewWalletDriver(db *sqlx.DB, logger *zap.SugaredLogger) (*WalletDriver, error) {

	if db == nil {
		return nil, errors.New("empty db instance while creating wallet db driver")
	}

	if logger == nil {
		return nil, errors.New("missing logger instance while creating wallet db driver")
	}

	driver := WalletDriver{db, logger}

	return &driver, nil
}

func (w *WalletDriver) FindWalletById(walletId string) (*entities.Wallet, *structs.DBException) {

	wallet := entities.Wallet{}

	error := w.db.Get(&wallet, `Select * from wallets where id = $1`, walletId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return nil, &exception

	}

	return &wallet, nil

}

func (w *WalletDriver) Create(wallet entities.Wallet) (*entities.Wallet, *structs.DBException) {

	result := entities.Wallet{}

	rows, error := w.db.NamedQuery(
		`INSERT INTO wallets 
			(id, total_balance, available_balance, version, owner_id, created_by, created_at, updated_by, updated_at) 
		VALUES 
			(:id, :total_balance, :available_balance, :version, :owner_id, :created_by, :created_at, :updated_by, :updated_at) 
		RETURNING *`, wallet)

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

func (w *WalletDriver) Update(wallet entities.Wallet) (*entities.Wallet, *structs.DBException) {

	result := entities.Wallet{}

	rows, error := w.db.NamedQuery(
		`UPDATE
			wallets 
		SET 
			 total_balance = :total_balance,
			 available_balance = :available_balance,
			 version =  :version + 1, 
			 owner_id = :owner_id,
			 updated_by = :updated_by, 
			 updated_at =  :updated_at
		WHERE 
			id = :id AND version = :version 
		RETURNING *`, wallet)

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

func (w *WalletDriver) Delete(walletId string) (bool, *structs.DBException) {

	_, error := w.db.Exec(
		`
		DELETE FROM
			wallets
		WHERE
			id = $1
	`,
		walletId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return false, &exception

	}

	return true, nil

}

func (w *WalletDriver) FindByOwnerId(ownerId string) (*entities.Wallet, *structs.DBException) {

	wallet := entities.Wallet{}

	error := w.db.Get(&wallet, `Select * from wallets where owner_id = $1`, ownerId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return nil, &exception

	}

	return &wallet, nil

}
