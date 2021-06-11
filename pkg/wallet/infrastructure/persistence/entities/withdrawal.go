package entities

import "time"

type Withdrawal struct {
	Id string

	ProcessedBy string

	Amount int64

	OpeningBalance int64

	TransactionReference string

	TransactionSource string

	SourceReference string

	Platform string

	ApprovedBy string

	Status string

	Comment string

	OwnerId string

	CreatedBy string

	UpdatedBy string

	CreatedAt time.Time

	UpdateAt time.Time
}
