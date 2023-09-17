package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type tagThreadRepository struct {
	TagThread *model.TagThread
	Pager     *Pager
	Repo      Repository
}

var TagThreadRepository = newTagThreadRepository()

func newTagThreadRepository() *tagThreadRepository {
	return new(tagThreadRepository)
}

func (obj *tagThreadRepository) Insert(tagThread model.TagThread) (rowsAffected int64, e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.Insert(&tagThread)
}

func (obj *tagThreadRepository) Update(tagThread model.TagThread) (rowsAffected int64, e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.Update(&tagThread)
}

func (obj *tagThreadRepository) FindByLocation(tagThread model.TagThread) (e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.FindByLocation(&tagThread)
}

func (obj *tagThreadRepository) DeleteByLocation(tagThread model.TagThread) (rowsAffected int64, e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.Update(&tagThread)
}

func (obj *tagThreadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	TagThreadRepository.Repo.Model = &model.TagThread{}
	return TagThreadRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *tagThreadRepository) SaveInRedis(tagThread model.TagThread) (e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.SaveInRedis(&tagThread)
}

func (obj *tagThreadRepository) FindInRedis(tagThread model.TagThread) (e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.FindInRedis(&tagThread)
}

func (obj *tagThreadRepository) DeleteInRedis(tagThread model.TagThread) (e error) {
	TagThreadRepository.Repo.Model = &tagThread
	return TagThreadRepository.Repo.DeleteInRedis(&tagThread)
}

func (obj *tagThreadRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	TagThreadRepository.Repo.Model = &model.TagThread{}
	return TagThreadRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *tagThreadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	TagThreadRepository.Repo.Model = &model.TagThread{}
	return TagThreadRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *tagThreadRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	TagThreadRepository.Repo.Model = &model.TagThread{}
	return TagThreadRepository.Repo.GetDataByWhereMap(where)
}

func (obj *tagThreadRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	TagThreadRepository.Repo.Model = &model.TagThread{}
	return TagThreadRepository.Repo.GetDataListByWhereMap(where)
}
