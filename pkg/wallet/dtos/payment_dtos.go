package dtos

import (
	"suxenia-finance/pkg/common/persistence"
	"suxenia-finance/pkg/externals/payments"
	"suxenia-finance/pkg/wallet/enums"
)

type IntitalizePaymentRequest struct {
	SourceReference string  `json:"sourceReference" validate:"required,len=8"`
	Source          string  `json:"source" validate:"required"`
	Email           *string `json:"email" validate:"omitempty,email"`
	Amount          string  `json:"amount" validate:"required"`
	OwnerId         *string `json:"ownerId" validate:"required"`

	Gateway  string `json:"gateway" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

type InitializedPayment struct {
	// store in redis for 30mins with payment_reference.
	TransactionReference string  `json:"transactionReference"`
	SourceReference      string  `json:"sourceReference"`
	Source               string  `json:"source"`
	Email                string  `json:"email"`
	Amount               int     `json:"amount"`
	AmountInMajor        float32 `json:"amountInMajor"`
	OwnerId              string  `json:"ownerId"`

	Gateway  payments.Processor `json:"gateway"`
	Platform enums.Platform     `json:"platform"`
}

type ConfirmPayment struct {
	OwnerId              *string `json:"ownerId"`
	TransactionReference string  `json:"transactionReference"`
	Gateway              string  `json:"gateway"`
	Platform             *string `json:"platform"`
}

type PaymentViewModel struct {
	Id string `json:"id"`

	ProcessedBy payments.Processor `json:"processedBy"`

	Amount int `json:"amount"`

	FormatedAmount string `json:"formatedAmount"`

	OpeningBalance *int32 `json:"openingBalance"`

	FormatedOpeningBalance string `json:"formatedOpeningBalance"`

	TransactionReference string `json:"transactionReference"`

	TransactionSource string `json:"transactionSource"`

	SourceReference *string `json:"sourceReference"`

	Platform enums.Platform `json:"platform"`

	Status enums.TransactionStatus `json:"status"`

	Comments string `json:"comments"`

	OwnerId string `json:"ownerId"`

	persistence.AuditInfo
}
