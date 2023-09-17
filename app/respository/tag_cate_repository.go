package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type tagCateRepository struct {
	TagCate *model.TagCate
	Pager   *Pager
	Repo    Repository
}

var TagCateRepository = newTagCateRepository()

func newTagCateRepository() *tagCateRepository {
	return new(tagCateRepository)
}

func (obj *tagCateRepository) Insert(tagCate model.TagCate) (rowsAffected int64, e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.Insert(&tagCate)
}

func (obj *tagCateRepository) Update(tagCate model.TagCate) (rowsAffected int64, e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.Update(&tagCate)
}

func (obj *tagCateRepository) FindByLocation(tagCate model.TagCate) (e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.FindByLocation(&tagCate)
}

func (obj *tagCateRepository) DeleteByLocation(tagCate model.TagCate) (rowsAffected int64, e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.Update(&tagCate)
}

func (obj *tagCateRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	TagCateRepository.Repo.Model = &model.TagCate{}
	return TagCateRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *tagCateRepository) SaveInRedis(tagCate model.TagCate) (e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.SaveInRedis(&tagCate)
}

func (obj *tagCateRepository) FindInRedis(tagCate model.TagCate) (e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.FindInRedis(&tagCate)
}

func (obj *tagCateRepository) DeleteInRedis(tagCate model.TagCate) (e error) {
	TagCateRepository.Repo.Model = &tagCate
	return TagCateRepository.Repo.DeleteInRedis(&tagCate)
}

func (obj *tagCateRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	TagCateRepository.Repo.Model = &model.TagCate{}
	return TagCateRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *tagCateRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	TagCateRepository.Repo.Model = &model.TagCate{}
	return TagCateRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *tagCateRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	TagCateRepository.Repo.Model = &model.TagCate{}
	return TagCateRepository.Repo.GetDataByWhereMap(where)
}

func (obj *tagCateRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	TagCateRepository.Repo.Model = &model.TagCate{}
	return TagCateRepository.Repo.GetDataListByWhereMap(where)
}
