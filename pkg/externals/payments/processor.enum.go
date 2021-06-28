package payments

import (
	"errors"
	"suxenia-finance/pkg/common/utils"
)

type Processor string

func newProcessor(name string) Processor {
	return Processor(name)
}

func (t *Processor) Name() string {
	return string(*t)
}

var (
	PAYSTACK    Processor = newProcessor("PAYSTACK")
	FLUTTERWAVE Processor = newProcessor("FLUTTERWAVE")
)

func ProcessorShortCode(processor Processor) (*string, error) {
	switch processor.Name() {

	case PAYSTACK.Name():

		return utils.StrToPr("PYSK"), nil

	case FLUTTERWAVE.Name():

		return utils.StrToPr("FLWV"), nil

	}

	return nil, errors.New("invalid payment processor provided")

}

func ParseProcessor(name string) (*Processor, error) {
	switch name {

	case PAYSTACK.Name():

		return &PAYSTACK, nil

	case FLUTTERWAVE.Name():

		return &FLUTTERWAVE, nil

	}

	return nil, errors.New("invalid payment processor provided")

}
