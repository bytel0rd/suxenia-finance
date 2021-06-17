package application

import (
	"errors"
	"fmt"
	"net/http"
	commonAggregates "suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	kycAggregates "suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/dtos"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/repos"
)

type BankingKYCApplication struct {
	bankRepo repos.IBankingKycRepo
}

func NewBankingKycApplication(bankRepo repos.IBankingKycRepo) (*BankingKYCApplication, error) {

	if bankRepo == nil {
		return nil, errors.New("please provide a valid instance of bankRepo to create instance")
	}

	return &BankingKYCApplication{
		bankRepo,
	}, nil
}

func (b *BankingKYCApplication) GetBankingKycById(id string) (*kycAggregates.BankingKYC, *structs.APIException) {

	kycAggregate, ok, error := b.bankRepo.RetrieveById(id)

	if error == nil && !ok {

		exception := structs.NewAPIExceptionFromString(fmt.Sprintf("Kyc with id: %s not found", id), http.StatusNotFound)

		return nil, &exception
	}

	if error != nil {

		exception := structs.NewAPIException(error, http.StatusInternalServerError)

		return nil, &exception
	}

	return kycAggregate, nil

}

func (b *BankingKYCApplication) CreateNewBankingKyc(profile commonAggregates.AuthorizeProfile, createRequest dtos.CreateBankKycDTO) (*kycAggregates.BankingKYC, *structs.APIException) {

	if ok, error := utils.Validate(createRequest); !ok {

		exception := structs.NewRawAPIException(error, "Bad request exception", http.StatusBadRequest)

		return nil, &exception
	}

	bankingKyc := kycAggregates.CreateBankingKYC(createRequest)

	savedbankingKycEntity, error := b.bankRepo.Create(bankingKyc)

	if error != nil {

		if error.IsExecError {
			exception := structs.NewInternalServerException(error)
			return nil, &exception
		}

		exception := structs.NewInternalServerException(error)
		return nil, &exception
	}

	return savedbankingKycEntity, nil
}

func (b *BankingKYCApplication) UpdateBankingKyc(profile commonAggregates.AuthorizeProfile, updateRequest dtos.UpdateBankKycDTO) (*kycAggregates.BankingKYC, *structs.APIException) {

	if ok, error := utils.Validate(updateRequest); !ok {

		exception := structs.NewRawAPIException(error, "Bad Request Exception", http.StatusBadRequest)

		return nil, &exception
	}

	bankingKyc, exception := b.GetBankingKycById(updateRequest.Id)

	if exception != nil {
		return nil, exception
	}

	if !profile.GetRole().IsAdmin() {

		ownerId, ok := profile.GetProfileId()

		bankingKycOwnerId, _ := bankingKyc.GetOwnerId()

		if ok && ownerId != bankingKycOwnerId {

			exception := structs.NewUnAuthorizedException(nil)

			return nil, &exception
		}

	}

	error := bankingKyc.ApplyUpdate(profile, updateRequest)

	if error != nil {

		exception := structs.NewBadRequestException(error)

		return nil, &exception
	}

	updatedKycEntity, updateException := b.bankRepo.Update(*bankingKyc)

	if updateException != nil {

		exception := structs.NewAPIException(updateException, http.StatusInternalServerError)

		return nil, &exception
	}

	return updatedKycEntity, nil
}

func (b *BankingKYCApplication) DeleteBankingKycById(profile commonAggregates.AuthorizeProfile, deleteRequest dtos.DeleteBankKycDTO) (bool, *structs.APIException) {

	if ok, error := utils.Validate(deleteRequest); !ok {

		exception := structs.NewRawAPIException(error, "Bad Request Exception", http.StatusBadRequest)

		return false, &exception
	}

	bankingKyc, exception := b.GetBankingKycById(deleteRequest.Id)

	if exception != nil {
		return false, exception
	}

	if !profile.GetRole().IsAdmin() {

		ownerId, ok := profile.GetProfileId()

		bankingKycOwnerId, _ := bankingKyc.GetOwnerId()

		if ok && ownerId != bankingKycOwnerId {

			exception := structs.NewUnAuthorizedException(nil)

			return false, &exception
		}

	}

	status, error := b.bankRepo.Delete(deleteRequest.Id)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewBadRequestException(error)

		return false, &exception
	}

	return status, nil

}
