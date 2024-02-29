package model

type StaffQuery struct {
	ID       string
	Position string
	Page     int
	Size     int
}

type Staff struct {
	ID          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Position    string  `json:"position"`
	Salary      float64 `json:"salary"`
	DateOfBirth string  `json:"date_of_birth"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	StartDate   string  `json:"start_date"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}

type Staffs struct {
	Staffs []Staff `json:"staffs"`
}

type AddStaffResponse struct {
	ID string `json:"id"`
}

type DeleteStaffRequest struct {
	ID string `json:"id"`
}

type DeleteStaffResponse struct {
	Result string `json:"result"`
}
