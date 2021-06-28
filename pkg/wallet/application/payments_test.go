package application

import (
	"strconv"
	"suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PaymentApp *PaymentApplication = nil

var Profile aggregates.AuthorizeProfile = aggregates.NewAuthorizedProfile()

func init() {

	Profile.GetEmail().SetAddress(utils.StrToPr("tayoadekunle@suxenia.com"))
	Profile.SetOrgId("random-org-id")
	Profile.SetProfileId("random-profile-id")
	Profile.SetFullName("Tayo Adekunle")

}

func TestRetrievePaymentById(t *testing.T) {

	id := "random-id"

	payment, exception := PaymentApp.RetrivePaymentById(id)

	assert.Nil(t, exception)
	assert.IsType(t, payment, new(entities.Payment))

}

func TestRetrievePaymentByIdNotFound(t *testing.T) {

	id := "random-id"

	payment, exception := PaymentApp.RetrivePaymentById(id)

	assert.Nil(t, exception)
	assert.Nil(t, payment)

}

func TestIntializePayment(t *testing.T) {

	request := dtos.IntitalizePaymentRequest{
		SourceReference: "source-reference",
		Source:          "EMR",
		Email:           utils.StrToPr("tayoadekunle@mail.com"),
		Amount:          "10000",
		Gateway:         "PAYSTACK",
		Platform:        "MOBILE",
	}

	payment, exception := PaymentApp.IntitalizePayment(Profile, request)

	assert.Nil(t, exception)
	// assert.Contains(t, payment.TransactionReference, enums.ParseTransactionStatus(enums.PAYSTACK))
	assert.Equal(t, request.Amount, strconv.FormatInt(int64(payment.Amount), 10))

}

func TestConfirmPayment(t *testing.T) {

	request := dtos.ConfirmPayment{
		TransactionReference: "",
		Gateway:              "PAYSTACK",
	}

	payment, exception := PaymentApp.ConfirmPayment(Profile, request)

	assert.Nil(t, exception)
	assert.Equal(t, payment.Amount, 1000)

}
