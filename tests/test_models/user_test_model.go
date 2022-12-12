package test_models

type TestGetUserResp struct {
	ID        string    `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
}

type TestSetUserUpdate struct {
	FirstName string    `json:"first_name,omitempty"`
}

type TestGetUserUpdateResp struct {
	FirstName string    `json:"first_name,omitempty"`
}

type TestCreateUser struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	VerifiedType string `json:"verified_type,omitempty"`
}

type TestCreateUserResponse struct {
	ID string `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	VerifiedType string `json:"verified_type,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}
