package routes

import (
	"net/http"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/kyc/application"
	"suxenia-finance/pkg/kyc/mappers"

	"github.com/gin-gonic/gin"
)

func GetBankingKycById(c *gin.Context) {

	bankKyc, exception := application.GetBankingKycById(c.Param("id"))

	if exception != nil {

		c.JSON(int(exception.GetStatusCode()), *exception)

		return
	}

	viewModel, error := mappers.BankingKycAggregateToViewModel(*bankKyc)

	if error != nil {

		exception := structs.NewAPIException(error, http.StatusInternalServerError)

		c.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	response := structs.NewAPIResponse(*viewModel, http.StatusOK)

	c.JSON(int(response.GetStatusCode()), response)

}
