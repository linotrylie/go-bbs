package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type forumAccessRepository struct {
	ForumAccess *model.ForumAccess
	Pager       *Pager
	Repo        Repository
}

var ForumAccessRepository = newForumAccessRepository()

func newForumAccessRepository() *forumAccessRepository {
	return new(forumAccessRepository)
}

func (obj *forumAccessRepository) Insert(forumAccess model.ForumAccess) (rowsAffected int64, e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.Insert(&forumAccess)
}

func (obj *forumAccessRepository) Update(forumAccess model.ForumAccess) (rowsAffected int64, e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.Update(&forumAccess)
}

func (obj *forumAccessRepository) FindByLocation(forumAccess model.ForumAccess) (e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.FindByLocation(&forumAccess)
}

func (obj *forumAccessRepository) DeleteByLocation(forumAccess model.ForumAccess) (rowsAffected int64, e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.Update(&forumAccess)
}

func (obj *forumAccessRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ForumAccessRepository.Repo.Model = &model.ForumAccess{}
	return ForumAccessRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *forumAccessRepository) SaveInRedis(forumAccess model.ForumAccess) (e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.SaveInRedis(&forumAccess)
}

func (obj *forumAccessRepository) FindInRedis(forumAccess model.ForumAccess) (e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.FindInRedis(&forumAccess)
}

func (obj *forumAccessRepository) DeleteInRedis(forumAccess model.ForumAccess) (e error) {
	ForumAccessRepository.Repo.Model = &forumAccess
	return ForumAccessRepository.Repo.DeleteInRedis(&forumAccess)
}

func (obj *forumAccessRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ForumAccessRepository.Repo.Model = &model.ForumAccess{}
	return ForumAccessRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *forumAccessRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ForumAccessRepository.Repo.Model = &model.ForumAccess{}
	return ForumAccessRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *forumAccessRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ForumAccessRepository.Repo.Model = &model.ForumAccess{}
	return ForumAccessRepository.Repo.GetDataByWhereMap(where)
}

func (obj *forumAccessRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ForumAccessRepository.Repo.Model = &model.ForumAccess{}
	return ForumAccessRepository.Repo.GetDataListByWhereMap(where)
}
