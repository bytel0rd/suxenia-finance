package enums

import "errors"

type TransactionType string

func newTransactionType(name string) TransactionType {
	return TransactionType(name)
}

var (
	PAYMENT    TransactionType = newTransactionType("PAYMENT")
	WITHDRAWAL TransactionType = newTransactionType("WITHDRAWAL")
)

func (t *TransactionType) Name() string {
	return string(*t)
}

func ParseTransactionType(name string) (*TransactionType, error) {
	switch name {

	case PAYMENT.Name():

		return &PAYMENT, nil

	case WITHDRAWAL.Name():

		return &WITHDRAWAL, nil

	}

	return nil, errors.New("invalid transaction type selected")

}
