package repos

import (
	"errors"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/wallet/domain/aggregates"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/wallet/mappers"
)

func NewWalletRepo(walletDriver *drivers.WalletDriver) (*WalletRepo, error) {

	if walletDriver == nil {
		return nil, errors.New("wallet repo cannot be instanciated due to missing wallet driver")
	}

	return &WalletRepo{
		driver: walletDriver,
	}, nil

}

type WalletRepo struct {
	driver *drivers.WalletDriver
}

func (r *WalletRepo) RetrieveById(id string) (*aggregates.WalletAggregate, bool, *structs.DBException) {

	wallet, error := r.driver.FindWalletById(id)

	if error == nil && wallet == nil {
		return nil, false, nil
	}

	if error != nil {
		return nil, false, error
	}

	aggregate := mappers.WalletAggregateFromPersistence(*wallet)

	return &aggregate, true, nil

}

func (r *WalletRepo) Create(wallet aggregates.WalletAggregate) (*aggregates.WalletAggregate, *structs.DBException) {

	entity := mappers.WalletAggregateToPersistence(wallet)

	savedWalletEntity, error := r.driver.Create(entity)

	if error != nil {
		return nil, error
	}

	wallet = mappers.WalletAggregateFromPersistence(*savedWalletEntity)

	return &wallet, nil
}

func (r *WalletRepo) Update(wallet aggregates.WalletAggregate) (*aggregates.WalletAggregate, *structs.DBException) {

	entity := mappers.WalletAggregateToPersistence(wallet)

	updatedWalletEntity, error := r.driver.Update(entity)

	if error != nil {

		return nil, error
	}

	wallet = mappers.WalletAggregateFromPersistence(*updatedWalletEntity)

	return &wallet, nil
}

func (r *WalletRepo) Delete(id string) (bool, *structs.DBException) {

	_, error := r.driver.FindWalletById(id)

	if error != nil {
		return false, error
	}

	status, error := r.driver.Delete(id)

	if error != nil {

		return false, error
	}

	return status, nil

}
