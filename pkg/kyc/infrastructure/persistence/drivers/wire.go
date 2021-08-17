package drivers

import (
	"github.com/google/wire"
)

var BuildSet wire.ProviderSet = wire.NewSet(NewBankycDriver, NewVirtualAccountDriver)
