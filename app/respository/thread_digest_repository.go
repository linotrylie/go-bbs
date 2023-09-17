package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type threadDigestRepository struct {
	ThreadDigest *model.ThreadDigest
	Pager        *Pager
	Repo         Repository
}

var ThreadDigestRepository = newThreadDigestRepository()

func newThreadDigestRepository() *threadDigestRepository {
	return new(threadDigestRepository)
}

func (obj *threadDigestRepository) Insert(threadDigest model.ThreadDigest) (rowsAffected int64, e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.Insert(&threadDigest)
}

func (obj *threadDigestRepository) Update(threadDigest model.ThreadDigest) (rowsAffected int64, e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.Update(&threadDigest)
}

func (obj *threadDigestRepository) FindByLocation(threadDigest model.ThreadDigest) (e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.FindByLocation(&threadDigest)
}

func (obj *threadDigestRepository) DeleteByLocation(threadDigest model.ThreadDigest) (rowsAffected int64, e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.Update(&threadDigest)
}

func (obj *threadDigestRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ThreadDigestRepository.Repo.Model = &model.ThreadDigest{}
	return ThreadDigestRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *threadDigestRepository) SaveInRedis(threadDigest model.ThreadDigest) (e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.SaveInRedis(&threadDigest)
}

func (obj *threadDigestRepository) FindInRedis(threadDigest model.ThreadDigest) (e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.FindInRedis(&threadDigest)
}

func (obj *threadDigestRepository) DeleteInRedis(threadDigest model.ThreadDigest) (e error) {
	ThreadDigestRepository.Repo.Model = &threadDigest
	return ThreadDigestRepository.Repo.DeleteInRedis(&threadDigest)
}

func (obj *threadDigestRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ThreadDigestRepository.Repo.Model = &model.ThreadDigest{}
	return ThreadDigestRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *threadDigestRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ThreadDigestRepository.Repo.Model = &model.ThreadDigest{}
	return ThreadDigestRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *threadDigestRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ThreadDigestRepository.Repo.Model = &model.ThreadDigest{}
	return ThreadDigestRepository.Repo.GetDataByWhereMap(where)
}

func (obj *threadDigestRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ThreadDigestRepository.Repo.Model = &model.ThreadDigest{}
	return ThreadDigestRepository.Repo.GetDataListByWhereMap(where)
}
