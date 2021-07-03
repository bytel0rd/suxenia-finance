package repos

import (
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/wallet/domain/aggregates"
)

type IWalletRepo interface {
	RetrieveById(id string) (*aggregates.WalletAggregate, bool, *structs.DBException)

	Create(wallet aggregates.WalletAggregate) (*aggregates.WalletAggregate, *structs.DBException)

	Update(wallet aggregates.WalletAggregate) (*aggregates.WalletAggregate, *structs.DBException)

	Delete(id string) (bool, *structs.DBException)
}
