package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func RegisterRoutes(router *gin.Engine, routes *KycRoutes) {

	bankKycRouter := router.Group(`/api/v1/bankkyc`)

	bankKycRouter.GET("/:id", routes.GetBankingKycById)

	bankKycRouter.POST("/", routes.CreateBankingKyc)
	bankKycRouter.PUT("/", routes.UpdateBankingKyc)
	bankKycRouter.DELETE("/", routes.DeleteBankingKyc)

}

var BuildSet wire.ProviderSet = wire.NewSet(NewKycRoute)
