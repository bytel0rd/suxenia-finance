package aggregates

import (
	"database/sql"
	"fmt"
	"strings"
	"suxenia-finance/pkg/common/domain/aggregates"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	common_enums "suxenia-finance/pkg/common/enums"
	"suxenia-finance/pkg/wallet/enums"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSetAvailableBalance(t *testing.T) {

	var wallet WalletAggregate = WalletAggregate{
		id:               "random-id",
		totalBalance:     decimal.NewFromInt(5),
		availableBalance: decimal.NewFromInt(5),
		version:          0,
		ownerId:          "",
		AuditData:        objects.AuditData{},
	}

	err := wallet.SetAvailableBalance(decimal.NewFromInt(10))

	assert.Error(t, err)

}

func TestSetTotalBalance(t *testing.T) {

	var wallet WalletAggregate = WalletAggregate{
		id:               "random-id",
		totalBalance:     decimal.NewFromInt(5),
		availableBalance: decimal.NewFromInt(5),
		version:          0,
		ownerId:          "",
		AuditData:        objects.AuditData{},
	}

	err := wallet.SetTotalBalance(decimal.NewFromInt(2))

	assert.Error(t, err)

}

func TestProcessPayment(t *testing.T) {

	var wallet WalletAggregate = WalletAggregate{
		id:               "random-id",
		totalBalance:     decimal.NewFromInt(10),
		availableBalance: decimal.NewFromInt(5),
		version:          0,
		ownerId:          "owner-id",
		AuditData:        objects.AuditData{},
	}

	verifiedPayment := entities.NewPayment(wallet.ownerId, "Tayo Adekunle")

	verifiedPayment.ProcessedBy = "PAYSTACK"
	verifiedPayment.Amount = 50

	verifiedPayment.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	verifiedPayment.TransactionSource = "Suxenia-HCS"
	verifiedPayment.SourceReference = sql.NullString{String: strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5]), Valid: true}

	verifiedPayment.Platform = "MOBILE"
	verifiedPayment.Status = "PENDING"
	verifiedPayment.Comments = "Testing payment update with paystack"

	payment, transaction, err := wallet.ProcessPayment(verifiedPayment)

	fmt.Printf("available balance %v", wallet.availableBalance.BigInt())

	assert.Equal(t, transaction.TransactionReference, payment.TransactionReference)

	assert.True(t, wallet.totalBalance.Equal(decimal.NewFromInt(60)))
	assert.True(t, wallet.availableBalance.Equal(decimal.NewFromInt(55)))

	assert.Equal(t, payment.OpeningBalance.Int32, int32(10))
	assert.Equal(t, transaction.OpeningBalance, 10)

	assert.Equal(t, payment.Status, enums.SUCCESS)
	assert.Nil(t, err)

}

func TestProcessWithrawal(t *testing.T) {

	var wallet WalletAggregate = WalletAggregate{
		id:               "random-id",
		totalBalance:     decimal.NewFromInt(10),
		availableBalance: decimal.NewFromInt(5),
		version:          0,
		ownerId:          "owner-id",
		AuditData:        objects.AuditData{},
	}

	withdrawalRequest := entities.NewWithdrawal(wallet.ownerId, "Tayo Adekunle")

	withdrawalRequest.ProcessedBy = "PAYSTACK"
	withdrawalRequest.Amount = 5

	withdrawalRequest.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	withdrawalRequest.TransactionSource = "Suxenia-HCS"
	withdrawalRequest.SourceReference = uuid.NewString()
	withdrawalRequest.ApprovedBy = sql.NullString{String: strings.ToUpper("SYSTEM" + uuid.NewString()[0:5]), Valid: true}

	withdrawalRequest.Platform = "MOBILE"
	withdrawalRequest.Status = enums.INITIATED
	withdrawalRequest.Comments = "Testing Withdrawal update with paystack"

	withdrawal, transaction, err := wallet.ProcessWithdrawal(withdrawalRequest)

	fmt.Printf("available balance %v", wallet.availableBalance.BigInt())

	assert.Equal(t, transaction.TransactionReference, withdrawal.TransactionReference)

	assert.True(t, wallet.availableBalance.Equal(decimal.NewFromInt(0)))
	assert.True(t, wallet.totalBalance.Equal(decimal.NewFromInt(10)))

	assert.Equal(t, withdrawal.OpeningBalance, 10)
	assert.Equal(t, transaction.OpeningBalance, 10)

	assert.Equal(t, withdrawal.Status, enums.PROCESSING)
	assert.Nil(t, err)

}

