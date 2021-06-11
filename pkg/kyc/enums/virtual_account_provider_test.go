package enums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVirtualAccountFromName(t *testing.T) {

	assert.Equal(t, VirtualAccountProviderFromName("PAYSTACK"), PAYSTACK)
	assert.Panics(t, func() {
		VirtualAccountProviderFromName("PANIC")
	})

}

func TestVirtualAccountFromCode(t *testing.T) {

	assert.Equal(t, VirtualAccountProviderFromCode("PYSTK"), PAYSTACK)
	assert.Panics(t, func() {
		VirtualAccountProviderFromCode("PANIC")
	})

}

func TestVirtualAccountEnum(t *testing.T) {

	assert.Equal(t, PAYSTACK.GetName(), PAYSTACK.name)
	assert.Equal(t, PAYSTACK.GetCode(), PAYSTACK.code)

	assert.True(t, PAYSTACK.Equal(PAYSTACK))
	assert.False(t, PAYSTACK.Equal(FLUTTERWAVE))

}

func TestVirtualAccountReferanceGenrator(t *testing.T) {
	assert.Contains(t, GenerateVirtualAccountReference(PAYSTACK), PAYSTACK.code)
}
