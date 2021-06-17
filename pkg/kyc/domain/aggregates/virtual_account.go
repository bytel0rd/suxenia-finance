package aggregates

import (
	"errors"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/enums"

	"github.com/google/uuid"
)

type VirtualAccount struct {
	id uuid.UUID

	accountName *string

	accountNumber *string

	bankName *string

	provider enums.VirtualAccountProvider

	reference string

	ownerId string

	objects.AuditData
}

func (p *VirtualAccount) GetId() uuid.UUID {
	return p.id
}

func (p *VirtualAccount) SetId(uuid uuid.UUID) error {
	p.id = uuid
	return nil
}

func (p *VirtualAccount) GetReference() string {
	return p.reference
}

func (p *VirtualAccount) SetReference(reference string) error {
	p.reference = reference
	return nil
}

func (p *VirtualAccount) GetOwnerId() string {
	return p.ownerId
}

func (p *VirtualAccount) SetOwnerId(id string) error {

	if utils.IsValidString(id) {
		p.ownerId = id
		return nil
	}

	return errors.New("invalid owner id provided")
}

func (p *VirtualAccount) GetProvider() enums.VirtualAccountProvider {
	return p.provider
}

func (p *VirtualAccount) SetProvider(provider enums.VirtualAccountProvider) {

	p.provider = provider
}

func (p *VirtualAccount) GetAccountName() (*string, bool) {

	if p.accountName != nil && utils.IsValidStringPointer(p.accountName) {
		return p.accountName, true
	}

	return nil, false
}

func (p *VirtualAccount) SetAccountName(acctName *string) error {

	if utils.IsValidStringPointer(acctName) {
		p.accountName = acctName
		return nil
	}

	return errors.New("invalid account name provided for virtual account")
}

func (p *VirtualAccount) GetAccountNumber() (*string, bool) {

	if p.accountNumber != nil && utils.IsValidStringPointer(p.accountNumber) {
		return p.accountNumber, true
	}

	return nil, false
}

func (p *VirtualAccount) SetAccountNumber(number *string) error {

	if utils.IsValidStringPointer(number) && len(*number) == 10 {
		p.accountNumber = number
		return nil
	}

	return errors.New("invalid account number provided for virtual account")
}

type NewVirtualAccountRequest struct {
	createdBy string
	ownerId   string
	provider  enums.VirtualAccountProvider
}

func NewVirtualAccount(acctReq NewVirtualAccountRequest) VirtualAccount {

	audit := objects.AuditData{}
	audit.SetCreatedBy(acctReq.createdBy)
	audit.SetUpdatedBy(acctReq.createdBy)

	return VirtualAccount{
		id:            uuid.New(),
		accountName:   nil,
		accountNumber: nil,
		bankName:      nil,
		provider:      acctReq.provider,
		reference:     enums.GenerateVirtualAccountReference(acctReq.provider),
		ownerId:       acctReq.ownerId,
		AuditData:     audit,
	}

}
