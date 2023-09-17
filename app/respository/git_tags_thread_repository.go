package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type gitTagsThreadRepository struct {
	GitTagsThread *model.GitTagsThread
	Pager         *Pager
	Repo          Repository
}

var GitTagsThreadRepository = newGitTagsThreadRepository()

func newGitTagsThreadRepository() *gitTagsThreadRepository {
	return new(gitTagsThreadRepository)
}

func (obj *gitTagsThreadRepository) Insert(gitTagsThread model.GitTagsThread) (rowsAffected int64, e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.Insert(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) Update(gitTagsThread model.GitTagsThread) (rowsAffected int64, e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.Update(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) FindByLocation(gitTagsThread model.GitTagsThread) (e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.FindByLocation(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) DeleteByLocation(gitTagsThread model.GitTagsThread) (rowsAffected int64, e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.Update(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	GitTagsThreadRepository.Repo.Model = &model.GitTagsThread{}
	return GitTagsThreadRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *gitTagsThreadRepository) SaveInRedis(gitTagsThread model.GitTagsThread) (e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.SaveInRedis(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) FindInRedis(gitTagsThread model.GitTagsThread) (e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.FindInRedis(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) DeleteInRedis(gitTagsThread model.GitTagsThread) (e error) {
	GitTagsThreadRepository.Repo.Model = &gitTagsThread
	return GitTagsThreadRepository.Repo.DeleteInRedis(&gitTagsThread)
}

func (obj *gitTagsThreadRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	GitTagsThreadRepository.Repo.Model = &model.GitTagsThread{}
	return GitTagsThreadRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *gitTagsThreadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	GitTagsThreadRepository.Repo.Model = &model.GitTagsThread{}
	return GitTagsThreadRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *gitTagsThreadRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	GitTagsThreadRepository.Repo.Model = &model.GitTagsThread{}
	return GitTagsThreadRepository.Repo.GetDataByWhereMap(where)
}

func (obj *gitTagsThreadRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	GitTagsThreadRepository.Repo.Model = &model.GitTagsThread{}
	return GitTagsThreadRepository.Repo.GetDataListByWhereMap(where)
}
