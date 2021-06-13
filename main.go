package main

import (
	"log"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/application"
	"suxenia-finance/pkg/kyc/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// This  project is been implemented with domain driven design (DDD)
// corresponding test are beside the aggregates or value objects
// the code are under the pkg/kyc/domain and pkg/common at the moment
// the reason for that is the project is been modeled such that the business logic is pure and
// independent of other service dependencies.
// The justification for DDD is because is project is aimed towards a reliable financial systems
// from my experience th fastest way to loose trust is inaccurate financial accounting

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	error := application.InstancateRepos(db)

	if err != nil {
		log.Fatalln(error)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	utils.LoggerInstance.Infof("Server listen on port %s", ":5050")
	r.Run(":5005")

}
