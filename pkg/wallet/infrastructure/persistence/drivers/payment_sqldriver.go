package drivers

import (
	"errors"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PaymentDriver struct {
	db *sqlx.DB
}

func NewPaymentDriver(db *sqlx.DB) (*PaymentDriver, error) {

	if db == nil {
		return nil, errors.New("empty db instance while creating payment db driver")
	}

	driver := PaymentDriver{db: db}

	return &driver, nil
}

func (w *PaymentDriver) FindById(paymentId string) (*entities.Payment, *structs.DBException) {

	payment := entities.Payment{}

	error := w.db.Get(&payment, `Select * from payments where id = $1`, paymentId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return nil, &exception

	}

	return &payment, nil

}

func (w *PaymentDriver) Create(payment entities.Payment) (*entities.Payment, *structs.DBException) {

	result := entities.Payment{}

	rows, error := w.db.NamedQuery(
		`INSERT INTO payments 
			(id, processed_by, amount, opening_balance, owner_id, transaction_reference, transaction_source, source_reference, platform, status, comments, created_by, created_at, updated_by, updated_at) 
		VALUES 
			(:id, :processed_by, :amount, :opening_balance, :owner_id, :transaction_reference, :transaction_source, :source_reference, :platform, :status, :comments, :created_by, :created_at, :updated_by, :updated_at) 
		RETURNING *`, payment)

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

func (w *PaymentDriver) Update(payment entities.Payment) (*entities.Payment, *structs.DBException) {

	result := entities.Payment{}

	rows, error := w.db.NamedQuery(
		`UPDATE
			payments
		SET 
			processed_by = :processed_by, 
			amount = :amount, 
			opening_balance = :opening_balance, 
			transaction_reference = :transaction_reference, 
			transaction_source = :transaction_source, 
			source_reference = :source_reference, 
			platform = :platform, 
			status = :status, 
			owner_id = :owner_id, 
			comments = :comments, 
	        updated_by = :updated_by, 
			updated_at = :updated_at
		WHERE 
			id = :id
		RETURNING *`, payment)

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

func (w *PaymentDriver) Delete(paymentId string) (bool, *structs.DBException) {

	_, error := w.db.Exec(
		`
		DELETE FROM
			payments
		WHERE
			id = $1
	`,
		paymentId)

	if error != nil {

		exception := structs.NewDBException(error, false)

		return false, &exception

	}

	return true, nil

}

func (w *PaymentDriver) FindByReference(reference string) (*entities.Payment, *structs.DBException) {

	payment := entities.Payment{}

	error := w.db.Get(&payment, `Select * from payments where transaction_reference = $1`, reference)

	if error != nil {

		if err, ok := error.(*pq.Error); ok {
			utils.LoggerInstance.Error(err)

			exception := structs.NewDBException(err, true)
			return nil, &exception
		}

		utils.LoggerInstance.Error(error)

		return nil, nil

	}

	return &payment, nil

}
