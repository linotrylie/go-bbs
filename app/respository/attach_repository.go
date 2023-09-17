package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type attachRepository struct {
	Attach *model.Attach
	Pager  *Pager
	Repo   Repository
}

var AttachRepository = newAttachRepository()

func newAttachRepository() *attachRepository {
	return new(attachRepository)
}

func (obj *attachRepository) Insert(attach model.Attach) (rowsAffected int64, e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.Insert(&attach)
}

func (obj *attachRepository) Update(attach model.Attach) (rowsAffected int64, e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.Update(&attach)
}

func (obj *attachRepository) FindByLocation(attach model.Attach) (e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.FindByLocation(&attach)
}

func (obj *attachRepository) DeleteByLocation(attach model.Attach) (rowsAffected int64, e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.Update(&attach)
}

func (obj *attachRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	AttachRepository.Repo.Model = &model.Attach{}
	return AttachRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *attachRepository) SaveInRedis(attach model.Attach) (e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.SaveInRedis(&attach)
}

func (obj *attachRepository) FindInRedis(attach model.Attach) (e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.FindInRedis(&attach)
}

func (obj *attachRepository) DeleteInRedis(attach model.Attach) (e error) {
	AttachRepository.Repo.Model = &attach
	return AttachRepository.Repo.DeleteInRedis(&attach)
}

func (obj *attachRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	AttachRepository.Repo.Model = &model.Attach{}
	return AttachRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *attachRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	AttachRepository.Repo.Model = &model.Attach{}
	return AttachRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *attachRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	AttachRepository.Repo.Model = &model.Attach{}
	return AttachRepository.Repo.GetDataByWhereMap(where)
}

func (obj *attachRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	AttachRepository.Repo.Model = &model.Attach{}
	return AttachRepository.Repo.GetDataListByWhereMap(where)
}
