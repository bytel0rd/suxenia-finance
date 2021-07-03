package main

import (
	"suxenia-finance/pkg/common/infrastructure/cache"
	"suxenia-finance/pkg/common/utils"

	kycApplication "suxenia-finance/pkg/kyc/application"
	walletApplication "suxenia-finance/pkg/wallet/application"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")

	if err != nil {
		utils.LoggerInstance.Fatal(err)
	}

	cache := cache.NewRedisCache()

	error := kycApplication.Instancate(db)

	if err != nil {
		utils.LoggerInstance.Fatal(error)
	}

	error = walletApplication.Instancate(db, &cache)

	if err != nil {
		utils.LoggerInstance.Fatal(error)
	}

	r := gin.Default()

	mountHttpInfrastructure(r)

	utils.LoggerInstance.Infof("Server listen on port %s", ":5050")
	r.Run(":5005")

}
