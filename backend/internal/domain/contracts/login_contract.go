package contracts

import "time"

type NewLogin struct {
	Jwt  string    `json:"jwt"`
	Time time.Time `json:"time"`
}

type LoginDtoRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
