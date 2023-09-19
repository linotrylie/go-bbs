package repository

import (
	"database/sql"
	"encoding/json"
	"go-bbs/app/http/model"
	"strconv"
)

type RepositoryInterface interface {
	Insert(model model.Model) (rowsAffected int64, e error)
	Update(model model.Model) (rowsAffected int64, e error)
	FindByLocation(model model.Model) (e error)
	DeleteByLocation(model model.Model) (rowsAffected int64, e error)
	TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error)
	SaveInRedis(model model.Model) (e error)
	FindInRedis(model model.Model) (e error)
	DeleteInRedis(model model.Model) (e error)
	SaveInRedisByKey(redisKey string, data string) (e error)
	FindInRedisByKey(redisKey string) (redisRes string, e error)
	GetDataByWhereMap(where map[string]interface{}) (e error)
	GetDataListByWhereMap(query string, args []interface{}) (list []model.User, e error)
}

type Pager struct {
	PageSize    int
	Page        int
	TotalPage   int64
	FieldsOrder []string // []{"id desc","name asc"}
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
