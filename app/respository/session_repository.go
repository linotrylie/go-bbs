package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type sessionRepository struct {
	Session *model.Session
	Pager   *Pager
	Repo    Repository
}

var SessionRepository = newSessionRepository()

func newSessionRepository() *sessionRepository {
	return new(sessionRepository)
}

func (obj *sessionRepository) Insert(session model.Session) (rowsAffected int64, e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.Insert(&session)
}

func (obj *sessionRepository) Update(session model.Session) (rowsAffected int64, e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.Update(&session)
}

func (obj *sessionRepository) FindByLocation(session model.Session) (e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.FindByLocation(&session)
}

func (obj *sessionRepository) DeleteByLocation(session model.Session) (rowsAffected int64, e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.Update(&session)
}

func (obj *sessionRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	SessionRepository.Repo.Model = &model.Session{}
	return SessionRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *sessionRepository) SaveInRedis(session model.Session) (e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.SaveInRedis(&session)
}

func (obj *sessionRepository) FindInRedis(session model.Session) (e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.FindInRedis(&session)
}

func (obj *sessionRepository) DeleteInRedis(session model.Session) (e error) {
	SessionRepository.Repo.Model = &session
	return SessionRepository.Repo.DeleteInRedis(&session)
}

func (obj *sessionRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	SessionRepository.Repo.Model = &model.Session{}
	return SessionRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *sessionRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	SessionRepository.Repo.Model = &model.Session{}
	return SessionRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *sessionRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	SessionRepository.Repo.Model = &model.Session{}
	return SessionRepository.Repo.GetDataByWhereMap(where)
}

func (obj *sessionRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	SessionRepository.Repo.Model = &model.Session{}
	return SessionRepository.Repo.GetDataListByWhereMap(where)
}
