package routes

import (
	"net/http"
	"suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/wallet/application"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/mappers"

	"github.com/gin-gonic/gin"
)

type PaymentApi struct {
	payment *application.PaymentApplication
}

func (self *PaymentApi) GetPaymentById(ctx *gin.Context) {

	payment, exception := self.payment.RetrivePaymentById(ctx.Param("id"))

	if exception != nil {

		ctx.JSON(exception.GetStatusCode(), exception)

		return
	}

	paymentView := mappers.PaymentEntityToView(*payment)

	response := structs.NewAPIResponse(paymentView, http.StatusOK)

	ctx.JSON(response.GetStatusCode(), response)

}

func (self *PaymentApi) InitializePayment(ctx *gin.Context) {

	initializeRequest := dtos.IntitalizePaymentRequest{}

	authProfile := ctx.MustGet("user").(aggregates.AuthorizeProfile)

	error := ctx.ShouldBindJSON(&initializeRequest)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewBadRequestException(nil)

		ctx.JSON(int(exception.GetStatusCode()), exception)

		return
	}

	payment, exception := self.payment.IntitalizePayment(authProfile, initializeRequest)

	if exception != nil {

		ctx.JSON(exception.GetStatusCode(), exception)

		return
	}

	response := structs.NewAPIResponse(payment, http.StatusOK)

	ctx.JSON(response.GetStatusCode(), response)

}
