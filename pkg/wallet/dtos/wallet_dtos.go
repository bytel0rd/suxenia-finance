package dtos

import "suxenia-finance/pkg/common/persistence"

type WalletViewModel struct {
	persistence.AuditInfo

	Id string `json:"id"`

	TotalBalance int `json:"totalBalance"`

	AvailableBalance int `json:"availableBalance"`

	Version int `json:"version"`

	OwnerId string `json:"ownerId"`
}

type CreateWalletDTO struct {
	OwnerId string `json:"ownerId" validate:"required"`
}

type DeleteWalletDTO struct {
	Id string `json:"id" validate:"required"`
}
