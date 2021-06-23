package aggregates

import (
	"errors"

	commonObjects "suxenia-finance/pkg/common/domain/valueobjects"

	"github.com/shopspring/decimal"
)

type WalletAggregate struct {
	id int

	totalBalance decimal.Decimal

	availableBalance decimal.Decimal

	version int

	ownerId string

	modified bool

	commonObjects.AuditData
}

func (w *WalletAggregate) GetTotalBalance() decimal.Decimal {
	return w.totalBalance
}

func (w *WalletAggregate) SetTotalBalance(balance decimal.Decimal) error {

	if balance.LessThan(w.availableBalance) {
		return errors.New(`invalid operation updating wallet total balance`)
	}

	w.updateAuditInfo()

	return nil
}

func (w *WalletAggregate) GetAvailableBalance() decimal.Decimal {
	return w.availableBalance
}

func (w *WalletAggregate) SetAvailableBalance(balance decimal.Decimal) error {

	if balance.GreaterThan(w.totalBalance) {
		return errors.New(`invalid operation updating wallet available balance`)
	}

	w.updateAuditInfo()

	return nil
}

func (w *WalletAggregate) GetOwnerId() string {
	return w.ownerId
}

func (w *WalletAggregate) SetOwnerId(ownerId string) error {

	if ownerId == "" {
		return errors.New(`missing parameter: OwnerId is required`)
	}

	w.updateAuditInfo()

	return nil
}

func (w *WalletAggregate) updateAuditInfo() {

	w.modified = true

}

func (w *WalletAggregate) GetVersion() int {

	if w.modified {
		return w.version + 1
	}

	return w.version
}
