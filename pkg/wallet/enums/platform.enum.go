package enums

import "errors"

type Platform string

func newPlatform(name string) Platform {
	return Platform(name)
}

func (t *Platform) Name() string {
	return string(*t)
}

var (
	MOBILE Platform = newPlatform("MOBILE")
	WEB    Platform = newPlatform("WEB")
	API    Platform = newPlatform("API")
)

func ParsePlatform(name string) (*Platform, error) {
	switch name {

	case MOBILE.Name():

		return &MOBILE, nil

	case WEB.Name():

		return &WEB, nil

	case API.Name():

		return &API, nil

	}

	return nil, errors.New("invalid platform selected")

}
