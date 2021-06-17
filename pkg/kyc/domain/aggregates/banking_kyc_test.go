package aggregates

import (
	"errors"
	"suxenia-finance/pkg/common/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankingKYCOwnerId(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	ownerId, error := bank.GetOwnerId()

	assert.Nil(t, error)
	assert.Equal(t, *ownerId, "ownerId")

	bank.SetOwnerId(utils.StrToPr("A"))

}
func TestBankingKYCName(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	name, error := bank.GetName()
	assert.Nil(t, error)
	assert.IsType(t, name, new(string))

}

func TestBankingKYCBankAccountName(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	acctName, error := bank.GetBankAccountName()
	assert.Error(t, error)
	assert.Nil(t, acctName)

	bank.SetBankAccountName(&testName)
	acctName, error = bank.GetBankAccountName()

	assert.Nil(t, error)
	assert.IsType(t, acctName, new(string))
}

func TestBankingKYCBankAccountNumber(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	_, error := bank.GetBankAccountNumber()
	assert.Error(t, error)

	bank.SetBankAccountNumber(utils.StrToPr("0125397373"))
	accountNo, error := bank.GetBankAccountNumber()

	assert.Nil(t, error)
	assert.IsType(t, accountNo, new(string))
}

func TestBankingKYCBankCode(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	_, error := bank.GetBankCode()
	assert.Error(t, error)

	bank.SetBankCode(utils.StrToPr("GTB"))
	code, error := bank.GetBankCode()

	assert.Nil(t, error)
	assert.IsType(t, code, new(string))

}

func TestBankingKYCBVN(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	_, error := bank.GetBVN()
	assert.Error(t, error)

	bank.SetBVN(utils.StrToPr("22222222222"))
	bvn, error := bank.GetBVN()

	assert.Nil(t, error)
	assert.IsType(t, bvn, new(string))

}

func TestBankingKYCIsVerified(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	ok := bank.IsVerified()
	assert.False(t, ok)

	bank.SetBankCode(utils.StrToPr("GTB"))

	bank.SetBankAccountNumber(utils.StrToPr(testName))

	bank.SetBVN(utils.StrToPr("11111111111"))

	bank.SetBankAccountNumber(utils.StrToPr("0125397733"))

	err := bank.SetVerificationStatus(true)

	assert.Nil(t, err)

	ok = bank.IsVerified()
	assert.True(t, ok)

}

func TestBankingKYCBVNValidation(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	var failVerification BVNValidator = func(account *BankingKYC, bvn string) (*string, bool, error) {
		return nil, false, errors.New("Error during account validation")
	}

	bvn := "11111111111"

	err := bank.VerifyAndSetBVN(utils.StrToPr(bvn), failVerification)

	assert.Error(t, err)

	var successVerification BVNValidator = func(account *BankingKYC, bvn string) (*string, bool, error) {
		return &testName, true, nil
	}

	err = bank.VerifyAndSetBVN(utils.StrToPr(bvn), successVerification)

	updatedBvn, _ := bank.GetBVN()

	assert.Equal(t, *updatedBvn, bvn)
	assert.Nil(t, err)

}

func TestBankingKYCAccountValidation(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	bank.SetBankCode(utils.StrToPr("GTB"))

	bank.SetBVN(utils.StrToPr("11111111111"))

	bank.SetBankAccountNumber(utils.StrToPr("0125397733"))

	var failVerification AccountValidator = func(account *BankingKYC) (*string, bool, error) {
		return nil, false, errors.New("Error during account validation")
	}

	err := bank.VerifyKYCAcct(failVerification)

	assert.Error(t, err)

	var successVerification AccountValidator = func(account *BankingKYC) (*string, bool, error) {
		return &testName, true, nil
	}

	err = bank.VerifyKYCAcct(successVerification)

	assert.Nil(t, err)

}
