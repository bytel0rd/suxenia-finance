package drivers

import (
	"fmt"
	"log"
	"suxenia-finance/pkg/common/infrastructure/logs"
	"suxenia-finance/pkg/kyc/enums"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var VirtualAccountDriverInstance *VirtualAccountDriver

func init() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// defer func() {
	// 	db.Close()
	// }()

	VirtualAccountDriverInstance, err = NewVirtualAccountDriver(db, logs.NewLogger())

}

func TestCreateVirtualAccount(t *testing.T) {

	virtualAccount := entities.NewVirtualAccountEntity()

	virtualAccount.AccountName = fmt.Sprintf("Test-%s-Name", uuid.NewString())
	virtualAccount.AccountNumber = "0123333333"
	virtualAccount.BankName = "WEMA"
	virtualAccount.Provider = enums.PAYSTACK.GetName()
	virtualAccount.Reference = enums.GenerateVirtualAccountReference(enums.PAYSTACK)
	virtualAccount.OwnerId = uuid.NewString()
	virtualAccount.AuditInfo.CreatedBy = "Tayo Adekunle"
	virtualAccount.AuditInfo.UpdatedBy = "Tayo Adekunle"

	savedAccount, error := VirtualAccountDriverInstance.Create(virtualAccount)

	assert.Nil(t, error)
	assert.Equal(t, savedAccount.Id, virtualAccount.Id)

}

func TestFindVirtualAccountById(t *testing.T) {

	id := "e6d5f114-542a-4f1f-b09e-7f87453bcc01"

	virtualAccount, error := VirtualAccountDriverInstance.FindById(id)

	assert.Nil(t, error)
	assert.Equal(t, virtualAccount.Id, id)

}

func TestUpdateVirtualAccount(t *testing.T) {

	virtualAccount, _ := VirtualAccountDriverInstance.FindById("e6d5f114-542a-4f1f-b09e-7f87453bcc01")

	virtualAccount.AccountName = fmt.Sprintf("Test-%s-Name", uuid.NewString())
	virtualAccount.AccountNumber = "0123333933"
	virtualAccount.Provider = enums.FLUTTERWAVE.GetName()
	virtualAccount.OwnerId = uuid.NewString()

	update, error := VirtualAccountDriverInstance.Update(virtualAccount)

	assert.Nil(t, error)
	assert.Equal(t, update.Id, virtualAccount.Id)

}

func TestDeleteVirtualAccount(t *testing.T) {

	VirtualAccountRecord := entities.NewVirtualAccountEntity()

	virtualAccount, _ := VirtualAccountDriverInstance.Create(VirtualAccountRecord)

	ok, error := VirtualAccountDriverInstance.Delete(virtualAccount.Id)

	assert.Nil(t, error)
	assert.True(t, ok)

}
