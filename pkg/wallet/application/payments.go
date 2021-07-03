package application

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	commonAggregates "suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/infrastructure/cache"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/externals/payments"
	"suxenia-finance/pkg/wallet/domain/aggregates"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/enums"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/repos"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentApplication struct {
	paymentDriver *drivers.PaymentDriver
	verifyPayment payments.ConfirmPayment
	// mail                    func(email string, payment entities.Payment) (bool, error)
	cache                   cache.Cache
	walletRepo              *repos.WalletRepo
	walletTransactionDriver *drivers.WalletTransactionDriver
}

func (p *PaymentApplication) RetrivePaymentById(id string) (*entities.Payment,
	*structs.APIException) {

	payment, error := p.paymentDriver.FindById(id)

	if error != nil {

		if error.IsExecError {

			exception := structs.NewInternalServerException(error)

			return nil, &exception
		}

		exception := structs.NewBadRequestException(error)

		return nil, &exception

	}

	if payment == nil {
		exception := structs.NewAPIExceptionFromString(
			fmt.Sprintf(" Payment with id: %s not found",
				id),
			http.StatusNotFound)

		return nil, &exception
	}

	return payment, nil

}

func (p *PaymentApplication) IntitalizePayment(
	profile commonAggregates.AuthorizeProfile,
	request dtos.IntitalizePaymentRequest,
) (*dtos.InitializedPayment, *structs.APIException) {

	if ok, error := utils.Validate(request); !ok {

		exception := structs.NewRawAPIException(error, "Bad Request Exception", http.StatusBadRequest)

		return nil, &exception
	}

	intializedPayment, exception := p.createInitializePaymentRequest(profile, request)

	if exception != nil {

		return nil, exception

	}

	exception = p.cacheIntitializePayment(intializedPayment)

	if exception != nil {

		return nil, exception

	}

	return intializedPayment, nil
}

func (p *PaymentApplication) cacheIntitializePayment(intializedPayment *dtos.InitializedPayment) *structs.APIException {

	intializePaymentJSON, error := json.Marshal(intializedPayment)

	internalServerError := errors.New("error occurred intitalizing transaction, please try again later")

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewInternalServerException(internalServerError)

		return &exception
	}

	error = p.cache.Put(intializedPayment.TransactionReference, string(intializePaymentJSON))

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewInternalServerException(internalServerError)

		return &exception
	}

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewInternalServerException(internalServerError)

		return &exception
	}

	return nil
}

func (p *PaymentApplication) createInitializePaymentRequest(
	profile commonAggregates.AuthorizeProfile,
	request dtos.IntitalizePaymentRequest,
) (*dtos.InitializedPayment, *structs.APIException) {

	email, ok := profile.GetEmail().GetAddress()

	if !ok {

		exception := structs.NewUnAuthorizedException(
			errors.New("incomplete authorization profile, please contact admin"),
		)

		return nil, &exception
	}

	processor, error := payments.ParseProcessor(request.Gateway)

	if error != nil {

		exception := structs.NewBadRequestException(error)

		return nil, &exception
	}

	platform, error := enums.ParsePlatform(request.Platform)

	if error != nil {

		exception := structs.NewBadRequestException(error)

		return nil, &exception
	}

	amount, error := decimal.NewFromString(request.Amount)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewBadRequestException(errors.New("invalid amount provided"))

		return nil, &exception
	}

	if amount.LessThan(decimal.NewFromInt(1000)) {

		exception := structs.NewBadRequestException(
			errors.New("mininum amount for a payment transaction is #1000.00"),
		)

		return nil, &exception
	}

	shortCode, _ := payments.ProcessorShortCode(*processor)

	transactionReference := fmt.Sprintf("SZX-%s-PYT-%s-%s", *platform, *shortCode, uuid.NewString()[0:6])

	majorAmount, _ := amount.BigFloat().Float32()

	minorAmount := amount.Mul(decimal.NewFromInt(100)).BigInt().Int64()

	intializedPayment := dtos.InitializedPayment{
		TransactionReference: transactionReference,
		SourceReference:      request.SourceReference,
		Source:               request.Source,
		Email:                *email,
		Amount:               int(minorAmount),
		AmountInMajor:        majorAmount,
		Gateway:              *processor,
		Platform:             *platform,
		OwnerId:              *request.OwnerId,
	}

	return &intializedPayment, nil
}

