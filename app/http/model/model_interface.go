package model

type Model interface {
	TableName() string
	Location() map[string]interface{}
	GetChanges() map[string]interface{}
	Update(name string, value interface{})
	RedisKey() string
}
