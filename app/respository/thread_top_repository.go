package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type threadTopRepository struct {
	ThreadTop *model.ThreadTop
	Pager     *Pager
	Repo      Repository
}

var ThreadTopRepository = newThreadTopRepository()

func newThreadTopRepository() *threadTopRepository {
	return new(threadTopRepository)
}

func (obj *threadTopRepository) Insert(threadTop model.ThreadTop) (rowsAffected int64, e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.Insert(&threadTop)
}

func (obj *threadTopRepository) Update(threadTop model.ThreadTop) (rowsAffected int64, e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.Update(&threadTop)
}

func (obj *threadTopRepository) FindByLocation(threadTop model.ThreadTop) (e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.FindByLocation(&threadTop)
}

func (obj *threadTopRepository) DeleteByLocation(threadTop model.ThreadTop) (rowsAffected int64, e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.Update(&threadTop)
}

func (obj *threadTopRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ThreadTopRepository.Repo.Model = &model.ThreadTop{}
	return ThreadTopRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *threadTopRepository) SaveInRedis(threadTop model.ThreadTop) (e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.SaveInRedis(&threadTop)
}

func (obj *threadTopRepository) FindInRedis(threadTop model.ThreadTop) (e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.FindInRedis(&threadTop)
}

func (obj *threadTopRepository) DeleteInRedis(threadTop model.ThreadTop) (e error) {
	ThreadTopRepository.Repo.Model = &threadTop
	return ThreadTopRepository.Repo.DeleteInRedis(&threadTop)
}

func (obj *threadTopRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ThreadTopRepository.Repo.Model = &model.ThreadTop{}
	return ThreadTopRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *threadTopRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ThreadTopRepository.Repo.Model = &model.ThreadTop{}
	return ThreadTopRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *threadTopRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ThreadTopRepository.Repo.Model = &model.ThreadTop{}
	return ThreadTopRepository.Repo.GetDataByWhereMap(where)
}

func (obj *threadTopRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ThreadTopRepository.Repo.Model = &model.ThreadTop{}
	return ThreadTopRepository.Repo.GetDataListByWhereMap(where)
}
