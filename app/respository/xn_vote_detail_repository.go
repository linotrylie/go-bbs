package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type xnVoteDetailRepository struct {
	XnVoteDetail *model.XnVoteDetail
	Pager        *Pager
	Repo         Repository
}

var XnVoteDetailRepository = newXnVoteDetailRepository()

func newXnVoteDetailRepository() *xnVoteDetailRepository {
	return new(xnVoteDetailRepository)
}

func (obj *xnVoteDetailRepository) Insert(xnVoteDetail model.XnVoteDetail) (rowsAffected int64, e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.Insert(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) Update(xnVoteDetail model.XnVoteDetail) (rowsAffected int64, e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.Update(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) FindByLocation(xnVoteDetail model.XnVoteDetail) (e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.FindByLocation(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) DeleteByLocation(xnVoteDetail model.XnVoteDetail) (rowsAffected int64, e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.Update(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	XnVoteDetailRepository.Repo.Model = &model.XnVoteDetail{}
	return XnVoteDetailRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *xnVoteDetailRepository) SaveInRedis(xnVoteDetail model.XnVoteDetail) (e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.SaveInRedis(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) FindInRedis(xnVoteDetail model.XnVoteDetail) (e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.FindInRedis(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) DeleteInRedis(xnVoteDetail model.XnVoteDetail) (e error) {
	XnVoteDetailRepository.Repo.Model = &xnVoteDetail
	return XnVoteDetailRepository.Repo.DeleteInRedis(&xnVoteDetail)
}

func (obj *xnVoteDetailRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	XnVoteDetailRepository.Repo.Model = &model.XnVoteDetail{}
	return XnVoteDetailRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *xnVoteDetailRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	XnVoteDetailRepository.Repo.Model = &model.XnVoteDetail{}
	return XnVoteDetailRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *xnVoteDetailRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	XnVoteDetailRepository.Repo.Model = &model.XnVoteDetail{}
	return XnVoteDetailRepository.Repo.GetDataByWhereMap(where)
}

func (obj *xnVoteDetailRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	XnVoteDetailRepository.Repo.Model = &model.XnVoteDetail{}
	return XnVoteDetailRepository.Repo.GetDataListByWhereMap(where)
}
