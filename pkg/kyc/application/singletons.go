package application

import (
	"reflect"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/repos"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var bankingKycInstance *repos.BankingKycRepo = nil
var BankingApplicationInstance *BankingKYCApplication = nil

func Instancate(db *sqlx.DB) error {

	var error error = nil

	bankingKycInstance, error = repos.NewBankycRepo(db)

	if error != nil {
		utils.LoggerInstance.Fatalf(error.Error(), zap.String("Instance", reflect.TypeOf(bankingKycInstance).String()))
		return error
	}

	utils.LoggerInstance.Infof("Successfully created %s Instance", reflect.TypeOf(bankingKycInstance).String())

	bankingApplicationInstance, error := NewBankingKycApplication(bankingKycInstance)

	if error != nil {
		utils.LoggerInstance.Fatalf(error.Error(), zap.String("Instance", reflect.TypeOf(bankingApplicationInstance).String()))
		return error
	}

	BankingApplicationInstance = bankingApplicationInstance

	return nil
}
