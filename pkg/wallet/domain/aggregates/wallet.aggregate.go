package aggregates

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"suxenia-finance/pkg/common/domain/aggregates"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/wallet/enums"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletAggregate struct {
	id string

	totalBalance decimal.Decimal

	availableBalance decimal.Decimal

	version int

	ownerId string

	objects.AuditData
}

func NewWalletAggeregate(ownerId string) WalletAggregate {
	return WalletAggregate{
		id:      uuid.NewString(),
		ownerId: ownerId,
	}
}

func (w *WalletAggregate) GetId() string {
	return w.id
}

func (w *WalletAggregate) SetId(id string) {
	w.id = id
}

func (w *WalletAggregate) SetVersion(version int) {
	w.version = version
}

func (w *WalletAggregate) GetTotalBalanceInBankerView() decimal.Decimal {
	return w.totalBalance.RoundBank(2)
}

func (w *WalletAggregate) GetTotalBalance() decimal.Decimal {
	return w.totalBalance
}

func (w *WalletAggregate) SetTotalBalance(balance decimal.Decimal) error {

	if balance.LessThan(w.availableBalance) {
		return errors.New(`invalid operation: total balance cannot be lesser than available balance`)
	}

	w.totalBalance = balance

	return nil
}

func (w *WalletAggregate) GetAvailableBalance() decimal.Decimal {
	return w.availableBalance
}

func (w *WalletAggregate) SetAvailableBalance(balance decimal.Decimal) error {

	if balance.GreaterThan(w.totalBalance) {
		return errors.New(`invalid operation: wallet avaliable balance cannot be greater than total balance`)
	}

	w.availableBalance = balance

	return nil
}

func (w *WalletAggregate) GetOwnerId() string {
	return w.ownerId
}

func (w *WalletAggregate) SetOwnerId(ownerId string) error {

	if ownerId == "" {
		return errors.New(`missing parameter: OwnerId is required`)
	}

	return nil
}

func (w *WalletAggregate) GetVersion() int {
	return w.version
}

func (w *WalletAggregate) ProcessPayment(payment entities.Payment) (*entities.Payment, *entities.WalletTransaction, *structs.APIException) {

	ok, exception := w.validatePayment(payment)

	if !ok {
		return nil, nil, exception
	}

	payment.OpeningBalance = sql.NullInt32{Int32: int32(w.GetTotalBalance().BigInt().Int64()), Valid: true}

	transaction := w.paymentToWalletTransaction(payment)

	w.SetTotalBalance(w.totalBalance.Add(decimal.NewFromInt(int64(payment.Amount))))

	w.SetAvailableBalance(w.availableBalance.Add(decimal.NewFromInt(int64(payment.Amount))))

	payment.Status = enums.SUCCESS

	transaction.Status = payment.Status

	transaction.Comments = fmt.Sprintf(" %s created  Payment with reference %s created At %v, added to wallet at %s ",
		payment.CreatedBy, payment.TransactionReference, payment.CreatedAt, time.Now())

	payment.Comments = fmt.Sprintf("Payment as been processed by wallet and corresponding wallet transaction Id: %s was created", transaction.Id)

	return &payment, &transaction, nil
}

func (w *WalletAggregate) paymentToWalletTransaction(payment entities.Payment) entities.WalletTransaction {

	transaction := entities.NewWalletTransaction(payment.OwnerId, payment.CreatedBy)

	transaction.TransactionType = "PAYMENT"

	transaction.TransactionReference = payment.TransactionReference

	transaction.Source = payment.TransactionSource

	transaction.Platform = payment.Platform

	transaction.OpeningBalance = int(payment.OpeningBalance.Int32)

	transaction.Amount = payment.Amount

	return transaction
}

