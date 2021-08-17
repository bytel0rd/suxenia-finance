package application

import (
	"errors"
	"fmt"
	"net/http"
	commonAggregates "suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/wallet/domain/aggregates"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/repos"

	"go.uber.org/zap"
)

type WalletApplication struct {
	walletRepo *repos.WalletRepo
	logger     *zap.SugaredLogger
}

func NewWalletApplication(walletRepo *repos.WalletRepo, logger *zap.SugaredLogger) (*WalletApplication, error) {

	if walletRepo == nil {
		return nil, errors.New("missing walletRepo instance required for wallet application")
	}

	if logger == nil {
		return nil, errors.New("missing logger instance required for wallet application")
	}

	app := WalletApplication{walletRepo, logger}

	return &app, nil
}

func (w *WalletApplication) GetWalletById(
	id string,
) (*aggregates.WalletAggregate, *structs.APIException) {

	aggregate, ok, error := w.walletRepo.RetrieveById(id)

	if error == nil && !ok {

		exception := structs.NewAPIExceptionFromString(
			fmt.Sprintf("wallet with id: %s not found", id),
			http.StatusNotFound,
		)

		return nil, &exception
	}

	if error != nil {

		exception := structs.NewAPIException(error, http.StatusInternalServerError)

		return nil, &exception
	}

	return aggregate, nil

}

func (w *WalletApplication) CreateNewWallet(
	profile commonAggregates.AuthorizeProfile,
	createRequest dtos.CreateWalletDTO,
) (*aggregates.WalletAggregate, *structs.APIException) {

	if ok, error := utils.Validate(createRequest); !ok {

		exception := structs.NewRawAPIException(error, "Bad request exception", http.StatusBadRequest)

		return nil, &exception
	}

	Wallet := aggregates.NewWalletAggeregate(createRequest.OwnerId)

	savedWalletEntity, error := w.walletRepo.Create(Wallet)

	if error != nil {

		w.logger.Error(error)

		if error.IsExecError {
			exception := structs.NewInternalServerException(errors.New("error occurred while creating wallet, please try again later."))
			return nil, &exception
		}

		exception := structs.NewBadRequestException(error)

		return nil, &exception
	}

	return savedWalletEntity, nil
}

func (w *WalletApplication) DeleteWalletById(
	profile commonAggregates.AuthorizeProfile, deleteRequest dtos.DeleteWalletDTO,
) (bool, *structs.APIException) {

	if ok, error := utils.Validate(deleteRequest); !ok {

		exception := structs.NewRawAPIException(error, "Bad Request Exception", http.StatusBadRequest)

		return false, &exception
	}

	_, exception := w.GetWalletById(deleteRequest.Id)

	if exception != nil {
		return false, exception
	}

	if !profile.GetRole().IsSuperAdmin() {

		exception := structs.NewUnAuthorizedException(nil)

		return false, &exception

	}

	status, error := w.walletRepo.Delete(deleteRequest.Id)

	if error != nil {

		w.logger.Error(error)

		exception := structs.NewBadRequestException(error)

		return false, &exception
	}

	return status, nil

}
