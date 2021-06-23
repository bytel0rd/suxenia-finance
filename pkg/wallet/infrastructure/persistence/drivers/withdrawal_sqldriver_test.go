package drivers

import (
	"database/sql"
	"log"
	"strings"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var WithdrawalDriverInstance *WithdrawalDriver

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	WithdrawalDriverInstance, err = NewWithdrawalDriver(db)

}

func TestCreateWithdrawal(t *testing.T) {

	Withdrawal := entities.NewWithdrawal(uuid.NewString(), "Tayo Adekunle")

	Withdrawal.ProcessedBy = "PAYSTACK"
	Withdrawal.Amount = 50

	Withdrawal.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	Withdrawal.OpeningBalance = 10
	Withdrawal.TransactionSource = "Suxenia-HCS"
	Withdrawal.SourceReference = uuid.NewString()
	Withdrawal.ApprovedBy = sql.NullString{String: strings.ToUpper("SYSTEM" + uuid.NewString()[0:5]), Valid: true}

	Withdrawal.Platform = "MOBILE"
	Withdrawal.Status = "PENDING"
	Withdrawal.Comments = "Testing Withdrawal update with paystack"

	savedWithdrawal, exception := WithdrawalDriverInstance.Create(Withdrawal)

	assert.Nil(t, exception)
	assert.IsType(t, savedWithdrawal, new(entities.Withdrawal))

}

func TestFindWithdrawalById(t *testing.T) {

	id := "5521fcb9-2a33-4eb0-a96c-241fa8a7decf"

	savedWithdrawal, exception := WithdrawalDriverInstance.FindById(id)

	assert.Nil(t, exception)
	assert.IsType(t, savedWithdrawal, new(entities.Withdrawal))

}

func TestUpdateWithdrawal(t *testing.T) {

	id := "5521fcb9-2a33-4eb0-a96c-241fa8a7decf"

	Withdrawal, _ := WithdrawalDriverInstance.FindById(id)

	Withdrawal.ProcessedBy = "FLUTTERWAVE"
	Withdrawal.Amount = 75

	Withdrawal.TransactionReference = strings.ToUpper("SZX-API-FLTW-" + uuid.NewString()[0:5])

	Withdrawal.TransactionSource = "Suxenia-HCS"
	Withdrawal.ApprovedBy = sql.NullString{String: strings.ToUpper("SUPER_ADMIN" + uuid.NewString()[0:5]), Valid: true}
	Withdrawal.SourceReference = uuid.NewString()

	Withdrawal.Platform = "WEB"
	Withdrawal.Status = "SUCCESS"
	Withdrawal.Comments = "Testing Withdrawal update with flutterwave"

	savedWithdrawal, exception := WithdrawalDriverInstance.Update(*Withdrawal)

	assert.Nil(t, exception)
	assert.IsType(t, savedWithdrawal, new(entities.Withdrawal))

}

func TestDeleteWithdrawal(t *testing.T) {

	Withdrawal := entities.NewWithdrawal(uuid.NewString(), "Tayo Adekunle")

	Withdrawal.ProcessedBy = "PAYSTACK"
	Withdrawal.Amount = 50

	Withdrawal.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	Withdrawal.OpeningBalance = 100
	Withdrawal.TransactionSource = "Suxenia-HCS"
	Withdrawal.SourceReference = strings.ToUpper("SZX-TEST-PYSK-" + uuid.NewString()[0:5])
	Withdrawal.ApprovedBy = sql.NullString{String: strings.ToUpper("SUPER_ADMIN" + uuid.NewString()[0:5]), Valid: true}

	Withdrawal.Platform = "MOBILE"
	Withdrawal.Status = "PENDING"
	Withdrawal.Comments = "Testing Withdrawal update with paystack"

	savedWithdrawal, _ := WithdrawalDriverInstance.Create(Withdrawal)

	ok, exception := WithdrawalDriverInstance.Delete(savedWithdrawal.Id)

	assert.Nil(t, exception)
	assert.True(t, ok)

}

func TestFindWithdrawalByOwnerId(t *testing.T) {

	ownerId := "f5740aad-2045-4323-8203-043ab33d3b98"

	savedWithdrawal, exception := WithdrawalDriverInstance.FindById(ownerId)

	assert.Nil(t, exception)
	assert.IsType(t, savedWithdrawal, new(entities.Withdrawal))

}
