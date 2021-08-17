package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, paymentApi *PaymentApi) {

	paymentRouter := router.Group(`/api/v1/payments`)

	paymentRouter.GET("/:id", paymentApi.GetPaymentById)

	paymentRouter.POST("/initialize", paymentApi.InitializePayment)

	paymentRouter.POST("/confirm", paymentApi.ConfirmPayment)

}
