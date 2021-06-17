package main

import (
	"log"
	"suxenia-finance/pkg/common/utils"
	kycApplication "suxenia-finance/pkg/kyc/application"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	error := kycApplication.Instancate(db)

	if err != nil {
		log.Fatalln(error)
	}

	r := gin.Default()

	mountHttpInfrastructure(r)

	utils.LoggerInstance.Infof("Server listen on port %s", ":5050")
	r.Run(":5005")

}
