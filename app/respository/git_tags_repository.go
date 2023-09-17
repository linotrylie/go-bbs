package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type gitTagsRepository struct {
	GitTags *model.GitTags
	Pager   *Pager
	Repo    Repository
}

var GitTagsRepository = newGitTagsRepository()

func newGitTagsRepository() *gitTagsRepository {
	return new(gitTagsRepository)
}

func (obj *gitTagsRepository) Insert(gitTags model.GitTags) (rowsAffected int64, e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.Insert(&gitTags)
}

func (obj *gitTagsRepository) Update(gitTags model.GitTags) (rowsAffected int64, e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.Update(&gitTags)
}

func (obj *gitTagsRepository) FindByLocation(gitTags model.GitTags) (e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.FindByLocation(&gitTags)
}

func (obj *gitTagsRepository) DeleteByLocation(gitTags model.GitTags) (rowsAffected int64, e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.Update(&gitTags)
}

func (obj *gitTagsRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	GitTagsRepository.Repo.Model = &model.GitTags{}
	return GitTagsRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *gitTagsRepository) SaveInRedis(gitTags model.GitTags) (e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.SaveInRedis(&gitTags)
}

func (obj *gitTagsRepository) FindInRedis(gitTags model.GitTags) (e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.FindInRedis(&gitTags)
}

func (obj *gitTagsRepository) DeleteInRedis(gitTags model.GitTags) (e error) {
	GitTagsRepository.Repo.Model = &gitTags
	return GitTagsRepository.Repo.DeleteInRedis(&gitTags)
}

func (obj *gitTagsRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	GitTagsRepository.Repo.Model = &model.GitTags{}
	return GitTagsRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *gitTagsRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	GitTagsRepository.Repo.Model = &model.GitTags{}
	return GitTagsRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *gitTagsRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	GitTagsRepository.Repo.Model = &model.GitTags{}
	return GitTagsRepository.Repo.GetDataByWhereMap(where)
}

func (obj *gitTagsRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	GitTagsRepository.Repo.Model = &model.GitTags{}
	return GitTagsRepository.Repo.GetDataListByWhereMap(where)
}
