package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type kvRepository struct {
	Kv    *model.Kv
	Pager *Pager
	Repo  Repository
}

var KvRepository = newKvRepository()

func newKvRepository() *kvRepository {
	return new(kvRepository)
}

func (obj *kvRepository) Insert(kv model.Kv) (rowsAffected int64, e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.Insert(&kv)
}

func (obj *kvRepository) Update(kv model.Kv) (rowsAffected int64, e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.Update(&kv)
}

func (obj *kvRepository) FindByLocation(kv model.Kv) (e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.FindByLocation(&kv)
}

func (obj *kvRepository) DeleteByLocation(kv model.Kv) (rowsAffected int64, e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.Update(&kv)
}

func (obj *kvRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	KvRepository.Repo.Model = &model.Kv{}
	return KvRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *kvRepository) SaveInRedis(kv model.Kv) (e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.SaveInRedis(&kv)
}

func (obj *kvRepository) FindInRedis(kv model.Kv) (e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.FindInRedis(&kv)
}

func (obj *kvRepository) DeleteInRedis(kv model.Kv) (e error) {
	KvRepository.Repo.Model = &kv
	return KvRepository.Repo.DeleteInRedis(&kv)
}

func (obj *kvRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	KvRepository.Repo.Model = &model.Kv{}
	return KvRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *kvRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	KvRepository.Repo.Model = &model.Kv{}
	return KvRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *kvRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	KvRepository.Repo.Model = &model.Kv{}
	return KvRepository.Repo.GetDataByWhereMap(where)
}

func (obj *kvRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	KvRepository.Repo.Model = &model.Kv{}
	return KvRepository.Repo.GetDataListByWhereMap(where)
}
