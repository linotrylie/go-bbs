package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type cacheRepository struct {
	Cache *model.Cache
	Pager *Pager
	Repo  Repository
}

var CacheRepository = newCacheRepository()

func newCacheRepository() *cacheRepository {
	return new(cacheRepository)
}

func (obj *cacheRepository) Insert(cache model.Cache) (rowsAffected int64, e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.Insert(&cache)
}

func (obj *cacheRepository) Update(cache model.Cache) (rowsAffected int64, e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.Update(&cache)
}

func (obj *cacheRepository) FindByLocation(cache model.Cache) (e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.FindByLocation(&cache)
}

func (obj *cacheRepository) DeleteByLocation(cache model.Cache) (rowsAffected int64, e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.Update(&cache)
}

func (obj *cacheRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	CacheRepository.Repo.Model = &model.Cache{}
	return CacheRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *cacheRepository) SaveInRedis(cache model.Cache) (e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.SaveInRedis(&cache)
}

func (obj *cacheRepository) FindInRedis(cache model.Cache) (e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.FindInRedis(&cache)
}

func (obj *cacheRepository) DeleteInRedis(cache model.Cache) (e error) {
	CacheRepository.Repo.Model = &cache
	return CacheRepository.Repo.DeleteInRedis(&cache)
}

func (obj *cacheRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	CacheRepository.Repo.Model = &model.Cache{}
	return CacheRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *cacheRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	CacheRepository.Repo.Model = &model.Cache{}
	return CacheRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *cacheRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	CacheRepository.Repo.Model = &model.Cache{}
	return CacheRepository.Repo.GetDataByWhereMap(where)
}

func (obj *cacheRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	CacheRepository.Repo.Model = &model.Cache{}
	return CacheRepository.Repo.GetDataListByWhereMap(where)
}
