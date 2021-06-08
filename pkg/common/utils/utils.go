package utils

func IsValidString(s *string) bool {

	if s == nil {
		return false
	}

	if *s == "" {
		return false
	}

	return true
}

func StrToPr(s string) *string {

	return &s

}
