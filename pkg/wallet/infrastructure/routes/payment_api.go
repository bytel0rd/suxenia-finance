package routes

import (
	"errors"
	"net/http"
	"suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/wallet/application"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/mappers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentApi struct {
	payment *application.PaymentApplication
	logger  *zap.SugaredLogger
}

func NewPaymentApi(
	payment *application.PaymentApplication,
	logger *zap.SugaredLogger,
) (*PaymentApi, error) {

	if payment == nil {
		return nil, errors.New("missing payment appication instance is required to create paymentAPI")
	}

	if logger == nil {
		return nil, errors.New("missing logger instance is required to create paymentAPI")
	}

	api := PaymentApi{
		payment: payment,
		logger:  logger,
	}

	return &api, nil

}

func (api *PaymentApi) GetPaymentById(ctx *gin.Context) {

	payment, exception := api.payment.RetrivePaymentById(ctx.Param("id"))

	if exception != nil {

		ctx.JSON(exception.GetStatusCode(), exception)

		return
	}

	paymentView := mappers.PaymentEntityToView(*payment)

	response := structs.NewAPIResponse(paymentView, http.StatusOK)

	ctx.JSON(response.GetStatusCode(), response)

}

func (api *PaymentApi) InitializePayment(ctx *gin.Context) {

	initializeRequest := dtos.IntitalizePaymentRequest{}

	authProfile := ctx.MustGet("user").(aggregates.AuthorizeProfile)

	error := ctx.ShouldBindJSON(&initializeRequest)

	if error != nil {

		api.logger.Error(error)

		exception := structs.NewBadRequestException(nil)

		ctx.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	if initializeRequest.OwnerId == nil {

		userId, _ := authProfile.GetProfileId()

		initializeRequest.OwnerId = userId

	}

	payment, exception := api.payment.IntitalizePayment(authProfile, initializeRequest)

	if exception != nil {

		ctx.JSON(exception.GetStatusCode(), exception)

		return
	}

	response := structs.NewAPIResponse(payment, http.StatusOK)

	ctx.JSON(response.GetStatusCode(), response)

}

func (api *PaymentApi) ConfirmPayment(ctx *gin.Context) {

	confirmPaymentRequest := dtos.ConfirmPayment{}

	authProfile := ctx.MustGet("user").(aggregates.AuthorizeProfile)

	error := ctx.ShouldBindJSON(&confirmPaymentRequest)

	if error != nil {

		api.logger.Error(error)

		exception := structs.NewBadRequestException(nil)

		ctx.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	if confirmPaymentRequest.OwnerId == nil {

		userId, _ := authProfile.GetProfileId()

		confirmPaymentRequest.OwnerId = userId

	}

	payment, exception := api.payment.ConfirmPayment(authProfile, confirmPaymentRequest)

	if exception != nil {

		ctx.JSON(exception.GetStatusCode(), exception)

		return
	}

	response := structs.NewAPIResponse(payment, http.StatusOK)

	ctx.JSON(response.GetStatusCode(), response)

}
