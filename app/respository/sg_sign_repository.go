package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type sgSignRepository struct {
	SgSign *model.SgSign
	Pager  *Pager
	Repo   Repository
}

var SgSignRepository = newSgSignRepository()

func newSgSignRepository() *sgSignRepository {
	return new(sgSignRepository)
}

func (obj *sgSignRepository) Insert(sgSign model.SgSign) (rowsAffected int64, e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.Insert(&sgSign)
}

func (obj *sgSignRepository) Update(sgSign model.SgSign) (rowsAffected int64, e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.Update(&sgSign)
}

func (obj *sgSignRepository) FindByLocation(sgSign model.SgSign) (e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.FindByLocation(&sgSign)
}

func (obj *sgSignRepository) DeleteByLocation(sgSign model.SgSign) (rowsAffected int64, e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.Update(&sgSign)
}

func (obj *sgSignRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	SgSignRepository.Repo.Model = &model.SgSign{}
	return SgSignRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *sgSignRepository) SaveInRedis(sgSign model.SgSign) (e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.SaveInRedis(&sgSign)
}

func (obj *sgSignRepository) FindInRedis(sgSign model.SgSign) (e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.FindInRedis(&sgSign)
}

func (obj *sgSignRepository) DeleteInRedis(sgSign model.SgSign) (e error) {
	SgSignRepository.Repo.Model = &sgSign
	return SgSignRepository.Repo.DeleteInRedis(&sgSign)
}

func (obj *sgSignRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	SgSignRepository.Repo.Model = &model.SgSign{}
	return SgSignRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *sgSignRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	SgSignRepository.Repo.Model = &model.SgSign{}
	return SgSignRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *sgSignRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	SgSignRepository.Repo.Model = &model.SgSign{}
	return SgSignRepository.Repo.GetDataByWhereMap(where)
}

func (obj *sgSignRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	SgSignRepository.Repo.Model = &model.SgSign{}
	return SgSignRepository.Repo.GetDataListByWhereMap(where)
}
