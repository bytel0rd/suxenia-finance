package application

import (
	"reflect"
	"suxenia-finance/pkg/common/infrastructure/cache"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/externals/payments"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/repos"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var PaymentApplicationInstance *PaymentApplication = nil

var verifyPayment payments.ConfirmPayment = func(processor payments.Processor, reference string) (*payments.VerifiedPayment, error) {

	verified := payments.VerifiedPayment{
		ProcessedBy:          processor,
		Amount:               decimal.NewFromInt(1000),
		TransactionReference: reference,
	}

	return &verified, nil
}

func Instancate(db *sqlx.DB, cache cache.Cache) error {

	walletDriver, error := drivers.NewWalletDriver(db)

	if error != nil {

		utils.LoggerInstance.Fatalf(
			error.Error(),
			zap.String("Instance", reflect.TypeOf(walletDriver).String()),
		)

		return error
	}

	paymentDriver, error := drivers.NewPaymentDriver(db)

	if error != nil {

		utils.LoggerInstance.Fatalf(
			error.Error(),
			zap.String("Instance", reflect.TypeOf(paymentDriver).String()),
		)

		return error
	}

	walletTransactionDriver, error := drivers.NewWalletTransactionDriver(db)

	if error != nil {

		utils.LoggerInstance.Fatalf(
			error.Error(),
			zap.String("Instance", reflect.TypeOf(walletTransactionDriver).String()),
		)

		return error
	}

	walletRepo, error := repos.NewWalletRepo(walletDriver)

	if error != nil {

		utils.LoggerInstance.Fatalf(
			error.Error(),
			zap.String("Instance", reflect.TypeOf(walletRepo).String()),
		)

		return error
	}

	paymentApplication := PaymentApplication{
		paymentDriver:           paymentDriver,
		verifyPayment:           verifyPayment,
		cache:                   cache,
		walletRepo:              walletRepo,
		walletTransactionDriver: walletTransactionDriver,
	}

	PaymentApplicationInstance = &paymentApplication

	utils.LoggerInstance.Infof("Successfully created %s Instance", reflect.TypeOf(paymentApplication).String())

	return nil
}
