package repos

import (
	"log"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"

	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Test(t *testing.T) {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=postgres dbname=suxenia-finance-staging  password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	bank, err := NewBankycRepo(db)

	if err != nil {
		panic(err)
	}

	bankRecord := entities.NewBankingKycEntity("Tayo Adekunle", uuid.NewString())

	bank.Create(bankRecord)

	// bank.FindById("9a2cce61-2b62-44cf-ab28-606b96975185")
	bank.FindById("f6c6bc77-rabd-41ae-82d6-84b1f24ccf83")

	bankRecord.Name = "Oyegoke Abiodun"

	bank.Update(&bankRecord)

	bank.Delete(bankRecord.Id)

	// db.Close()

	// fmt.Printf("err: %v \n", er)
	// fmt.Printf("value: %v \n", v)

}
