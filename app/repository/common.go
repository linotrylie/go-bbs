package repository

import (
	"database/sql"
	"go-bbs/app/http/model"
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
	TotalPage   int
	FieldsOrder []string // []{"id desc","name asc"}
}
