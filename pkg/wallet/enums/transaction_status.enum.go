package enums

import "errors"

type TransactionStatus string

func newTransactionStatus(name string) TransactionStatus {
	return TransactionStatus(name)
}

var (
	PENDING    TransactionStatus = newTransactionStatus("PENDING")
	SUCCESS    TransactionStatus = newTransactionStatus("SUCCESS")
	FAILED     TransactionStatus = newTransactionStatus("FAILED")
	REJECTED   TransactionStatus = newTransactionStatus("REJECTED")
	PROCESSING TransactionStatus = newTransactionStatus("PROCESSING")
	INITIATED  TransactionStatus = newTransactionStatus("INITIATED")
)

func (t *TransactionStatus) Name() string {
	return string(*t)
}

func ParseTransactionStatus(name string) (*TransactionStatus, error) {
	switch name {

	case PENDING.Name():

		return &PENDING, nil

	case SUCCESS.Name():

		return &SUCCESS, nil

	case FAILED.Name():

		return &FAILED, nil

	case REJECTED.Name():

		return &REJECTED, nil

	case PROCESSING.Name():

		return &PROCESSING, nil

	case INITIATED.Name():

		return &INITIATED, nil

	}

	return nil, errors.New("invalid transaction status")

}
