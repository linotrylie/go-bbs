package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type xnVoteInfoRepository struct {
	XnVoteInfo *model.XnVoteInfo
	Pager      *Pager
	Repo       Repository
}

var XnVoteInfoRepository = newXnVoteInfoRepository()

func newXnVoteInfoRepository() *xnVoteInfoRepository {
	return new(xnVoteInfoRepository)
}

func (obj *xnVoteInfoRepository) Insert(xnVoteInfo model.XnVoteInfo) (rowsAffected int64, e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.Insert(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) Update(xnVoteInfo model.XnVoteInfo) (rowsAffected int64, e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.Update(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) FindByLocation(xnVoteInfo model.XnVoteInfo) (e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.FindByLocation(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) DeleteByLocation(xnVoteInfo model.XnVoteInfo) (rowsAffected int64, e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.Update(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	XnVoteInfoRepository.Repo.Model = &model.XnVoteInfo{}
	return XnVoteInfoRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *xnVoteInfoRepository) SaveInRedis(xnVoteInfo model.XnVoteInfo) (e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.SaveInRedis(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) FindInRedis(xnVoteInfo model.XnVoteInfo) (e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.FindInRedis(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) DeleteInRedis(xnVoteInfo model.XnVoteInfo) (e error) {
	XnVoteInfoRepository.Repo.Model = &xnVoteInfo
	return XnVoteInfoRepository.Repo.DeleteInRedis(&xnVoteInfo)
}

func (obj *xnVoteInfoRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	XnVoteInfoRepository.Repo.Model = &model.XnVoteInfo{}
	return XnVoteInfoRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *xnVoteInfoRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	XnVoteInfoRepository.Repo.Model = &model.XnVoteInfo{}
	return XnVoteInfoRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *xnVoteInfoRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	XnVoteInfoRepository.Repo.Model = &model.XnVoteInfo{}
	return XnVoteInfoRepository.Repo.GetDataByWhereMap(where)
}

func (obj *xnVoteInfoRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	XnVoteInfoRepository.Repo.Model = &model.XnVoteInfo{}
	return XnVoteInfoRepository.Repo.GetDataListByWhereMap(where)
}
