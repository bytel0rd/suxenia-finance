package main

import (
	"net/http"
	"strings"
	"suxenia-finance/pkg/common/mappers"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"

	kycRoutes "suxenia-finance/pkg/kyc/infrastructure/routes"

	"gopkg.in/dgrijalva/jwt-go.v3"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func mountHttpInfrastructure(r *gin.Engine) {

	r.Use(authJWTMiddleWare)

	kycRoutes.RegisterRoutes(r)
}

func authJWTMiddleWare(r *gin.Context) {

	rawToken := r.Request.Header.Get("authorization")

	tokenSplit := strings.Split(rawToken, " ")

	if len(tokenSplit) != 2 || strings.ToLower(tokenSplit[0]) != "bearer" {

		exception := structs.NewAPIExceptionFromString("Invalid Acccess Token provided", http.StatusUnauthorized)

		r.AbortWithStatusJSON(exception.GetStatusCode(), exception)

		return
	}

	token := tokenSplit[1]

	claim := jwt.MapClaims{}

	parsedToken, _ := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return "Zuxenia", nil
	})

	if parsedToken == nil {

		exception := structs.NewUnAuthorizedException(nil)

		r.AbortWithStatusJSON(exception.GetStatusCode(), exception)

		return
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	error := claims.Valid()

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewUnAuthorizedException(nil)

		r.AbortWithStatusJSON(exception.GetStatusCode(), exception)

		return
	}

	profile := structs.AuthProfile{}

	error = mapstructure.Decode(claims, &profile)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewUnAuthorizedException(nil)

		r.AbortWithStatusJSON(exception.GetStatusCode(), exception)

		return
	}

	authorizedProfile, exception := mappers.NewAuthorizedProfileFromAuthProfile(profile)

	if exception != nil {

		utils.LoggerInstance.Error(exception)

		r.AbortWithStatusJSON(exception.GetStatusCode(), exception)

		return
	}

	r.Set("user", *authorizedProfile)

	r.Next()

}
