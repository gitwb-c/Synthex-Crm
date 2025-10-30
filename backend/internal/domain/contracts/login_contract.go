package contracts

import "time"

type NewLogin struct {
	SessionId    string    `json:"sessionId"`
	Jwt          string    `json:"jwt"`
	EmployeeId   string    `json:"employeeId"`
	TenantId     string    `json:"tenantId"`
	DepartmentId string    `json:"departmentId"`
	Time         time.Time `json:"time"`
}

type LoginDtoRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
