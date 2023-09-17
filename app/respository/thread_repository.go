package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type threadRepository struct {
	Thread *model.Thread
	Pager  *Pager
	Repo   Repository
}

var ThreadRepository = newThreadRepository()

func newThreadRepository() *threadRepository {
	return new(threadRepository)
}

func (obj *threadRepository) Insert(thread model.Thread) (rowsAffected int64, e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.Insert(&thread)
}

func (obj *threadRepository) Update(thread model.Thread) (rowsAffected int64, e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.Update(&thread)
}

func (obj *threadRepository) FindByLocation(thread model.Thread) (e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.FindByLocation(&thread)
}

func (obj *threadRepository) DeleteByLocation(thread model.Thread) (rowsAffected int64, e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.Update(&thread)
}

func (obj *threadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ThreadRepository.Repo.Model = &model.Thread{}
	return ThreadRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *threadRepository) SaveInRedis(thread model.Thread) (e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.SaveInRedis(&thread)
}

func (obj *threadRepository) FindInRedis(thread model.Thread) (e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.FindInRedis(&thread)
}

func (obj *threadRepository) DeleteInRedis(thread model.Thread) (e error) {
	ThreadRepository.Repo.Model = &thread
	return ThreadRepository.Repo.DeleteInRedis(&thread)
}

func (obj *threadRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ThreadRepository.Repo.Model = &model.Thread{}
	return ThreadRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *threadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ThreadRepository.Repo.Model = &model.Thread{}
	return ThreadRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *threadRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ThreadRepository.Repo.Model = &model.Thread{}
	return ThreadRepository.Repo.GetDataByWhereMap(where)
}

func (obj *threadRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ThreadRepository.Repo.Model = &model.Thread{}
	return ThreadRepository.Repo.GetDataListByWhereMap(where)
}
