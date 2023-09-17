package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type threadSearchRepository struct {
	ThreadSearch *model.ThreadSearch
	Pager        *Pager
	Repo         Repository
}

var ThreadSearchRepository = newThreadSearchRepository()

func newThreadSearchRepository() *threadSearchRepository {
	return new(threadSearchRepository)
}

func (obj *threadSearchRepository) Insert(threadSearch model.ThreadSearch) (rowsAffected int64, e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.Insert(&threadSearch)
}

func (obj *threadSearchRepository) Update(threadSearch model.ThreadSearch) (rowsAffected int64, e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.Update(&threadSearch)
}

func (obj *threadSearchRepository) FindByLocation(threadSearch model.ThreadSearch) (e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.FindByLocation(&threadSearch)
}

func (obj *threadSearchRepository) DeleteByLocation(threadSearch model.ThreadSearch) (rowsAffected int64, e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.Update(&threadSearch)
}

func (obj *threadSearchRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ThreadSearchRepository.Repo.Model = &model.ThreadSearch{}
	return ThreadSearchRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *threadSearchRepository) SaveInRedis(threadSearch model.ThreadSearch) (e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.SaveInRedis(&threadSearch)
}

func (obj *threadSearchRepository) FindInRedis(threadSearch model.ThreadSearch) (e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.FindInRedis(&threadSearch)
}

func (obj *threadSearchRepository) DeleteInRedis(threadSearch model.ThreadSearch) (e error) {
	ThreadSearchRepository.Repo.Model = &threadSearch
	return ThreadSearchRepository.Repo.DeleteInRedis(&threadSearch)
}

func (obj *threadSearchRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ThreadSearchRepository.Repo.Model = &model.ThreadSearch{}
	return ThreadSearchRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *threadSearchRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ThreadSearchRepository.Repo.Model = &model.ThreadSearch{}
	return ThreadSearchRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *threadSearchRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ThreadSearchRepository.Repo.Model = &model.ThreadSearch{}
	return ThreadSearchRepository.Repo.GetDataByWhereMap(where)
}

func (obj *threadSearchRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ThreadSearchRepository.Repo.Model = &model.ThreadSearch{}
	return ThreadSearchRepository.Repo.GetDataListByWhereMap(where)
}
