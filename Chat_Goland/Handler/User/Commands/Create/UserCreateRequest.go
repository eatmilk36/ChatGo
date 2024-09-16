package Create

import "time"

type UserCreateRequest struct {
	Account     string    `json:"account"`
	Password    string    `json:"password"`
	Createdtime time.Time `json:"createdTime"`
}
