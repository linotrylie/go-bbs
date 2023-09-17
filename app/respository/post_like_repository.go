package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type postLikeRepository struct {
	PostLike *model.PostLike
	Pager    *Pager
	Repo     Repository
}

var PostLikeRepository = newPostLikeRepository()

func newPostLikeRepository() *postLikeRepository {
	return new(postLikeRepository)
}

func (obj *postLikeRepository) Insert(postLike model.PostLike) (rowsAffected int64, e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.Insert(&postLike)
}

func (obj *postLikeRepository) Update(postLike model.PostLike) (rowsAffected int64, e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.Update(&postLike)
}

func (obj *postLikeRepository) FindByLocation(postLike model.PostLike) (e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.FindByLocation(&postLike)
}

func (obj *postLikeRepository) DeleteByLocation(postLike model.PostLike) (rowsAffected int64, e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.Update(&postLike)
}

func (obj *postLikeRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	PostLikeRepository.Repo.Model = &model.PostLike{}
	return PostLikeRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *postLikeRepository) SaveInRedis(postLike model.PostLike) (e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.SaveInRedis(&postLike)
}

func (obj *postLikeRepository) FindInRedis(postLike model.PostLike) (e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.FindInRedis(&postLike)
}

func (obj *postLikeRepository) DeleteInRedis(postLike model.PostLike) (e error) {
	PostLikeRepository.Repo.Model = &postLike
	return PostLikeRepository.Repo.DeleteInRedis(&postLike)
}

func (obj *postLikeRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	PostLikeRepository.Repo.Model = &model.PostLike{}
	return PostLikeRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *postLikeRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	PostLikeRepository.Repo.Model = &model.PostLike{}
	return PostLikeRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *postLikeRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	PostLikeRepository.Repo.Model = &model.PostLike{}
	return PostLikeRepository.Repo.GetDataByWhereMap(where)
}

func (obj *postLikeRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	PostLikeRepository.Repo.Model = &model.PostLike{}
	return PostLikeRepository.Repo.GetDataListByWhereMap(where)
}
