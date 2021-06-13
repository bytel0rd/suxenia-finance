package repos

import (
	"errors"
	"fmt"
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

	"github.com/jmoiron/sqlx"
)

func NewVirtualAccountRepo(db *sqlx.DB) (*VirtualAccountRepo, error) {

	if db == nil {
		return nil, errors.New("cannot create virtual account repo due to invalid connection provided")
	}

	return &VirtualAccountRepo{db}, nil

}

type VirtualAccountRepo struct {
	db *sqlx.DB
}

func (b *VirtualAccountRepo) Create(kyc entities.VirtualAccountEntity) (*entities.VirtualAccountEntity, error) {

	_, err := b.db.NamedExec(

		`INSERT INTO virtual_accounts (
			id, bank_name, account_name, account_number, provider, reference, owner_id,  created_by, updated_by, created_at, updated_at
		)
		VALUES (
			:id, :bank_name, :account_name, :account_number, :provider, :reference, :owner_id, :created_by, :updated_by, :created_at, :updated_at
		)`, kyc)

	if err != nil {
		return nil, err
	}

	return &kyc, nil
}

func (b VirtualAccountRepo) FindById(id string) (*entities.VirtualAccountEntity, error) {

	kyc := entities.VirtualAccountEntity{}

	err := b.db.Get(&kyc, "SELECT * FROM virtual_accounts WHERE id = $1", id)

	if err != nil {

		message := err.Error()
		utils.LoggerInstance.Warn(
			message,
		)

		return nil, nil
	}

	return &kyc, nil
}

func (b *VirtualAccountRepo) Update(kyc *entities.VirtualAccountEntity) (*entities.VirtualAccountEntity, error) {

	result := entities.VirtualAccountEntity{
		Id:            "",
		AccountName:   "",
		AccountNumber: "",
		BankName:      "",
		Provider:      "",
		Reference:     "",
		OwnerId:       "",
		AuditInfo:     persistence.AuditInfo{},
	}

	rows, err := b.db.NamedQuery(
		`UPDATE virtual_accounts SET
			bank_name = :bank_name,
			account_name = :account_name,
			account_number = :account_number,
			provider = :provider,
			reference = :reference,
			updated_by =  :updated_by,
			updated_at = :updated_at
		WHERE
			id = :id
		RETURNING
			id, bank_name, account_name, account_number, bank_name, provider, reference, owner_id, created_by, updated_by, created_at, updated_at
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

func (b *VirtualAccountRepo) Delete(id string) (bool, error) {

	_, err := b.db.Exec("delete from virtual_accounts where id = $1", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
