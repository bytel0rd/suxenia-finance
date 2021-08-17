package main

import (
	"errors"
	"log"
	kycRoutes "suxenia-finance/pkg/kyc/infrastructure/routes"
	walletRoutes "suxenia-finance/pkg/wallet/infrastructure/routes"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Application struct {
	kyc     *kycRoutes.KycRoutes
	payment *walletRoutes.PaymentApi
}

func NewApplication(kyc *kycRoutes.KycRoutes, payment *walletRoutes.PaymentApi) (*Application, error) {

	if kyc == nil {
		return nil, errors.New("cannot create application instance due to missing kyc router instance")
	}

	if payment == nil {
		return nil, errors.New("cannot create application instance due to missing payment router instance")
	}

	return &Application{kyc, payment}, nil

}

func NewDBInstance() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
