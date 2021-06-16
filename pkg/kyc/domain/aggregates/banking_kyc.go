package aggregates

import (
	"errors"
	"suxenia-finance/pkg/common/domain/aggregates"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/dtos"

	"github.com/google/uuid"
)

func NewBankingKYC(ownerId string, name string) BankingKYC {

	auditData := objects.AuditData{}

	auditData.SetCreatedBy(name)
	auditData.SetUpdatedBy(name)

	kyc := BankingKYC{
		id:                utils.StrToPr(uuid.New().String()),
		name:              nil,
		bankAccountName:   nil,
		bankAccountNumber: nil,
		bankCode:          nil,
		bvn:               nil,
		ownerId:           &ownerId,
		AuditData:         auditData,
	}

	kyc.SetName(&name)

	return kyc

}

func CreateBankingKYC(request dtos.CreateBankKycDTO) BankingKYC {

	auditData := objects.NewAuditData(request.Name)

	bankingKyc := BankingKYC{
		id:                utils.StrToPr(uuid.New().String()),
		name:              nil,
		bankAccountName:   nil,
		bankAccountNumber: nil,
		bankCode:          nil,
		bvn:               nil,
		ownerId:           nil,
		verified:          false,
		AuditData:         auditData,
	}

	bankingKyc.SetName(&request.Name)
	bankingKyc.SetBankAccountName(request.BankAccountName)
	bankingKyc.SetBankAccountNumber(request.BankAccountNumber)
	bankingKyc.SetBankCode(request.BankCode)
	bankingKyc.SetBVN(request.BVN)
	bankingKyc.SetOwnerId(request.OwnerId)

	return bankingKyc

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

func (p *BankingKYC) GetId() (*string, error) {

	if p.id != nil {
		return p.id, nil
	}

	return nil, errors.New("invalid id provided for banking kyc")
}

func (p *BankingKYC) SetId(id *string) {

	if utils.IsValidStringPointer(id) {

		p.id = id

	}

}

func (p *BankingKYC) GetOwnerId() (*string, error) {

	if p.ownerId != nil {
		return p.ownerId, nil
	}

	return nil, errors.New("invalid id provided for banking kyc")
}

func (p *BankingKYC) SetOwnerId(id *string) {

	if utils.IsValidStringPointer(id) {

		p.ownerId = id

	}

}

func (p *BankingKYC) GetName() (*string, error) {

	if utils.IsValidStringPointer(p.name) {
		return p.name, nil
	}

	return nil, errors.New("missing name provided for banking kyc")
}

func (p *BankingKYC) SetName(name *string) {

	if utils.IsValidStringPointer(name) {

		p.name = name

	}

}

func (p *BankingKYC) GetBankAccountName() (*string, error) {

	if p.bankAccountName != nil && utils.IsValidStringPointer(p.bankAccountName) {
		return p.bankAccountName, nil
	}

	return nil, errors.New("invalid account name provided for banking kyc")
}

func (p *BankingKYC) SetBankAccountName(acctName *string) {

	if utils.IsValidStringPointer(acctName) {

		p.bankAccountName = acctName

	}

}

func (p *BankingKYC) GetBankAccountNumber() (*string, error) {

	if p.bankAccountNumber != nil {
		return p.bankAccountNumber, nil
	}

	return nil, errors.New("invalid account number provided for banking kyc")
}

func (p *BankingKYC) SetBankAccountNumber(number *string) {

	if utils.IsValidStringPointer(number) && len(*number) == 10 {

		p.bankAccountNumber = number

	}

}

func (p *BankingKYC) GetBankCode() (*string, error) {

	if utils.IsValidStringPointer(p.bankCode) {
		return p.bankCode, nil
	}

	return nil, errors.New("invalid bank code provided for banking kyc")
}

func (p *BankingKYC) SetBankCode(code *string) {

	if utils.IsValidStringPointer(code) {

		p.bankCode = code

	}

}

func (p *BankingKYC) GetBVN() (*string, error) {

	if utils.IsValidStringPointer(p.bvn) {

		return p.bvn, nil

	}

	return nil, errors.New("invalid bvn provided for banking kyc")
}

func (p *BankingKYC) SetBVN(bvn *string) {

	if utils.IsValidStringPointer(bvn) {

		p.bvn = bvn

	}

}

func (p *BankingKYC) VerifyAndSetBVN(bvn *string, validator BVNValidator) error {

	if utils.IsValidStringPointer(bvn) {

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

func (p *BankingKYC) ApplyUpdate(authProfile aggregates.AuthorizeProfile, updateRequest dtos.UpdateBankKycDTO) error {

	p.SetName(updateRequest.Name)

	p.SetBankAccountName(updateRequest.BankAccountName)

	p.SetBankAccountNumber(updateRequest.BankAccountNumber)

	p.SetBankCode(updateRequest.BankCode)

	p.SetBVN(updateRequest.BVN)

	p.SetVerificationStatus(false)

	return nil
}

func (p *BankingKYC) AnyFieldError() error {

	_, error := p.GetId()

	if error != nil {
		return error
	}

	_, error = p.GetName()

	if error != nil {
		return error
	}

	_, error = p.GetBankAccountName()

	if error != nil {
		return error
	}

	_, error = p.GetBankAccountNumber()

	if error != nil {
		return error
	}

	_, error = p.GetBVN()

	if error != nil {
		return error
	}

	_, error = p.GetBankCode()

	if error != nil {
		return error
	}

	_, error = p.GetOwnerId()

	if error != nil {
		return error
	}

	return nil

}
