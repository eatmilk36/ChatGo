package RedisModels

import "time"

type RedisUserModel struct {
	Account    string
	Jwt        string
	CreateTime time.Time
}
