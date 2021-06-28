package drivers

import (
	"log"
	"suxenia-finance/pkg/wallet/enums"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var WalletTransactionDriverInstance *WalletTransactionDriver

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// defer func() {
	// 	db.Close()
	// }()

	WalletTransactionDriverInstance, err = NewWalletTransactionDriver(db)

}

func TestCreateWalletTransaction(t *testing.T) {

	walletTransaction := entities.NewWalletTransaction(uuid.NewString(), "Tayo Adekunle")

	walletTransaction.TransactionType = "PAYMENT"
	walletTransaction.TransactionReference = "SZX-API-PYSK-" + uuid.NewString()[0:5]
	walletTransaction.Source = "PAYSTACK"
	walletTransaction.Amount = 50
	walletTransaction.Comments = "Testing payment update with paystack"
	walletTransaction.Platform = "MOBILE"
	walletTransaction.Status = enums.PENDING

	savedWalletTransaction, exception := WalletTransactionDriverInstance.Create(walletTransaction)

	assert.Nil(t, exception)
	assert.IsType(t, savedWalletTransaction, new(entities.WalletTransaction))

}

func TestFindWalletTransactionById(t *testing.T) {

	id := "543b11d3-9970-4289-9919-e5492ae8f4dd"

	savedWalletTransaction, exception := WalletTransactionDriverInstance.FindById(id)

	assert.Nil(t, exception)
	assert.IsType(t, savedWalletTransaction, new(entities.WalletTransaction))

}

func TestUpdateWalletTransaction(t *testing.T) {

	id := "543b11d3-9970-4289-9919-e5492ae8f4dd"

	walletTransaction, _ := WalletTransactionDriverInstance.FindById(id)

	walletTransaction.TransactionType = "PAYMENT"
	walletTransaction.TransactionReference = "SZX-API-FLWV-" + uuid.NewString()[0:5]
	walletTransaction.Source = "FLUTTERWAVE"
	walletTransaction.Amount = 100
	walletTransaction.Comments = "Testing payment update"
	walletTransaction.Platform = "WEB"
	walletTransaction.Status = enums.SUCCESS

	savedWalletTransaction, exception := WalletTransactionDriverInstance.Update(*walletTransaction)

	assert.Nil(t, exception)
	assert.IsType(t, savedWalletTransaction, new(entities.WalletTransaction))

}

func TestDeleteWalletTransaction(t *testing.T) {

	walletTransaction := entities.NewWalletTransaction(uuid.NewString(), "Tayo Adekunle")

	savedWalletTransaction, _ := WalletTransactionDriverInstance.Create(walletTransaction)

	ok, exception := WalletTransactionDriverInstance.Delete(savedWalletTransaction.Id)

	assert.Nil(t, exception)
	assert.True(t, ok)

}

func TestFindWalletTransactionByOwnerId(t *testing.T) {

	ownerId := "f5740aad-2045-4323-8203-043ab33d3b98"

	savedWalletTransaction, exception := WalletTransactionDriverInstance.FindById(ownerId)

	assert.Nil(t, exception)
	assert.IsType(t, savedWalletTransaction, new(entities.WalletTransaction))

}
