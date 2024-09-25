package Create

import "time"

type UserCreateRequest struct {
	Account     string    `json:"Account"`
	Password    string    `json:"Password"`
	Createdtime time.Time `json:"CreatedTime"`
}