func (p *PaymentApplication) ConfirmPayment(
	profile commonAggregates.AuthorizeProfile,
	request dtos.ConfirmPayment,
) (*entities.Payment, *structs.APIException) {

	if ok, error := utils.Validate(request); !ok {

		exception := structs.NewRawAPIException(error, "Bad Request Exception", http.StatusBadRequest)

		return nil, &exception
	}

	confirmedPayment, exception := p.getVerifiedPayment(request)

	if exception != nil {

		return nil, exception

	}

	name, ok := profile.GetFullName()

	if !ok {

		exception := structs.NewAPIExceptionFromString(
			"Incomplete authorization profile, please contact admin for any complain.",
			http.StatusUnauthorized,
		)

		return nil, &exception
	}

	_, exception = p.isPaymentAlreadyProcessed(confirmedPayment)

	if exception != nil {

		return nil, exception

	}

	intializedTransaction, exception := p.retrieveCachedIntializePaymentInformation(confirmedPayment)

	if exception != nil {

		return nil, exception

	}

	newPayment := p.intializePaymentRecord(*name, *intializedTransaction, *confirmedPayment)

	processedPayment, exception := p.processPayment(newPayment)

	if exception != nil {

		utils.LoggerInstance.Error(exception)

		return p.savePaymentRecord(newPayment)
	}

	return processedPayment, nil

}

func (p *PaymentApplication) processPayment(payment entities.Payment) (*entities.Payment, *structs.APIException) {

	wallet, ok, dbException := p.walletRepo.RetrieveById(payment.Id)

	if dbException != nil {

		exception := structs.NewAPIExceptionFromString(
			"Error occurred while processing payment on the wallet, please try again later",
			http.StatusInternalServerError,
		)

		return nil, &exception
	}

	if !ok {

		message := fmt.Sprintf("Wallet with Id %s  not found", wallet.GetId())

		exception := structs.NewAPIExceptionFromString(message, http.StatusInternalServerError)

		return nil, &exception

	}

	processedPayment, walletTransaction, exception := wallet.ProcessPayment(payment)

	if exception != nil {
		return nil, exception
	}

	processedPayment, exception = p.persistPaymentTransactionOrRollBack(wallet, processedPayment, walletTransaction)

	if exception != nil {

		return p.savePaymentRecord(payment)

	}

	return processedPayment, nil
}

func (p *PaymentApplication) persistPaymentTransactionOrRollBack(
	wallet *aggregates.WalletAggregate,
	processedPayment *entities.Payment,
	walletTransaction *entities.WalletTransaction,
) (*entities.Payment, *structs.APIException) {

	processedPayment, exception := p.paymentDriver.Create(*processedPayment)

	if exception != nil {

		exception := structs.NewAPIExceptionFromString(
			"Error occurred while processing payment on the wallet, please try again later",
			http.StatusInternalServerError,
		)

		return nil, &exception
	}

	walletTransaction, exception = p.walletTransactionDriver.Create(*walletTransaction)

	if exception != nil {

		status, _ := p.paymentDriver.Delete(processedPayment.Id)

		if !status {
			utils.LoggerInstance.Error(
				"Payment transaction record could not be deleted to revert processing information",
				processedPayment.Id,
			)
		}

		exception := structs.NewAPIExceptionFromString(
			"Error occurred while processing payment on the wallet, please try again later",
			http.StatusInternalServerError,
		)

		return nil, &exception
	}

	_, exception = p.walletRepo.Update(*wallet)

	if exception != nil {

		status, _ := p.paymentDriver.Delete(processedPayment.Id)

		if !status {
			utils.LoggerInstance.Error(
				"Payment transaction record could not be deleted to revert processing information",
				processedPayment.Id,
			)
		}

		status, _ = p.walletTransactionDriver.Delete(walletTransaction.Id)

		if !status {
			utils.LoggerInstance.Error(
				"Wallet transaction could not be deleted to revert processing information",
				processedPayment.Id,
			)
		}

		exception := structs.NewAPIExceptionFromString(
			"Error occurred while processing payment on the wallet, please try again later",
			http.StatusInternalServerError,
		)

		return nil, &exception
	}

	return processedPayment, nil
}

