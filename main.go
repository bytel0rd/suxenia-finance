package main

import (
	"fmt"
	"suxenia-finance/pkg/kyc/domain/aggregates"
)

// This  project is been implemented with domain driven design (DDD)
// corresponding test are beside the aggregates or value objects
// the code are under the pkg/kyc/domain and pkg/common at the moment
// the reason for that is the project is been modeled such that the business logic is pure and
// independent of other service dependencies.
// The justification for DDD is because is project is aimed towards a reliable financial systems
// from my experience th fastest way to loose trust is inaccurate financial accounting

func main() {

	bank := aggregates.NewBankingKYC("test-ownerId", "Oyegoke Abiodun")

	fmt.Println(bank.GetName())
}
