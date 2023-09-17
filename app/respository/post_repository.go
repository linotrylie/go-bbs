package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type postRepository struct {
	Post  *model.Post
	Pager *Pager
	Repo  Repository
}

var PostRepository = newPostRepository()

func newPostRepository() *postRepository {
	return new(postRepository)
}

func (obj *postRepository) Insert(post model.Post) (rowsAffected int64, e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.Insert(&post)
}

func (obj *postRepository) Update(post model.Post) (rowsAffected int64, e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.Update(&post)
}

func (obj *postRepository) FindByLocation(post model.Post) (e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.FindByLocation(&post)
}

func (obj *postRepository) DeleteByLocation(post model.Post) (rowsAffected int64, e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.Update(&post)
}

func (obj *postRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	PostRepository.Repo.Model = &model.Post{}
	return PostRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *postRepository) SaveInRedis(post model.Post) (e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.SaveInRedis(&post)
}

func (obj *postRepository) FindInRedis(post model.Post) (e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.FindInRedis(&post)
}

func (obj *postRepository) DeleteInRedis(post model.Post) (e error) {
	PostRepository.Repo.Model = &post
	return PostRepository.Repo.DeleteInRedis(&post)
}

func (obj *postRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	PostRepository.Repo.Model = &model.Post{}
	return PostRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *postRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	PostRepository.Repo.Model = &model.Post{}
	return PostRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *postRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	PostRepository.Repo.Model = &model.Post{}
	return PostRepository.Repo.GetDataByWhereMap(where)
}

func (obj *postRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	PostRepository.Repo.Model = &model.Post{}
	return PostRepository.Repo.GetDataListByWhereMap(where)
}