func (w *WalletAggregate) validatePayment(payment entities.Payment) (bool, *structs.APIException) {

	if payment.OwnerId != w.ownerId {

		exception := structs.NewAPIExceptionFromString("payment transaction cannot be processed for another user", http.StatusUnavailableForLegalReasons)

		return false, &exception
	}

	if payment.Status == enums.FAILED || payment.Status == enums.REJECTED {

		exception := structs.NewAPIExceptionFromString("payment transaction cannot be processed because it failed during confirmation", http.StatusBadRequest)

		return false, &exception
	}

	if payment.Status == enums.SUCCESS {

		exception := structs.NewAPIExceptionFromString("payment transaction as already been processed", http.StatusConflict)

		return false, &exception
	}

	return true, nil
}

// withdrawal (status initiated) -> pending (compared to automatic withdrawal limit) -> processing -> disturbment (callback -> success)
// partial settlement (over available_balance alone) || complete_settlement (over total_balance alone)

func (w *WalletAggregate) ProcessWithdrawal(withdrawal entities.Withdrawal, automaticLimit decimal.Decimal) (*entities.Withdrawal, *entities.WalletTransaction, *structs.APIException) {

	ok, exception := w.validateIntiatedWithdrawal(withdrawal)

	if !ok {
		return nil, nil, exception
	}

	withdrawal.OpeningBalance = int(w.totalBalance.BigInt().Int64())

	w.SetAvailableBalance(w.availableBalance.Sub(decimal.NewFromInt(int64(withdrawal.Amount))))

	if w.availableBalance.IsNegative() {

		exception := structs.NewAPIExceptionFromString("Insufficent funds to process withdrawal, please try again later.", http.StatusNotAcceptable)

		return nil, nil, &exception
	}

	transaction := w.withdrawalToWalletTransaction(withdrawal)

	withdrawal.Status = w.initiatedWithdrawalStatus(automaticLimit, withdrawal.Amount)

	transaction.Status = withdrawal.Status

	transaction.Comments = fmt.Sprintf(" %s created  withdrawal with reference %s created At %v, added to wallet at %s ",
		withdrawal.CreatedBy, withdrawal.TransactionReference, withdrawal.CreatedAt, time.Now())

	return &withdrawal, &transaction, nil

}

func (w *WalletAggregate) initiatedWithdrawalStatus(automaticLimit decimal.Decimal, withdrawalAmount int) enums.TransactionStatus {

	if automaticLimit.GreaterThan(decimal.NewFromInt(int64(withdrawalAmount))) {

		return enums.PROCESSING

	} else {

		return enums.PENDING
	}
}

func (w *WalletAggregate) withdrawalToWalletTransaction(withdrawal entities.Withdrawal) entities.WalletTransaction {

	transaction := entities.NewWalletTransaction(withdrawal.OwnerId, withdrawal.CreatedBy)

	transaction.TransactionType = "WITHDRAWAL"

	transaction.TransactionReference = withdrawal.TransactionReference

	transaction.Source = withdrawal.TransactionSource

	transaction.Platform = withdrawal.Platform

	transaction.OpeningBalance = withdrawal.OpeningBalance

	transaction.Amount = withdrawal.Amount

	return transaction
}

func (w *WalletAggregate) validateIntiatedWithdrawal(withdrawal entities.Withdrawal) (bool, *structs.APIException) {

	if w.GetOwnerId() != withdrawal.OwnerId {

		exception := structs.NewAPIExceptionFromString("withdrawal cannot be processed for another user", http.StatusUnavailableForLegalReasons)

		return false, &exception
	}

	if withdrawal.Status != enums.INITIATED {

		exception := structs.NewAPIExceptionFromString("withdrawal has already been processed, before, please try to requery the transaction", http.StatusForbidden)

		return false, &exception

	}

	return true, nil
}

