package application

import (
	"errors"
	"net/http"
	commonAggregates "suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/enums"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	kycAggregates "suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BankRepoMock struct {
	aggregate kycAggregates.BankingKYC
	status    bool
	fatal     bool
}

func NewBankRepoMock(operationStatus bool) BankRepoMock {
	return BankRepoMock{
		status:    operationStatus,
		aggregate: kycAggregates.NewBankingKYC("random-owner", "Tayo Adekunle"),
	}
}

func (r *BankRepoMock) RetrieveById(id string) (*kycAggregates.BankingKYC, bool, *structs.DBException) {

	if r.status && r.fatal {
		return nil, false, nil
	}

	if r.status {
		r.aggregate.SetId(&id)
		return &r.aggregate, true, nil
	}

	exception := structs.NewDBException(errors.New("Not Found"), true)

	return nil, false, &exception
}

func (r *BankRepoMock) Create(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, *structs.DBException) {

	if r.status {
		return &bankingKyc, nil
	}

	exception := structs.NewDBException(errors.New("Failed To Create"), true)

	return nil, &exception
}

func (r *BankRepoMock) Update(bankingKyc kycAggregates.BankingKYC) (*kycAggregates.BankingKYC, *structs.DBException) {

	if r.status {
		return &bankingKyc, nil
	}

	exception := structs.NewDBException(errors.New("failed To Update"), false)

	return nil, &exception
}

func (r *BankRepoMock) Delete(id string) (bool, *structs.DBException) {

	if r.status {
		return true, nil
	}

	exception := structs.NewDBException(errors.New("Failed To Delete"), false)

	return false, &exception
}

func NewMockAuthorizedProfile() commonAggregates.AuthorizeProfile {

	profile := commonAggregates.NewAuthorizedProfile()

	profile.SetRole(enums.SUPER_ADMIN)

	return profile
}

func TestGetBankingKycByIdFatal(t *testing.T) {

	repo := NewBankRepoMock(false)
	bankApp, _ := NewBankingKycApplication(&repo)

	bankKyc, error := bankApp.GetBankingKycById("random-id")

	assert.Nil(t, bankKyc)
	assert.Equal(t, error.GetStatusCode(), http.StatusInternalServerError)

}

func TestGetBankingKycByIdNotFound(t *testing.T) {

	repo := NewBankRepoMock(true)
	bankApp, _ := NewBankingKycApplication(&repo)

	repo.fatal = true
	bankKyc, error := bankApp.GetBankingKycById("random-id")

	assert.Nil(t, bankKyc)
	assert.Equal(t, error.GetStatusCode(), http.StatusNotFound)

}

func TestGetBankingKycByIdSuccess(t *testing.T) {

	repo := NewBankRepoMock(true)
	bankApp, _ := NewBankingKycApplication(&repo)

	bankKyc, error := bankApp.GetBankingKycById("random-id")

	assert.Nil(t, error)
	assert.IsType(t, bankKyc, new(kycAggregates.BankingKYC))

}

func TestCreateNewBankingKycBadRequest(t *testing.T) {

	repo := NewBankRepoMock(false)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	createRequest := dtos.CreateBankKycDTO{
		Name:              "",
		BankAccountName:   nil,
		BankAccountNumber: nil,
		BVN:               nil,
		BankCode:          nil,
		OwnerId:           nil,
	}

	bankKyc, error := bankApp.CreateNewBankingKyc(authProfile, createRequest)

	assert.Nil(t, bankKyc)
	assert.Equal(t, error.GetPtr().StatusCode, http.StatusBadRequest)

}

func TestCreateNewBankingKycDBBadRequest(t *testing.T) {

	repo := NewBankRepoMock(false)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	createRequest := dtos.CreateBankKycDTO{
		Name:              "Tayo Adekunle",
		BankAccountName:   nil,
		BankAccountNumber: nil,
		BVN:               nil,
		BankCode:          nil,
		OwnerId:           utils.StrToPr("random-owner-id"),
	}

	bankKyc, error := bankApp.CreateNewBankingKyc(authProfile, createRequest)

	assert.Nil(t, bankKyc)
	assert.Equal(t, error.GetPtr().StatusCode, http.StatusInternalServerError)

}

func TestCreateNewBankingKySuccess(t *testing.T) {

	repo := NewBankRepoMock(true)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	createRequest := dtos.CreateBankKycDTO{
		Name:              "Tayo Adekunle",
		BankAccountName:   nil,
		BankAccountNumber: nil,
		BVN:               nil,
		BankCode:          nil,
		OwnerId:           utils.StrToPr("random-owner-id"),
	}

	bankKyc, error := bankApp.CreateNewBankingKyc(authProfile, createRequest)

	assert.Nil(t, error)
	assert.IsType(t, bankKyc, new(kycAggregates.BankingKYC))

}

func TestUpdateBankingKycValidation(t *testing.T) {
	repo := NewBankRepoMock(true)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	updateRequest := dtos.UpdateBankKycDTO{
		Id:                "",
		Name:              utils.StrToPr("Tayo Adekunle"),
		BankAccountName:   utils.StrToPr("Tayo Adekunle T."),
		BankAccountNumber: nil,
		BVN:               nil,
		BankCode:          nil,
	}

	bankKyc, error := bankApp.UpdateBankingKyc(authProfile, updateRequest)

	assert.Nil(t, bankKyc)
	assert.Equal(t, error.GetStatusCode(), http.StatusBadRequest)

}

func TestUpdateBankingKycDBError(t *testing.T) {
	repo := NewBankRepoMock(false)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	updateRequest := dtos.UpdateBankKycDTO{
		Id:                "Random-Id",
		Name:              utils.StrToPr("Tayo Adekunle"),
		BankAccountName:   utils.StrToPr("Tayo Adekunle T."),
		BankAccountNumber: nil,
		BVN:               nil,
		BankCode:          nil,
	}

	bankKyc, error := bankApp.UpdateBankingKyc(authProfile, updateRequest)

	assert.Nil(t, bankKyc)
	assert.Equal(t, error.GetStatusCode(), http.StatusInternalServerError)

}

func TestUpdateBankingKycSuccess(t *testing.T) {
	repo := NewBankRepoMock(true)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	updateRequest := dtos.UpdateBankKycDTO{
		Id:                "Random-Id",
		Name:              utils.StrToPr("Tayo Adekunle"),
		BankAccountName:   utils.StrToPr("Tayo Adekunle T."),
		BankAccountNumber: nil,
		BVN:               nil,
		BankCode:          nil,
	}

	bankKyc, error := bankApp.UpdateBankingKyc(authProfile, updateRequest)
	name, _ := bankKyc.GetName()

	assert.Nil(t, error)
	assert.Equal(t, name, updateRequest.Name)

}

func TestDeleteBankingKycByIdFailed(t *testing.T) {
	repo := NewBankRepoMock(false)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	deleteRequest := dtos.DeleteBankKycDTO{
		Id: "Random-Id",
	}

	ok, error := bankApp.DeleteBankingKycById(authProfile, deleteRequest)

	assert.False(t, ok)
	assert.IsType(t, error, new(structs.APIException))

}

func TestDeleteBankingKycByIdSuccess(t *testing.T) {
	repo := NewBankRepoMock(true)
	bankApp, _ := NewBankingKycApplication(&repo)

	authProfile := NewMockAuthorizedProfile()
	deleteRequest := dtos.DeleteBankKycDTO{
		Id: "Random-Id",
	}

	ok, error := bankApp.DeleteBankingKycById(authProfile, deleteRequest)

	assert.True(t, ok)
	assert.Nil(t, error)

}
