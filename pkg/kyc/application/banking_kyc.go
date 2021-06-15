package application

import (
	"fmt"
	"net/http"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/mappers"
)

func NewBankingKycApplication(bankRepo *repos.BankingKycRepo) (*BankingKYCApplication, error) {

	if bankRepo == nil {
		return nil, errors.New("please provide a valid instance of bankRepo to create instance")
	}

	return &BankingKYCApplication{
		bankRepo,
	}, nil
}

func (b *BankingKYCApplication) GetBankingKycById(id string) (*kycAggregates.BankingKYC, *structs.APIException) {

	kycEntity, error := bankingKycInstance.FindById(id)

	if error == nil && kycEntity == nil {

		exception := structs.NewAPIExceptionFromString(fmt.Sprintf("Kyc with id: %s not found", id), http.StatusNotFound)

		return nil, &exception
	}

	if error != nil {
		exception := structs.NewAPIException(error, http.StatusInternalServerError)

		return nil, &exception
	}

	kycAggregate, error := mappers.BankingKycAggregateFromPersistence(*kycEntity)

	if error != nil {
		exception := structs.NewAPIException(error, http.StatusInternalServerError)
		return nil, &exception
	}

	return kycAggregate, nil

}
