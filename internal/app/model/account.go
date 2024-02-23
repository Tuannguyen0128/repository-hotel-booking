package model

type AccountQuery struct {
	ID       string
	StaffID  string
	Username string
	Page     int
	Size     int
}

type Account struct {
	ID          string  `json:"id"`
	StaffID     string  `json:"staff_id"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	UserRoleID  string  `json:"user_role_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
	LastLoginAt *string `json:"last_login_at,omitempty"`
}
type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type AddAccountResponse struct {
	ID string `json:"id"`
}
