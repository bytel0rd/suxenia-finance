package enums

import (
	"fmt"
)

type VirtualAccountProvider struct {
	name string
	code string
}

func (provider *VirtualAccountProvider) GetName() string {
	return provider.name
}

func (provider *VirtualAccountProvider) GetCode() string {
	return provider.code
}

func (provider *VirtualAccountProvider) Equal(enum VirtualAccountProvider) bool {
	return provider.name == enum.name && provider.code == enum.code
}

func newVirtualAccountProvider(name string, code string) VirtualAccountProvider {
	return VirtualAccountProvider{
		name,
		code,
	}
}

var (
	PAYSTACK    = newVirtualAccountProvider("PAYSTACK", "PYSTK")
	FLUTTERWAVE = newVirtualAccountProvider("FLUTTERWAVE", "FLTWV")
	MONNIFY     = newVirtualAccountProvider("MONNIFY", "MNIFY")
)

func VirtualAccountProviderFromName(name string) VirtualAccountProvider {

	providers := []VirtualAccountProvider{PAYSTACK, FLUTTERWAVE, MONNIFY}

	for i := 0; i < len(providers); i++ {

		if providers[i].name == name {
			return providers[i]
		}

	}

	panic("Invalid Account Provider From Name")

}

func VirtualAccountProviderFromCode(code string) VirtualAccountProvider {

	providers := []VirtualAccountProvider{PAYSTACK, FLUTTERWAVE, MONNIFY}

	for i := 0; i < len(providers); i++ {

		if providers[i].code == code {
			return providers[i]
		}

	}

	panic("Invalid Account Provider From Code")

}

func GenerateVirtualAccountReference(provider VirtualAccountProvider) string {
	return fmt.Sprintf(`SZX-%v-%v`, provider.code, provider.code)
}
