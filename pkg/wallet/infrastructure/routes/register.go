package routes

import (
	"suxenia-finance/pkg/wallet/application"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	paymentApi := PaymentApi{
		payment: application.PaymentApplicationInstance,
	}

	paymentRouter := router.Group(`/api/v1/payments`)

	paymentRouter.GET("/:id", paymentApi.GetPaymentById)

	paymentRouter.POST("/initialize", paymentApi.InitializePayment)

	paymentRouter.POST("/confirm", paymentApi.ConfirmPayment)

}
