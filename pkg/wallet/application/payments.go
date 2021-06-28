package application

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	commonAggregates "suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/externals/payments"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/enums"
	"suxenia-finance/pkg/wallet/infrastructure/cache"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/drivers"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentApplication struct {
	driver        *drivers.PaymentDriver
	verifyPayment payments.ConfirmPayment
	mail          func(email string, payment entities.Payment) (bool, error)
	cache         cache.Cache
	// walletRepo repos.Wallet
}

func (p *PaymentApplication) RetrivePaymentById(id string) (*entities.Payment, *structs.APIException) {
	payment, error := p.driver.FindById(id)

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

func (p *PaymentApplication) IntitalizePayment(profile commonAggregates.AuthorizeProfile, request dtos.IntitalizePaymentRequest) (*dtos.InitializedPayment, *structs.APIException) {

	email, ok := profile.GetEmail().GetAddress()

	if !ok {
		exception := structs.NewUnAuthorizedException(errors.New("incomplete authorization profile, please contact admin"))
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
		exception := structs.NewBadRequestException(errors.New("mininum amount for a payment transaction is #1000.00"))
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
	}

	error = p.cache.Put(intializedPayment.TransactionReference, intializedPayment)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewInternalServerException(errors.New("error occurred intitalizing transaction, please try again later"))
		return nil, &exception
	}

	return &intializedPayment, nil
}

func (p *PaymentApplication) ConfirmPayment(profile commonAggregates.AuthorizeProfile, request dtos.ConfirmPayment) (*entities.Payment, *structs.APIException) {

	if request.OwnerId == nil {
		exception := structs.NewAPIExceptionFromString("Transaction OwnerId is missing", http.StatusUnavailableForLegalReasons)
		return nil, &exception
	}

	if ok, error := utils.Validate(request); !ok {

		exception := structs.NewRawAPIException(error, "Invalid fields provided for payment confirmation", http.StatusBadRequest)

		return nil, &exception
	}

	processor, error := payments.ParseProcessor(request.Gateway)

	if error != nil {

		utils.LoggerInstance.Error(error)

		exception := structs.NewAPIExceptionFromString("Payment verification failed, please try again later.", http.StatusBadRequest)
		return nil, &exception
	}

	confirmedPayment, error := p.verifyPayment(*processor, request.TransactionReference)

	if error != nil {
		utils.LoggerInstance.Error(error)

		exception := structs.NewAPIExceptionFromString("Payment verification failed from processor, please try again later.", http.StatusBadRequest)
		return nil, &exception
	}

	payment, dbException := p.driver.FindByReference(confirmedPayment.TransactionReference)

	if dbException != nil {
		exception := structs.NewInternalServerException(dbException)
		return nil, &exception
	}

	if payment != nil {
		exception := structs.NewAPIExceptionFromString("Payment Transaction as already been processed, please contact admin for any complain.", http.StatusConflict)
		return nil, &exception
	}

	name, ok := profile.GetFullName()

	if !ok {
		exception := structs.NewAPIExceptionFromString("Incomplete authorization profile, Full Name is missing, please contact admin for any complain.", http.StatusUnauthorized)
		return nil, &exception
	}

	// intializedTransaction, error := p.cache.Retrieve(confirmedPayment.TransactionReference)

	// if error != nil {
	// 	utils.LoggerInstance.Error(error)
	// }

	// var transactDetail dtos.InitializedPayment = intializedTransaction

	var transactDetail dtos.InitializedPayment = dtos.InitializedPayment{}

	newPayment := entities.NewPayment(*request.OwnerId, *name)

	newPayment.Comments = "Payment has been verified but not proccessed by the wallet yet."

	newPayment.Amount = int(confirmedPayment.Amount.BigInt().Int64())

	newPayment.Status = enums.PENDING
	newPayment.TransactionReference = confirmedPayment.TransactionReference
	newPayment.SourceReference = sql.NullString{String: transactDetail.SourceReference}
	newPayment.Platform = transactDetail.Platform
	newPayment.TransactionSource = transactDetail.Source
	newPayment.ProcessedBy = confirmedPayment.ProcessedBy

	savedPayment, exception := p.driver.Create(newPayment)

	if exception != nil {

		utils.LoggerInstance.Error(exception)

		if exception.IsExecError {

			apiError := structs.NewInternalServerException(errors.New("error occurred while processing payment"))

			return nil, &apiError
		}

		apiError := structs.NewBadRequestException(dbException)

		return nil, &apiError
	}

	return savedPayment, nil
}
