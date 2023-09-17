package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type postUpdateLogRepository struct {
	PostUpdateLog *model.PostUpdateLog
	Pager         *Pager
	Repo          Repository
}

var PostUpdateLogRepository = newPostUpdateLogRepository()

func newPostUpdateLogRepository() *postUpdateLogRepository {
	return new(postUpdateLogRepository)
}

func (obj *postUpdateLogRepository) Insert(postUpdateLog model.PostUpdateLog) (rowsAffected int64, e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.Insert(&postUpdateLog)
}

func (obj *postUpdateLogRepository) Update(postUpdateLog model.PostUpdateLog) (rowsAffected int64, e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.Update(&postUpdateLog)
}

func (obj *postUpdateLogRepository) FindByLocation(postUpdateLog model.PostUpdateLog) (e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.FindByLocation(&postUpdateLog)
}

func (obj *postUpdateLogRepository) DeleteByLocation(postUpdateLog model.PostUpdateLog) (rowsAffected int64, e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.Update(&postUpdateLog)
}

func (obj *postUpdateLogRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	PostUpdateLogRepository.Repo.Model = &model.PostUpdateLog{}
	return PostUpdateLogRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *postUpdateLogRepository) SaveInRedis(postUpdateLog model.PostUpdateLog) (e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.SaveInRedis(&postUpdateLog)
}

func (obj *postUpdateLogRepository) FindInRedis(postUpdateLog model.PostUpdateLog) (e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.FindInRedis(&postUpdateLog)
}

func (obj *postUpdateLogRepository) DeleteInRedis(postUpdateLog model.PostUpdateLog) (e error) {
	PostUpdateLogRepository.Repo.Model = &postUpdateLog
	return PostUpdateLogRepository.Repo.DeleteInRedis(&postUpdateLog)
}

func (obj *postUpdateLogRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	PostUpdateLogRepository.Repo.Model = &model.PostUpdateLog{}
	return PostUpdateLogRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *postUpdateLogRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	PostUpdateLogRepository.Repo.Model = &model.PostUpdateLog{}
	return PostUpdateLogRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *postUpdateLogRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	PostUpdateLogRepository.Repo.Model = &model.PostUpdateLog{}
	return PostUpdateLogRepository.Repo.GetDataByWhereMap(where)
}

func (obj *postUpdateLogRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	PostUpdateLogRepository.Repo.Model = &model.PostUpdateLog{}
	return PostUpdateLogRepository.Repo.GetDataListByWhereMap(where)
}
