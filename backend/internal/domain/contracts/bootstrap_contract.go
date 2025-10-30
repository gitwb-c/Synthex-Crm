package contracts

import "time"

type BootstrapDtoRequest struct {
	CompanyName  string `json:"companyName"`
	Name         string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	EmployeeName string `json:"employeeName"`
}

type NewSetup struct {
	Time time.Time `json:"time"`
}
