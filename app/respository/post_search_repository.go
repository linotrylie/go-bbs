package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type postSearchRepository struct {
	PostSearch *model.PostSearch
	Pager      *Pager
	Repo       Repository
}

var PostSearchRepository = newPostSearchRepository()

func newPostSearchRepository() *postSearchRepository {
	return new(postSearchRepository)
}

func (obj *postSearchRepository) Insert(postSearch model.PostSearch) (rowsAffected int64, e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.Insert(&postSearch)
}

func (obj *postSearchRepository) Update(postSearch model.PostSearch) (rowsAffected int64, e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.Update(&postSearch)
}

func (obj *postSearchRepository) FindByLocation(postSearch model.PostSearch) (e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.FindByLocation(&postSearch)
}

func (obj *postSearchRepository) DeleteByLocation(postSearch model.PostSearch) (rowsAffected int64, e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.Update(&postSearch)
}

func (obj *postSearchRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	PostSearchRepository.Repo.Model = &model.PostSearch{}
	return PostSearchRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *postSearchRepository) SaveInRedis(postSearch model.PostSearch) (e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.SaveInRedis(&postSearch)
}

func (obj *postSearchRepository) FindInRedis(postSearch model.PostSearch) (e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.FindInRedis(&postSearch)
}

func (obj *postSearchRepository) DeleteInRedis(postSearch model.PostSearch) (e error) {
	PostSearchRepository.Repo.Model = &postSearch
	return PostSearchRepository.Repo.DeleteInRedis(&postSearch)
}

func (obj *postSearchRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	PostSearchRepository.Repo.Model = &model.PostSearch{}
	return PostSearchRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *postSearchRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	PostSearchRepository.Repo.Model = &model.PostSearch{}
	return PostSearchRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *postSearchRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	PostSearchRepository.Repo.Model = &model.PostSearch{}
	return PostSearchRepository.Repo.GetDataByWhereMap(where)
}

func (obj *postSearchRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	PostSearchRepository.Repo.Model = &model.PostSearch{}
	return PostSearchRepository.Repo.GetDataListByWhereMap(where)
}
