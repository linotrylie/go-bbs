package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type paylistRepository struct {
	Paylist *model.Paylist
	Pager   *Pager
	Repo    Repository
}

var PaylistRepository = newPaylistRepository()

func newPaylistRepository() *paylistRepository {
	return new(paylistRepository)
}

func (obj *paylistRepository) Insert(paylist model.Paylist) (rowsAffected int64, e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.Insert(&paylist)
}

func (obj *paylistRepository) Update(paylist model.Paylist) (rowsAffected int64, e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.Update(&paylist)
}

func (obj *paylistRepository) FindByLocation(paylist model.Paylist) (e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.FindByLocation(&paylist)
}

func (obj *paylistRepository) DeleteByLocation(paylist model.Paylist) (rowsAffected int64, e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.Update(&paylist)
}

func (obj *paylistRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	PaylistRepository.Repo.Model = &model.Paylist{}
	return PaylistRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *paylistRepository) SaveInRedis(paylist model.Paylist) (e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.SaveInRedis(&paylist)
}

func (obj *paylistRepository) FindInRedis(paylist model.Paylist) (e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.FindInRedis(&paylist)
}

func (obj *paylistRepository) DeleteInRedis(paylist model.Paylist) (e error) {
	PaylistRepository.Repo.Model = &paylist
	return PaylistRepository.Repo.DeleteInRedis(&paylist)
}

func (obj *paylistRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	PaylistRepository.Repo.Model = &model.Paylist{}
	return PaylistRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *paylistRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	PaylistRepository.Repo.Model = &model.Paylist{}
	return PaylistRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *paylistRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	PaylistRepository.Repo.Model = &model.Paylist{}
	return PaylistRepository.Repo.GetDataByWhereMap(where)
}

func (obj *paylistRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	PaylistRepository.Repo.Model = &model.Paylist{}
	return PaylistRepository.Repo.GetDataListByWhereMap(where)
}