func (p *PaymentApplication) savePaymentRecord(newPayment entities.Payment) (*entities.Payment, *structs.APIException) {

	savedPayment, exception := p.paymentDriver.Create(newPayment)

	if exception != nil {

		if exception.IsExecError {

			apiError := structs.NewInternalServerException(
				errors.New("error occurred during payment confirmation, please try again later"),
			)

			return nil, &apiError
		}

		apiError := structs.NewBadRequestException(exception)

		return nil, &apiError
	}

	return savedPayment, nil
}

func (p *PaymentApplication) intializePaymentRecord(
	auditor string,
	intializedTransaction dtos.InitializedPayment,
	confirmedPayment payments.VerifiedPayment,
) entities.Payment {

	newPayment := entities.NewPayment(intializedTransaction.OwnerId, auditor)

	newPayment.Comments = "Payment has been verified but not proccessed by the wallet yet."

	newPayment.Amount = int(confirmedPayment.Amount.BigInt().Int64())

	newPayment.Status = enums.PENDING

	newPayment.TransactionReference = confirmedPayment.TransactionReference

	newPayment.SourceReference = sql.NullString{String: intializedTransaction.SourceReference}

	newPayment.Platform = intializedTransaction.Platform

	newPayment.TransactionSource = intializedTransaction.Source

	newPayment.ProcessedBy = confirmedPayment.ProcessedBy

	return newPayment
}

func (p *PaymentApplication) retrieveCachedIntializePaymentInformation(
	confirmedPayment *payments.VerifiedPayment,
) (*dtos.InitializedPayment, *structs.APIException) {

	intializedTransaction := dtos.InitializedPayment{}

	intializedTransactionJSON, error := p.cache.Retrieve(confirmedPayment.TransactionReference)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewInternalServerException(error)

		return nil, &exception
	}

	if intializedTransactionJSON == nil {

		message := fmt.Sprintf(
			"Payment was completed outside transaction window, please forward %s transaction reference to admins.",
			confirmedPayment.TransactionReference,
		)

		exception := structs.NewAPIExceptionFromString(message, http.StatusPreconditionFailed)

		return nil, &exception
	}

	error = json.Unmarshal([]byte(*intializedTransactionJSON), &intializedTransaction)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewInternalServerException(
			errors.New("error occurred during payment confirmation, please try again later"),
		)

		return nil, &exception
	}

	return &intializedTransaction, nil
}

func (p *PaymentApplication) isPaymentAlreadyProcessed(
	confirmedPayment *payments.VerifiedPayment,
) (bool, *structs.APIException) {

	payment, dbException := p.paymentDriver.FindByReference(
		confirmedPayment.TransactionReference,
	)

	if dbException != nil {

		utils.LoggerInstance.Error(dbException)

		exception := structs.NewInternalServerException(
			errors.New("error occurred during payment confirmation, please try again later"),
		)

		return false, &exception
	}

	if payment != nil {

		exception := structs.NewAPIExceptionFromString(
			"Payment Transaction as already been processed, please contact admin for any complain.",
			http.StatusConflict,
		)

		return true, &exception
	}

	return false, nil
}

func (p *PaymentApplication) getVerifiedPayment(
	request dtos.ConfirmPayment,
) (*payments.VerifiedPayment, *structs.APIException) {

	if request.OwnerId == nil {

		exception := structs.NewAPIExceptionFromString(
			"Transaction OwnerId is missing",
			http.StatusPreconditionRequired,
		)

		return nil, &exception
	}

	if ok, error := utils.Validate(request); !ok {

		exception := structs.NewRawAPIException(
			error,
			"Invalid fields provided for payment confirmation",
			http.StatusBadRequest,
		)

		return nil, &exception
	}

	processor, error := payments.ParseProcessor(request.Gateway)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewAPIExceptionFromString(
			"Payment verification failed, please try again later.",
			http.StatusBadRequest,
		)

		return nil, &exception
	}

	confirmedPayment, error := p.verifyPayment(
		*processor,
		request.TransactionReference,
	)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewAPIExceptionFromString(
			"Payment verification failed from processor, please try again later.",
			http.StatusBadRequest,
		)

		return nil, &exception
	}

	return confirmedPayment, nil
}
