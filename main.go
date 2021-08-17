package main

import (

	// kycApplication "suxenia-finance/pkg/kyc/application"
	// walletApplication "suxenia-finance/pkg/wallet/application"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	app, error := InitalizeApplication()

	mountHttpInfrastructure(r, app)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Server listen on port %s", ":5050")

	r.Run(":5005")

}
