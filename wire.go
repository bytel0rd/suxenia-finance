//+build wireinject
// The build tag makes sure the stub is not built in the final build.
package main

import (
	kycApp "suxenia-finance/pkg/kyc/application"
	kycRoute "suxenia-finance/pkg/kyc/infrastructure/routes"

	walletApp "suxenia-finance/pkg/wallet/application"
	walletRoute "suxenia-finance/pkg/wallet/infrastructure/routes"

	"github.com/google/wire"
)

func InitalizeApplication() (*Application, error) {
	wire.Build(
		kycApp.BuildSet,
		kycRoute.NewKycRoute,
		walletApp.BuildSet,
		walletRoute.NewPaymentApi,
		NewApplication,
		NewDBInstance,
	)
	return &Application{}, nil
}
