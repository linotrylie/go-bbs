package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type forumRepository struct {
	Forum *model.Forum
	Pager *Pager
	Repo  Repository
}

var ForumRepository = newForumRepository()

func newForumRepository() *forumRepository {
	return new(forumRepository)
}

func (obj *forumRepository) Insert(forum model.Forum) (rowsAffected int64, e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.Insert(&forum)
}

func (obj *forumRepository) Update(forum model.Forum) (rowsAffected int64, e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.Update(&forum)
}

func (obj *forumRepository) FindByLocation(forum model.Forum) (e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.FindByLocation(&forum)
}

func (obj *forumRepository) DeleteByLocation(forum model.Forum) (rowsAffected int64, e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.Update(&forum)
}

func (obj *forumRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ForumRepository.Repo.Model = &model.Forum{}
	return ForumRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *forumRepository) SaveInRedis(forum model.Forum) (e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.SaveInRedis(&forum)
}

func (obj *forumRepository) FindInRedis(forum model.Forum) (e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.FindInRedis(&forum)
}

func (obj *forumRepository) DeleteInRedis(forum model.Forum) (e error) {
	ForumRepository.Repo.Model = &forum
	return ForumRepository.Repo.DeleteInRedis(&forum)
}

func (obj *forumRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ForumRepository.Repo.Model = &model.Forum{}
	return ForumRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *forumRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ForumRepository.Repo.Model = &model.Forum{}
	return ForumRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *forumRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ForumRepository.Repo.Model = &model.Forum{}
	return ForumRepository.Repo.GetDataByWhereMap(where)
}

func (obj *forumRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ForumRepository.Repo.Model = &model.Forum{}
	return ForumRepository.Repo.GetDataListByWhereMap(where)
}
