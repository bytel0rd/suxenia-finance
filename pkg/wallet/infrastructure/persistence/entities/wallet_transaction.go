package entities

import "time"

type WalletTransaction struct {
	Id string

	TransactionType string

	TransactionReference string

	Source string

	Amount int64

	OpeningBalance int64

	Platform string

	OwnerId string

	CreatedBy string

	UpdatedBy string

	CreatedAt time.Time

	UpdateAt time.Time
}
