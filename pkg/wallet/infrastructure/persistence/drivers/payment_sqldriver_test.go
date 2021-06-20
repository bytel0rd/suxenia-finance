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

var PaymentDriverInstance *PaymentDriver

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	PaymentDriverInstance, err = NewPaymentDriver(db)

}

func TestCreatePayment(t *testing.T) {

	payment := entities.NewPayment(uuid.NewString(), "Tayo Adekunle")

	payment.ProcessedBy = "PAYSTACK"
	payment.Amount = 50

	payment.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	payment.OpeningBalance = sql.NullInt32{Int32: 10, Valid: true}
	payment.TransactionSource = "Suxenia-HCS"
	payment.SourceReference = sql.NullString{String: strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5]), Valid: true}

	payment.Platform = "MOBILE"
	payment.Status = "PENDING"
	payment.Comments = "Testing payment update with paystack"

	savedPayment, exception := PaymentDriverInstance.Create(payment)

	assert.Nil(t, exception)
	assert.IsType(t, savedPayment, new(entities.Payment))

}

func TestFindPaymentById(t *testing.T) {

	id := "ea9ea2f4-11f5-4de4-89e3-c1b461623c43"

	savedPayment, exception := PaymentDriverInstance.FindById(id)

	assert.Nil(t, exception)
	assert.IsType(t, savedPayment, new(entities.Payment))

}

func TestUpdatePayment(t *testing.T) {

	id := "ea9ea2f4-11f5-4de4-89e3-c1b461623c43"

	payment, _ := PaymentDriverInstance.FindById(id)

	payment.ProcessedBy = "FLUTTERWAVE"
	payment.Amount = 75

	payment.TransactionReference = strings.ToUpper("SZX-API-FLTW-" + uuid.NewString()[0:5])

	payment.TransactionSource = "Suxenia-HCS"
	payment.SourceReference = sql.NullString{String: strings.ToUpper("SZX-API-FLTW-" + uuid.NewString()[0:5]), Valid: true}

	payment.Platform = "WEB"
	payment.Status = "SUCCESS"
	payment.Comments = "Testing payment update with flutterwave"

	savedPayment, exception := PaymentDriverInstance.Update(*payment)

	assert.Nil(t, exception)
	assert.IsType(t, savedPayment, new(entities.Payment))

}

func TestDeletePayment(t *testing.T) {

	payment := entities.NewPayment(uuid.NewString(), "Tayo Adekunle")

	payment.ProcessedBy = "PAYSTACK"
	payment.Amount = 50

	payment.TransactionReference = strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5])

	payment.OpeningBalance = sql.NullInt32{Int32: 10, Valid: true}
	payment.TransactionSource = "Suxenia-HCS"
	payment.SourceReference = sql.NullString{String: strings.ToUpper("SZX-API-PYSK-" + uuid.NewString()[0:5]), Valid: true}

	payment.Platform = "MOBILE"
	payment.Status = "PENDING"
	payment.Comments = "Testing payment update with paystack"

	savedPayment, _ := PaymentDriverInstance.Create(payment)

	ok, exception := PaymentDriverInstance.Delete(savedPayment.Id)

	assert.Nil(t, exception)
	assert.True(t, ok)

}

func TestFindPaymentByOwnerId(t *testing.T) {

	ownerId := "f5740aad-2045-4323-8203-043ab33d3b98"

	savedPayment, exception := PaymentDriverInstance.FindById(ownerId)

	assert.Nil(t, exception)
	assert.IsType(t, savedPayment, new(entities.Payment))

}
