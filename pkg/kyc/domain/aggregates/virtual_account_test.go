package aggregates

import (
	"suxenia-finance/pkg/kyc/enums"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVirtualAccount(t *testing.T) {

	acct := NewVirtualAccount(NewVirtualAccountRequest{
		createdBy: "Tayo Adekunle",
		ownerId:   uuid.NewString(),
		provider:  enums.PAYSTACK,
	})

	assert.IsType(t, acct, VirtualAccount{})
}