func (w *WalletAggregate) ApproveWithdrawal(profile aggregates.AuthorizeProfile, withdrawal entities.Withdrawal, transaction entities.WalletTransaction) (*entities.Withdrawal, *entities.WalletTransaction, *structs.APIException) {

	name, exception := w.getAuthorizedAuditor(profile)

	if exception != nil {
		return nil, nil, exception
	}

	ok, exception := w.validateTransactionApproval(withdrawal, transaction)

	if !ok {
		return nil, nil, exception
	}

	withdrawal.Status = enums.PROCESSING

	transaction.Status = withdrawal.Status

	withdrawal.ApprovedBy = sql.NullString{String: *name, Valid: true}

	transaction.Comments = fmt.Sprintf(" %s created  withdrawal with reference %s created At %v, added to wallet at %s ",
		withdrawal.CreatedBy, withdrawal.TransactionReference, withdrawal.CreatedAt, time.Now())

	return &withdrawal, &transaction, nil

}

func (w *WalletAggregate) validateTransactionApproval(withdrawal entities.Withdrawal, transaction entities.WalletTransaction) (bool, *structs.APIException) {

	if w.GetOwnerId() != withdrawal.OwnerId {

		exception := structs.NewAPIExceptionFromString("withdrawal cannot be processed for another user", http.StatusUnavailableForLegalReasons)

		return false, &exception
	}

	if withdrawal.Status != enums.PENDING {

		exception := structs.NewAPIExceptionFromString("withdrawal does not require an admin approval to be processed", http.StatusBadRequest)

		return false, &exception

	}

	if withdrawal.TransactionReference != transaction.TransactionReference {

		exception := structs.NewAPIExceptionFromString("withdrawal cannot be completed for two different transaction histories", http.StatusInternalServerError)

		return false, &exception

	}

	return true, nil
}

func (w *WalletAggregate) getAuthorizedAuditor(profile aggregates.AuthorizeProfile) (*string, *structs.APIException) {

	name, ok := profile.GetFullName()

	if !ok {

		exception := structs.NewUnAuthorizedException(errors.New("UnAuthorized Exception: Incomplete authorization information"))

		return nil, &exception
	}

	if !profile.GetRole().IsSuperAdmin() {

		exception := structs.NewUnAuthorizedException(errors.New("UnAuthorized Exception"))

		return nil, &exception

	}
	return name, nil
}

func (w *WalletAggregate) CompleteWithdrawal(withdrawal entities.Withdrawal, transaction entities.WalletTransaction) (*entities.Withdrawal, *entities.WalletTransaction, *structs.APIException) {

	ok, exception := w.validateCompleteWithdrawal(withdrawal, transaction)

	if !ok {
		return nil, nil, exception
	}

	w.SetTotalBalance(w.totalBalance.Sub(decimal.NewFromInt(int64(withdrawal.Amount))))

	withdrawal.Status = enums.SUCCESS

	transaction.Status = withdrawal.Status

	transaction.Comments = fmt.Sprintf(" %s created  withdrawal with reference %s created At %v, added to wallet at %s ",
		withdrawal.CreatedBy, withdrawal.TransactionReference, withdrawal.CreatedAt, time.Now())

	return &withdrawal, &transaction, nil
}

func (w *WalletAggregate) validateCompleteWithdrawal(withdrawal entities.Withdrawal, transaction entities.WalletTransaction) (bool, *structs.APIException) {

	if w.GetOwnerId() != withdrawal.OwnerId {

		exception := structs.NewAPIExceptionFromString("withdrawal cannot be processed for another user", http.StatusUnavailableForLegalReasons)

		return false, &exception
	}

	if withdrawal.Status != enums.PROCESSING {

		exception := structs.NewAPIExceptionFromString("withdrawal as not been not been authorized for processing by the admins", http.StatusUnauthorized)

		return false, &exception

	}

	if withdrawal.TransactionReference != transaction.TransactionReference {

		exception := structs.NewAPIExceptionFromString("withdrawal cannot be completed for two different transaction histories", http.StatusInternalServerError)

		return false, &exception

	}

	return true, nil
}
