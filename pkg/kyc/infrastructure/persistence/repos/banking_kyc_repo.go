package repos

import (
	"suxenia-finance/pkg/common/structs"
	kycAggregates "suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/kyc/mappers"

	"github.com/jmoiron/sqlx"
)

func NewBankycRepo(db *sqlx.DB) (*BankingKycRepo, error) {
	driver, error := drivers.NewBankycDriver(db)

	if error != nil {
		return nil, error
	}

	return &BankingKycRepo{
		driver: driver,
	}, nil
}

type BankingKycRepo struct {
	driver *drivers.BankKycDriver
}

func (r *BankingKycRepo) RetrieveById(id string) (*kycAggregates.BankingKYC, bool, *structs.DBException) {

	kycEntity, error := r.driver.FindById(id)

	if error == nil && kycEntity == nil {
		return nil, false, nil
	}

	if error != nil {
		return nil, false, error
	}

	kycAggregate := mappers.BankingKycAggregateFromPersistence(*kycEntity)

	return &kycAggregate, true, nil

}

func (r *BankingKycRepo) Create(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, *structs.DBException) {

	bankingKycEntity := mappers.BankingKycAggregateToPersistence(bankingKyc)

	savedbankingKycEntity, error := r.driver.Create(bankingKycEntity)

	if error != nil {
		return nil, error
	}

	bankingKyc = mappers.BankingKycAggregateFromPersistence(*savedbankingKycEntity)

	return &bankingKyc, nil
}

func (r *BankingKycRepo) Update(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, *structs.DBException) {

	bankingKycEntity := mappers.BankingKycAggregateToPersistence(bankingKyc)

	updatedKycEntity, error := r.driver.Update(bankingKycEntity)

	if error != nil {

		return nil, error
	}

	bankingKyc = mappers.BankingKycAggregateFromPersistence(*updatedKycEntity)

	return &bankingKyc, nil
}

func (r *BankingKycRepo) Delete(id string) (bool, *structs.DBException) {

	_, error := r.driver.FindById(id)

	if error != nil {
		return false, error
	}

	status, error := r.driver.Delete(id)

	if error != nil {

		return false, error
	}

	return status, nil

}
