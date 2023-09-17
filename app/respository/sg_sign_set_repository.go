package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type sgSignSetRepository struct {
	SgSignSet *model.SgSignSet
	Pager     *Pager
	Repo      Repository
}

var SgSignSetRepository = newSgSignSetRepository()

func newSgSignSetRepository() *sgSignSetRepository {
	return new(sgSignSetRepository)
}

func (obj *sgSignSetRepository) Insert(sgSignSet model.SgSignSet) (rowsAffected int64, e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.Insert(&sgSignSet)
}

func (obj *sgSignSetRepository) Update(sgSignSet model.SgSignSet) (rowsAffected int64, e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.Update(&sgSignSet)
}

func (obj *sgSignSetRepository) FindByLocation(sgSignSet model.SgSignSet) (e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.FindByLocation(&sgSignSet)
}

func (obj *sgSignSetRepository) DeleteByLocation(sgSignSet model.SgSignSet) (rowsAffected int64, e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.Update(&sgSignSet)
}

func (obj *sgSignSetRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	SgSignSetRepository.Repo.Model = &model.SgSignSet{}
	return SgSignSetRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *sgSignSetRepository) SaveInRedis(sgSignSet model.SgSignSet) (e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.SaveInRedis(&sgSignSet)
}

func (obj *sgSignSetRepository) FindInRedis(sgSignSet model.SgSignSet) (e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.FindInRedis(&sgSignSet)
}

func (obj *sgSignSetRepository) DeleteInRedis(sgSignSet model.SgSignSet) (e error) {
	SgSignSetRepository.Repo.Model = &sgSignSet
	return SgSignSetRepository.Repo.DeleteInRedis(&sgSignSet)
}

func (obj *sgSignSetRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	SgSignSetRepository.Repo.Model = &model.SgSignSet{}
	return SgSignSetRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *sgSignSetRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	SgSignSetRepository.Repo.Model = &model.SgSignSet{}
	return SgSignSetRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *sgSignSetRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	SgSignSetRepository.Repo.Model = &model.SgSignSet{}
	return SgSignSetRepository.Repo.GetDataByWhereMap(where)
}

func (obj *sgSignSetRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	SgSignSetRepository.Repo.Model = &model.SgSignSet{}
	return SgSignSetRepository.Repo.GetDataListByWhereMap(where)
}
