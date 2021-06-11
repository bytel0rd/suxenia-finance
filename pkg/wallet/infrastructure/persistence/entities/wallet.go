package entities

import (
	common "suxenia-finance/pkg/common/persistence"
)

type Wallet struct {
	common.AuditInfo

	Id int

	TotalBalance int64

	AvailableBalance int64

	Version int

	OwnerId string
}
