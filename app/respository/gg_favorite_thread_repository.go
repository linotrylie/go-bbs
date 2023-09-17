package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type ggFavoriteThreadRepository struct {
	GgFavoriteThread *model.GgFavoriteThread
	Pager            *Pager
	Repo             Repository
}

var GgFavoriteThreadRepository = newGgFavoriteThreadRepository()

func newGgFavoriteThreadRepository() *ggFavoriteThreadRepository {
	return new(ggFavoriteThreadRepository)
}

func (obj *ggFavoriteThreadRepository) Insert(ggFavoriteThread model.GgFavoriteThread) (rowsAffected int64, e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.Insert(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) Update(ggFavoriteThread model.GgFavoriteThread) (rowsAffected int64, e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.Update(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) FindByLocation(ggFavoriteThread model.GgFavoriteThread) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.FindByLocation(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) DeleteByLocation(ggFavoriteThread model.GgFavoriteThread) (rowsAffected int64, e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.Update(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &model.GgFavoriteThread{}
	return GgFavoriteThreadRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *ggFavoriteThreadRepository) SaveInRedis(ggFavoriteThread model.GgFavoriteThread) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.SaveInRedis(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) FindInRedis(ggFavoriteThread model.GgFavoriteThread) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.FindInRedis(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) DeleteInRedis(ggFavoriteThread model.GgFavoriteThread) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &ggFavoriteThread
	return GgFavoriteThreadRepository.Repo.DeleteInRedis(&ggFavoriteThread)
}

func (obj *ggFavoriteThreadRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &model.GgFavoriteThread{}
	return GgFavoriteThreadRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *ggFavoriteThreadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	GgFavoriteThreadRepository.Repo.Model = &model.GgFavoriteThread{}
	return GgFavoriteThreadRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *ggFavoriteThreadRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	GgFavoriteThreadRepository.Repo.Model = &model.GgFavoriteThread{}
	return GgFavoriteThreadRepository.Repo.GetDataByWhereMap(where)
}

func (obj *ggFavoriteThreadRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	GgFavoriteThreadRepository.Repo.Model = &model.GgFavoriteThread{}
	return GgFavoriteThreadRepository.Repo.GetDataListByWhereMap(where)
}
