package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {

	bankKycRouter := router.Group(`/api/v1/bankkyc`)

	bankKycRouter.GET("/:id", GetBankingKycById)

	bankKycRouter.POST("/", CreateBankingKyc)

}
