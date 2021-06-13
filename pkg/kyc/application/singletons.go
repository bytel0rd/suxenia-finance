package application

import (
	"suxenia-finance/pkg/kyc/infrastructure/persistence/repos"

	"github.com/jmoiron/sqlx"
)

var bankingKycInstance *repos.BankKycRepo = nil

func InstancateRepos(db *sqlx.DB) error {

	var error error = nil

	bankingKycInstance, error = repos.NewBankycRepo(db)

	if error != nil {
		return error
	}

	return nil
}
