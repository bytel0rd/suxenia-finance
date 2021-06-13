package aggregates

import (
	"errors"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/utils"

	"github.com/google/uuid"
)

func NewBankingKYC(ownerId string, name string) BankingKYC {

	auditData := objects.AuditData{}

	auditData.SetCreatedBy(name)
	auditData.SetUpdatedBy(name)

	return BankingKYC{
		id:                utils.StrToPr(uuid.New().String()),
		name:              nil,
		bankAccountName:   nil,
		bankAccountNumber: nil,
		bankCode:          nil,
		bvn:               nil,
		ownerId:           &ownerId,
		AuditData:         auditData,
	}

}

type BankingKYC struct {
	id *string

	name *string

	bankAccountName *string

	bankAccountNumber *string

	bankCode *string

	bvn *string

	ownerId *string

	verified bool

	objects.AuditData
}

func (p *BankingKYC) GetId() (*string, bool) {

	if p.id != nil {
		return p.id, true
	}

	return nil, false
}

func (p *BankingKYC) SetId(id *string) error {

	if utils.IsValidString(id) {
		p.id = id
		return nil
	}

	return errors.New("invalid id provided for banking kyc")
}

func (p *BankingKYC) GetOwnerId() (*string, bool) {

	if p.ownerId != nil {
		return p.ownerId, true
	}

	return nil, false
}

func (p *BankingKYC) SetOwnerId(id *string) error {

	if utils.IsValidString(id) {
		p.ownerId = id
		return nil
	}

	return errors.New("invalid id provided for banking kyc")
}

func (p *BankingKYC) GetName() (*string, bool) {

	if p.name != nil && utils.IsValidString(p.name) {
		return p.name, true
	}

	return nil, false
}

func (p *BankingKYC) SetName(name *string) error {

	if utils.IsValidString(name) {
		p.name = name
		return nil
	}

	return errors.New("missing name provided for banking kyc")
}

func (p *BankingKYC) GetBankAccountName() (*string, bool) {

	if p.bankAccountName != nil && utils.IsValidString(p.bankAccountName) {
		return p.bankAccountName, true
	}

	return nil, false
}

func (p *BankingKYC) SetBankAccountName(acctName *string) error {

	if utils.IsValidString(acctName) {
		p.bankAccountName = acctName
		return nil
	}

	return errors.New("invalid account name provided for banking kyc")
}

func (p *BankingKYC) GetBankAccountNumber() (*string, bool) {

	if p.bankAccountName != nil && utils.IsValidString(p.bankAccountName) {
		return p.bankAccountNumber, true
	}

	return nil, false
}

func (p *BankingKYC) SetBankAccountNumber(number *string) error {

	if utils.IsValidString(number) && len(*number) == 10 {
		p.bankAccountNumber = number
		return nil
	}

	return errors.New("invalid account number provided for banking kyc")
}

func (p *BankingKYC) GetBankCode() (*string, bool) {

	if utils.IsValidString(p.bankCode) {
		return p.bankCode, true
	}

	return nil, false
}

func (p *BankingKYC) SetBankCode(code *string) error {

	if utils.IsValidString(code) {
		p.bankCode = code
		return nil
	}

	return errors.New("invalid bank code provided for banking kyc")
}

func (p *BankingKYC) GetBVN() (*string, bool) {

	if utils.IsValidString(p.bvn) {
		return p.bvn, true
	}

	return nil, false
}

func (p *BankingKYC) SetBVN(bvn *string) error {

	if utils.IsValidString(bvn) {
		p.bvn = bvn
		return nil
	}

	return errors.New("invalid bvn provided for banking kyc")
}

func (p *BankingKYC) VerifyAndSetBVN(bvn *string, validator BVNValidator) error {

	if utils.IsValidString(bvn) {

		actName, ok, err := validator(p, *bvn)

		if err != nil && !ok {
			return err
		}

		p.name = actName
		p.bvn = bvn

		return nil
	}

	return errors.New("invalid bvn provided for banking kyc")
}

func (p *BankingKYC) IsVerified() bool {
	return p.verified
}

func (p *BankingKYC) SetVerificationStatus(status bool) error {

	isAccountComplete := p.isAccountDetailComplete()

	if isAccountComplete {

		p.verified = status

		return nil
	}

	if status && !isAccountComplete {
		return errors.New("incomplete account details cannot be verified")
	}

	p.verified = status

	return nil
}

func (p *BankingKYC) isAccountDetailComplete() bool {

	if p.IsVerified() {
		return p.IsVerified()
	}

	return p.bvn != nil && p.bankAccountNumber != nil && p.bankCode != nil
}

func (p *BankingKYC) VerifyKYCAcct(validator AccountValidator) error {

	if p.bvn == nil || p.name == nil {
		return errors.New("please complete bvn verification before account verification")
	}

	if p.bankAccountNumber == nil {
		return errors.New("account number is missing and required before continue account verification")

	}

	if p.bankCode == nil {
		return errors.New("please provide bank code to continue account verification")
	}

	acctName, ok, err := validator(p)

	if err != nil && !ok {
		return err
	}

	p.bankAccountName = acctName
	p.SetVerificationStatus(ok)

	return nil
}
