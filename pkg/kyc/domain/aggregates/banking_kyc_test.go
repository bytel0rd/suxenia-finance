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

	ownerId, ownerOk := bank.GetOwnerId()

	assert.True(t, ownerOk)
	assert.Equal(t, *ownerId, "ownerId")

	err := bank.SetOwnerId(utils.StrToPr(""))

	assert.Error(t, err)

}
func TestBankingKYCName(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	name, nameOk := bank.GetName()
	assert.True(t, nameOk)
	assert.Equal(t, name, testName)

	err := bank.SetName(utils.StrToPr(""))

	assert.Error(t, err)

}

func TestBankingKYCBankAccountName(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	acctName, ok := bank.GetBankAccountName()
	assert.False(t, ok)
	assert.Equal(t, acctName, nil)

	err := bank.SetBankAccountName(utils.StrToPr(""))

	assert.Error(t, err)

	err = bank.SetBankAccountName(&testName)

	assert.Nil(t, err)

	acctName, ok = bank.GetBankAccountName()
	assert.True(t, ok)
	assert.Equal(t, acctName, testName)
}

func TestBankingKYCBankAccountNumber(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	accountNo, ok := bank.GetBankAccountNumber()
	assert.False(t, ok)
	assert.Equal(t, accountNo, nil)

	err := bank.SetBankAccountNumber(utils.StrToPr(""))

	assert.Error(t, err)

	err = bank.SetBankAccountNumber(utils.StrToPr("0125397736"))

	assert.Nil(t, err)

	accountNo, ok = bank.GetBankAccountNumber()
	assert.True(t, ok)
	assert.Equal(t, accountNo, testName)
}

func TestBankingKYCBankCode(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	code, ok := bank.GetBankCode()
	assert.False(t, ok)
	assert.Equal(t, code, nil)

	err := bank.SetBankCode(utils.StrToPr(""))

	assert.Error(t, err)

	err = bank.SetBankCode(utils.StrToPr("GTB"))

	assert.Nil(t, err)

	code, ok = bank.GetBankCode()
	assert.True(t, ok)
	assert.Equal(t, code, "GTB")

}

func TestBankingKYCBVN(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	bvn, ok := bank.GetBVN()
	assert.False(t, ok)
	assert.Equal(t, bvn, nil)

	err := bank.SetBVN(utils.StrToPr(""))

	assert.Error(t, err)

	err = bank.SetBVN(utils.StrToPr("22222222222"))

	assert.Nil(t, err)

	bvn, ok = bank.GetBankCode()
	assert.True(t, ok)
	assert.Equal(t, bvn, "22222222222")

}

func TestBankingKYCIsVerified(t *testing.T) {

	testName := "Tayo Adekunle"

	bank := NewBankingKYC("ownerId", testName)

	ok := bank.IsVerified()
	assert.False(t, ok)

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

	_ = bank.SetBankCode(utils.StrToPr("GTB"))

	_ = bank.SetBVN(utils.StrToPr("11111111111"))

	_ = bank.SetBankAccountNumber(utils.StrToPr("08149464288"))

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
