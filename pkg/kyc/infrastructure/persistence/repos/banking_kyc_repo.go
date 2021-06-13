package repos

import (
	"errors"
	"fmt"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
)

type BankKycRepo struct {
	db *sqlx.DB
}

func (b *BankKycRepo) Create(kyc entities.BankingKycEntity) (*entities.BankingKycEntity, error) {

	// how to autogenrate via tags or reflection capacblites.

	_, err := b.db.NamedExec(

		`INSERT INTO banking_kyc (
			id, name, bank_account_name, bank_account_number, bvn, bank_code, owner_id, verified, created_by, updated_by, created_at, updated_at
		)
		VALUES (
			:id, :name, :bank_account_name, :bank_account_number, :bvn, :bank_code, :owner_id, :verified, :created_by, :updated_by, :created_at, :updated_at
		)`, kyc)

	if err != nil {
		return nil, err
	}

	return &kyc, nil
}

func (b *BankKycRepo) FindById(id string) (*entities.BankingKycEntity, error) {

	kyc := entities.BankingKycEntity{}

	err := b.db.Get(&kyc, "SELECT * FROM banking_kyc WHERE id = $1", id)

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return &kyc, nil
}

func (b *BankKycRepo) Update(kyc *entities.BankingKycEntity) (*entities.BankingKycEntity, error) {

	_, err := b.db.NamedExec(
		`UPDATE banking_kyc SET
			id = :id 
			name = :name 
			bank_account_name = :bank_account_name  
			bank_account_number = :bank_account_number
			bvn = :bvn 
			bank_code = :bank_code 
			owner_id = :owner_id 
			verified = :verified 
			created_by = :created_by 
			updated_by =  :updated_by 
			created_at = :created_at
			updated_at = :updated_at
		WHERE
			id = :id
		`, kyc)

	if err != nil {
		return nil, err
	}

	return kyc, nil

}

func (b *BankKycRepo) Delete(id string) (bool, error) {

	_, err := b.db.Exec("delete from banking_kyc where id = $1", id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func NewBankycRepo(db *sqlx.DB) (*BankKycRepo, error) {

	if db == nil {
		return nil, errors.New("cannot create banking repo due to invalid connection provided")
	}

	return &BankKycRepo{db}, nil

}
