package contracts

import "time"

type NewLogin struct {
	Jwt  string    `json:"jwt"`
	Time time.Time `json:"time"`
}
