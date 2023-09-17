package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type sessionDataRepository struct {
	SessionData *model.SessionData
	Pager       *Pager
	Repo        Repository
}

var SessionDataRepository = newSessionDataRepository()

func newSessionDataRepository() *sessionDataRepository {
	return new(sessionDataRepository)
}

func (obj *sessionDataRepository) Insert(sessionData model.SessionData) (rowsAffected int64, e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.Insert(&sessionData)
}

func (obj *sessionDataRepository) Update(sessionData model.SessionData) (rowsAffected int64, e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.Update(&sessionData)
}

func (obj *sessionDataRepository) FindByLocation(sessionData model.SessionData) (e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.FindByLocation(&sessionData)
}

func (obj *sessionDataRepository) DeleteByLocation(sessionData model.SessionData) (rowsAffected int64, e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.Update(&sessionData)
}

func (obj *sessionDataRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	SessionDataRepository.Repo.Model = &model.SessionData{}
	return SessionDataRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *sessionDataRepository) SaveInRedis(sessionData model.SessionData) (e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.SaveInRedis(&sessionData)
}

func (obj *sessionDataRepository) FindInRedis(sessionData model.SessionData) (e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.FindInRedis(&sessionData)
}

func (obj *sessionDataRepository) DeleteInRedis(sessionData model.SessionData) (e error) {
	SessionDataRepository.Repo.Model = &sessionData
	return SessionDataRepository.Repo.DeleteInRedis(&sessionData)
}

func (obj *sessionDataRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	SessionDataRepository.Repo.Model = &model.SessionData{}
	return SessionDataRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *sessionDataRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	SessionDataRepository.Repo.Model = &model.SessionData{}
	return SessionDataRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *sessionDataRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	SessionDataRepository.Repo.Model = &model.SessionData{}
	return SessionDataRepository.Repo.GetDataByWhereMap(where)
}

func (obj *sessionDataRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	SessionDataRepository.Repo.Model = &model.SessionData{}
	return SessionDataRepository.Repo.GetDataListByWhereMap(where)
}
