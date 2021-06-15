package structs

type AuthProfile struct {
	Email       string   `json:"email"`
	FullName    string   `json:"fullName"`
	ID          string   `json:"id"`
	OrgID       *string  `json:"orgId,omitempty"`
	Permissions []string `json:"permissions"`
	Role        string   `json:"role"`
}

func (profile *AuthProfile) Validate() error {
	return nil
}

func (profile *AuthProfile) IsValid() bool {
	return true
}
