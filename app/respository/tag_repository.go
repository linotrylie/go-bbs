package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type tagRepository struct {
	Tag   *model.Tag
	Pager *Pager
	Repo  Repository
}

var TagRepository = newTagRepository()

func newTagRepository() *tagRepository {
	return new(tagRepository)
}

func (obj *tagRepository) Insert(tag model.Tag) (rowsAffected int64, e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.Insert(&tag)
}

func (obj *tagRepository) Update(tag model.Tag) (rowsAffected int64, e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.Update(&tag)
}

func (obj *tagRepository) FindByLocation(tag model.Tag) (e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.FindByLocation(&tag)
}

func (obj *tagRepository) DeleteByLocation(tag model.Tag) (rowsAffected int64, e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.Update(&tag)
}

func (obj *tagRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	TagRepository.Repo.Model = &model.Tag{}
	return TagRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *tagRepository) SaveInRedis(tag model.Tag) (e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.SaveInRedis(&tag)
}

func (obj *tagRepository) FindInRedis(tag model.Tag) (e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.FindInRedis(&tag)
}

func (obj *tagRepository) DeleteInRedis(tag model.Tag) (e error) {
	TagRepository.Repo.Model = &tag
	return TagRepository.Repo.DeleteInRedis(&tag)
}

func (obj *tagRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	TagRepository.Repo.Model = &model.Tag{}
	return TagRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *tagRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	TagRepository.Repo.Model = &model.Tag{}
	return TagRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *tagRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	TagRepository.Repo.Model = &model.Tag{}
	return TagRepository.Repo.GetDataByWhereMap(where)
}

func (obj *tagRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	TagRepository.Repo.Model = &model.Tag{}
	return TagRepository.Repo.GetDataListByWhereMap(where)
}
