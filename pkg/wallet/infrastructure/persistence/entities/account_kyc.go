package entities

import "time"

type AccountKyc struct {
	Id string

	Name string

	BankAccountName string

	BankCode string

	OwnerName string

	BVN string

	OwnerId string

	CreatedBy string

	UpdatedBy string

	CreatedAt time.Time

	UpdateAt time.Time
}
