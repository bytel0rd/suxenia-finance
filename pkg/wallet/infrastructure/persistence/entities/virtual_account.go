package entities

import "time"

type VirtualAccount struct {
	Id string

	AccountName string

	AccountNumber string

	BankName string

	Provider string

	Reference string

	OwnerId string

	CreatedBy string

	UpdatedBy string

	CreatedAt time.Time

	UpdateAt time.Time
}
