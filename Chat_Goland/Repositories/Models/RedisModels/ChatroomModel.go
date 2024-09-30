package RedisModels

type RedisChatroomModel struct {
	Id   int64  `json:"id"`
	Hash string `json:"hash"`
	Name string `json:"name"`
}
