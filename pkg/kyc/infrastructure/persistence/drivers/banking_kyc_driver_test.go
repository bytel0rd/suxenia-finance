package drivers

import (
	"database/sql"
	"fmt"
	"log"
	"suxenia-finance/pkg/common/infrastructure/logs"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var BankKycDriverInstance *BankKycDriver

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// defer func() {
	// 	db.Close()
	// }()

	BankKycDriverInstance, err = NewBankycDriver(db, logs.NewLogger())

}

func TestCreateBankKyc(t *testing.T) {

	bankRecord := entities.NewBankingKycEntity("CREATE-TEST", uuid.NewString())

	kyc, error := BankKycDriverInstance.Create(bankRecord)

	assert.Nil(t, error)
	assert.Equal(t, kyc.Id, bankRecord.Id)

}

func TestFindBankById(t *testing.T) {

	id := "9a2cce61-2b62-44cf-ab28-606b96975185"

	// bank.FindById("9a2cce61-2b62-44cf-ab28-606b969751855")
	bankKyc, error := BankKycDriverInstance.FindById(id)

	assert.Nil(t, error)
	assert.Equal(t, bankKyc.Id, id)

}

func TestUpdateBankKYC(t *testing.T) {

	kyc, _ := BankKycDriverInstance.FindById("9a2cce61-2b62-44cf-ab28-606b96975185")

	kyc.Name = fmt.Sprintf("Test-%s-Name", uuid.NewString())
	kyc.BankAccountName = sql.NullString{
		String: "Oyegoke Abiodun A",
		Valid:  true,
	}
	kyc.BVN = sql.NullString{
		String: "22222222222",
		Valid:  true,
	}

	update, error := BankKycDriverInstance.Update(*kyc)

	assert.Nil(t, error)
	assert.Equal(t, update.Id, kyc.Id)

}

func TestDeleteBankKyc(t *testing.T) {

	bankRecord := entities.NewBankingKycEntity("TEST_DELETE_NAME", uuid.NewString())

	kyc, _ := BankKycDriverInstance.Create(bankRecord)

	ok, error := BankKycDriverInstance.Delete(kyc.Id)

	assert.Nil(t, error)
	assert.True(t, ok)

}
