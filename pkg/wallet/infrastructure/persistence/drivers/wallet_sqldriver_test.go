package drivers

import (
	"log"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var walletDriverInstance *WalletDriver

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// defer func() {
	// 	db.Close()
	// }()

	walletDriverInstance, err = NewWalletDriver(db)

}

func TestCreateWallet(t *testing.T) {

	wallet := entities.NewWallet(uuid.NewString(), "Tayo Adekunle")

	savedWallet, exception := walletDriverInstance.Create(wallet)

	assert.Nil(t, exception)
	assert.IsType(t, savedWallet, new(entities.Wallet))

}

func TestFindWalletById(t *testing.T) {

	id := "58a64c1e-1668-401f-a192-b0a8ac207bda"

	savedWallet, exception := walletDriverInstance.FindWalletById(id)

	assert.Nil(t, exception)
	assert.IsType(t, savedWallet, new(entities.Wallet))

}

func TestUpdateWallet(t *testing.T) {

	id := "58a64c1e-1668-401f-a192-b0a8ac207bda"

	wallet, _ := walletDriverInstance.FindWalletById(id)

	wallet.AvailableBalance = 10
	wallet.TotalBalance = 10

	savedWallet, exception := walletDriverInstance.Update(*wallet)

	assert.Nil(t, exception)
	assert.IsType(t, savedWallet, new(entities.Wallet))

}

func TestDeleteWallet(t *testing.T) {

	wallet := entities.NewWallet(uuid.NewString(), "Tayo Adekunle")

	savedWallet, _ := walletDriverInstance.Create(wallet)

	ok, exception := walletDriverInstance.Delete(savedWallet.Id)

	assert.Nil(t, exception)
	assert.True(t, ok)

}

func TestFindWalletByOwnerId(t *testing.T) {

	ownerId := "f5740aad-2045-4323-8203-043ab33d3b98"

	savedWallet, exception := walletDriverInstance.FindByOwnerId(ownerId)

	assert.Nil(t, exception)
	assert.IsType(t, savedWallet, new(entities.Wallet))

}
