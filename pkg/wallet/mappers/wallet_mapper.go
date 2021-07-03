package mappers

import (
	"suxenia-finance/pkg/wallet/domain/aggregates"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/shopspring/decimal"
)

func WalletAggregateFromPersistence(wallet entities.Wallet) aggregates.WalletAggregate {

	aggregate := aggregates.NewWalletAggeregate(wallet.OwnerId)

	aggregate.SetId(wallet.Id)

	aggregate.SetAvailableBalance(decimal.NewFromInt(int64(wallet.AvailableBalance)))
	aggregate.SetTotalBalance(decimal.NewFromInt(int64(wallet.TotalBalance)))
	aggregate.SetVersion(wallet.Version)

	aggregate.SetUpdatedBy(wallet.CreatedBy)
	aggregate.SetUpdatedBy(wallet.UpdatedBy)
	aggregate.SetCreatedAt(wallet.CreatedAt)
	aggregate.SetUpdatedAt(wallet.UpdateAt)

	return aggregate

}

func WalletAggregateToPersistence(aggregate aggregates.WalletAggregate) entities.Wallet {

	wallet := entities.Wallet{}

	wallet.Id = aggregate.GetId()
	wallet.OwnerId = aggregate.GetOwnerId()

	wallet.AvailableBalance = int(aggregate.GetAvailableBalance().BigInt().Int64())
	wallet.TotalBalance = int(aggregate.GetTotalBalance().BigInt().Int64())
	wallet.Version = aggregate.GetVersion()

	wallet.AuditInfo.CreatedBy = aggregate.GetCreatedBy()
	wallet.AuditInfo.UpdatedBy = aggregate.GetUpdatedBy()
	wallet.AuditInfo.CreatedAt = aggregate.GetCreatedAt()
	wallet.AuditInfo.UpdateAt = aggregate.GetUpdatedAt()

	return wallet

}

func WalletAggregateToViewModel(aggregate aggregates.WalletAggregate) dtos.WalletViewModel {

	wallet := dtos.WalletViewModel{}

	wallet.Id = aggregate.GetId()
	wallet.OwnerId = aggregate.GetOwnerId()

	wallet.AvailableBalance = int(aggregate.GetAvailableBalance().BigInt().Int64())
	wallet.TotalBalance = int(aggregate.GetTotalBalance().BigInt().Int64())
	wallet.Version = aggregate.GetVersion()

	wallet.AuditInfo.CreatedBy = aggregate.GetCreatedBy()
	wallet.AuditInfo.UpdatedBy = aggregate.GetUpdatedBy()
	wallet.AuditInfo.CreatedAt = aggregate.GetCreatedAt()
	wallet.AuditInfo.UpdateAt = aggregate.GetUpdatedAt()

	return wallet
}
