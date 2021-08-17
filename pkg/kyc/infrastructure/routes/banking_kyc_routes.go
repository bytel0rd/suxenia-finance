package routes

import (
	"errors"
	"net/http"
	"suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/application"
	"suxenia-finance/pkg/kyc/dtos"
	"suxenia-finance/pkg/kyc/mappers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KycRoutes struct {
	kyc    *application.BankingKYCApplication
	logger *zap.SugaredLogger
}

func NewKycRoute(kyc *application.BankingKYCApplication, logger *zap.SugaredLogger) (*KycRoutes, error) {

	if kyc == nil {
		return nil, errors.New("error intializing kyc route, missing application.BankingKYCApplication provider")
	}

	routes := KycRoutes{kyc, logger}

	return &routes, nil

}

func (r *KycRoutes) GetBankingKycById(c *gin.Context) {

	bankKyc, exception := r.kyc.GetBankingKycById(c.Param("id"))

	if exception != nil {

		c.JSON(int(exception.GetStatusCode()), *exception)

		return
	}

	viewModel := mappers.BankingKycAggregateToViewModel(*bankKyc)

	response := structs.NewAPIResponse(viewModel, http.StatusOK)

	c.JSON(int(response.GetStatusCode()), response)

}

func (r *KycRoutes) CreateBankingKyc(c *gin.Context) {

	createRequest := dtos.CreateBankKycDTO{}

	authProfile := c.MustGet("user").(aggregates.AuthorizeProfile)

	error := c.ShouldBindJSON(&createRequest)

	if error != nil {

		r.logger.Error(error)

		exception := structs.NewBadRequestException(nil)

		c.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	if !authProfile.GetRole().IsAdmin() || !utils.IsValidStringPointer(createRequest.OwnerId) {

		if ownerId, ok := authProfile.GetProfileId(); ok {

			createRequest.OwnerId = ownerId

		} else {

			exception := structs.NewUnAuthorizedException(nil)

			c.JSON(int(exception.GetStatusCode()), exception)

			return

		}

	}

	bankKyc, exception := r.kyc.CreateNewBankingKyc(authProfile, createRequest)

	if exception != nil {

		c.JSON(int(exception.GetStatusCode()), *exception)

		return
	}

	viewModel := mappers.BankingKycAggregateToViewModel(*bankKyc)

	response := structs.NewAPIResponse(viewModel, http.StatusOK)

	c.JSON(int(response.GetStatusCode()), response)

}

func (r *KycRoutes) UpdateBankingKyc(c *gin.Context) {

	updateRequest := dtos.UpdateBankKycDTO{}

	authProfile := c.MustGet("user").(aggregates.AuthorizeProfile)

	error := c.ShouldBindJSON(&updateRequest)

	if error != nil {

		r.logger.Error(error)

		exception := structs.NewBadRequestException(nil)

		c.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	bankKyc, exception := r.kyc.UpdateBankingKyc(authProfile, updateRequest)

	if exception != nil {

		c.JSON(int(exception.GetStatusCode()), *exception)

		return
	}

	viewModel := mappers.BankingKycAggregateToViewModel(*bankKyc)

	response := structs.NewAPIResponse(viewModel, http.StatusOK)

	c.JSON(int(response.GetStatusCode()), response)

}

func (r *KycRoutes) DeleteBankingKyc(c *gin.Context) {

	deleteRequest := dtos.DeleteBankKycDTO{}

	authProfile := c.MustGet("user").(aggregates.AuthorizeProfile)

	error := c.ShouldBindJSON(&deleteRequest)

	if error != nil {

		r.logger.Error(error)

		exception := structs.NewBadRequestException(nil)

		c.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	status, exception := r.kyc.DeleteBankingKycById(authProfile, deleteRequest)

	if exception != nil {

		c.JSON(int(exception.GetStatusCode()), *exception)

		return
	}

	response := structs.NewAPIResponse(status, http.StatusOK)

	c.JSON(int(response.GetStatusCode()), response)

}
