package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type xnVoteRepository struct {
	XnVote *model.XnVote
	Pager  *Pager
	Repo   Repository
}

var XnVoteRepository = newXnVoteRepository()

func newXnVoteRepository() *xnVoteRepository {
	return new(xnVoteRepository)
}

func (obj *xnVoteRepository) Insert(xnVote model.XnVote) (rowsAffected int64, e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.Insert(&xnVote)
}

func (obj *xnVoteRepository) Update(xnVote model.XnVote) (rowsAffected int64, e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.Update(&xnVote)
}

func (obj *xnVoteRepository) FindByLocation(xnVote model.XnVote) (e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.FindByLocation(&xnVote)
}

func (obj *xnVoteRepository) DeleteByLocation(xnVote model.XnVote) (rowsAffected int64, e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.Update(&xnVote)
}

func (obj *xnVoteRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	XnVoteRepository.Repo.Model = &model.XnVote{}
	return XnVoteRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *xnVoteRepository) SaveInRedis(xnVote model.XnVote) (e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.SaveInRedis(&xnVote)
}

func (obj *xnVoteRepository) FindInRedis(xnVote model.XnVote) (e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.FindInRedis(&xnVote)
}

func (obj *xnVoteRepository) DeleteInRedis(xnVote model.XnVote) (e error) {
	XnVoteRepository.Repo.Model = &xnVote
	return XnVoteRepository.Repo.DeleteInRedis(&xnVote)
}

func (obj *xnVoteRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	XnVoteRepository.Repo.Model = &model.XnVote{}
	return XnVoteRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *xnVoteRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	XnVoteRepository.Repo.Model = &model.XnVote{}
	return XnVoteRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *xnVoteRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	XnVoteRepository.Repo.Model = &model.XnVote{}
	return XnVoteRepository.Repo.GetDataByWhereMap(where)
}

func (obj *xnVoteRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	XnVoteRepository.Repo.Model = &model.XnVote{}
	return XnVoteRepository.Repo.GetDataListByWhereMap(where)
}
