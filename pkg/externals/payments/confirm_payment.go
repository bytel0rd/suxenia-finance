package payments

import "github.com/shopspring/decimal"

type VerifiedPayment struct {
	ProcessedBy Processor

	Amount decimal.Decimal

	TransactionReference string
}

type ConfirmPayment func(processor Processor, reference string) (*VerifiedPayment, error)