func TestCompleteWithrawal(t *testing.T) {

	authProfile := aggregates.NewAuthorizedProfile()
	authProfile.SetFullName("Tayo adekunle")

	var wallet WalletAggregate = WalletAggregate{
		id:               "random-id",
		totalBalance:     decimal.NewFromInt(10),
		availableBalance: decimal.NewFromInt(5),
		version:          0,
		ownerId:          "owner-id",
		AuditData:        objects.AuditData{},
	}

	partialWithdrawal := entities.NewWithdrawal(wallet.ownerId, "Tayo Adekunle")

	partialWithdrawal.ProcessedBy = "PAYSTACK"
	partialWithdrawal.Amount = 5
	partialWithdrawal.OwnerId = wallet.ownerId

	partialWithdrawal.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	partialWithdrawal.TransactionSource = "Suxenia-HCS"
	partialWithdrawal.SourceReference = uuid.NewString()

	partialWithdrawal.Platform = "MOBILE"
	partialWithdrawal.Status = enums.PROCESSING
	partialWithdrawal.Comments = "Testing Withdrawal update with paystack"

	partialTransaction := entities.NewWalletTransaction(uuid.NewString(), "Tayo Adekunle")

	partialTransaction.OwnerId = wallet.ownerId
	partialTransaction.TransactionType = "PAYMENT"
	partialTransaction.TransactionReference = partialWithdrawal.TransactionReference
	partialTransaction.Source = "PAYSTACK"
	partialTransaction.Amount = 5
	partialTransaction.Comments = "Testing payment update with paystack"
	partialTransaction.Platform = "MOBILE"
	partialTransaction.Status = enums.PROCESSING

	withdrawal, transaction, err := wallet.CompleteWithdrawal(partialWithdrawal, partialTransaction)

	fmt.Printf("total balance %v", wallet.totalBalance.BigInt())

	assert.Equal(t, transaction.TransactionReference, withdrawal.TransactionReference)

	assert.True(t, wallet.availableBalance.Equal(decimal.NewFromInt(5)))
	assert.True(t, wallet.totalBalance.Equal(decimal.NewFromInt(5)))

	assert.Equal(t, withdrawal.Status, enums.SUCCESS)
	assert.Nil(t, err)

}

func TestAprroveWithrawal(t *testing.T) {

	authProfile := aggregates.NewAuthorizedProfile()
	authProfile.SetFullName("Tayo adekunle")
	authProfile.SetRole(common_enums.SUPER_ADMIN)

	var wallet WalletAggregate = WalletAggregate{
		id:               "random-id",
		totalBalance:     decimal.NewFromInt(10),
		availableBalance: decimal.NewFromInt(5),
		version:          0,
		ownerId:          "owner-id",
		AuditData:        objects.AuditData{},
	}

	partialWithdrawal := entities.NewWithdrawal(wallet.ownerId, "Tayo Adekunle")

	partialWithdrawal.ProcessedBy = "PAYSTACK"
	partialWithdrawal.Amount = 5
	partialWithdrawal.OwnerId = wallet.ownerId

	partialWithdrawal.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	partialWithdrawal.TransactionSource = "Suxenia-HCS"
	partialWithdrawal.SourceReference = uuid.NewString()

	partialWithdrawal.Platform = "MOBILE"
	partialWithdrawal.Status = enums.PENDING
	partialWithdrawal.Comments = "Testing Withdrawal update with paystack"

	partialTransaction := entities.NewWalletTransaction(uuid.NewString(), "Tayo Adekunle")

	partialTransaction.OwnerId = wallet.ownerId
	partialTransaction.TransactionType = "PAYMENT"
	partialTransaction.TransactionReference = partialWithdrawal.TransactionReference
	partialTransaction.Source = "PAYSTACK"
	partialTransaction.Amount = 5
	partialTransaction.Comments = "Testing payment update with paystack"
	partialTransaction.Platform = "MOBILE"
	partialTransaction.Status = enums.PENDING

	withdrawal, transaction, err := wallet.ApproveWithdrawal(authProfile, partialWithdrawal, partialTransaction)

	name, _ := authProfile.GetFullName()
	assert.Equal(t, withdrawal.ApprovedBy.String, *name)
	assert.Equal(t, transaction.Status, enums.PROCESSING)
	assert.Equal(t, withdrawal.Status, enums.PROCESSING)
	assert.Nil(t, err)

}
