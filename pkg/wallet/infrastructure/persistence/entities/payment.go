package entities

import (
	"database/sql"
	"time"
)

type Payment struct {
	Id string

	ProcessedBy string

	Amount int64

	OpeningBalance int64

	TransactionReference string

	Currency string

	TransactionSource sql.NullString

	SourceReference sql.NullString

	Platform string

	Status string

	Comment string

	OwnerId string

	CreatedBy string

	UpdatedBy string

	CreatedAt time.Time

	UpdateAt time.Time
}
