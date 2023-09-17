package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type userPayRepository struct {
	UserPay *model.UserPay
	Pager   *Pager
	Repo    Repository
}

var UserPayRepository = newUserPayRepository()

func newUserPayRepository() *userPayRepository {
	return new(userPayRepository)
}

func (obj *userPayRepository) Insert(userPay model.UserPay) (rowsAffected int64, e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.Insert(&userPay)
}

func (obj *userPayRepository) Update(userPay model.UserPay) (rowsAffected int64, e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.Update(&userPay)
}

func (obj *userPayRepository) FindByLocation(userPay model.UserPay) (e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.FindByLocation(&userPay)
}

func (obj *userPayRepository) DeleteByLocation(userPay model.UserPay) (rowsAffected int64, e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.Update(&userPay)
}

func (obj *userPayRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	UserPayRepository.Repo.Model = &model.UserPay{}
	return UserPayRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *userPayRepository) SaveInRedis(userPay model.UserPay) (e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.SaveInRedis(&userPay)
}

func (obj *userPayRepository) FindInRedis(userPay model.UserPay) (e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.FindInRedis(&userPay)
}

func (obj *userPayRepository) DeleteInRedis(userPay model.UserPay) (e error) {
	UserPayRepository.Repo.Model = &userPay
	return UserPayRepository.Repo.DeleteInRedis(&userPay)
}

func (obj *userPayRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	UserPayRepository.Repo.Model = &model.UserPay{}
	return UserPayRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *userPayRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	UserPayRepository.Repo.Model = &model.UserPay{}
	return UserPayRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *userPayRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	UserPayRepository.Repo.Model = &model.UserPay{}
	return UserPayRepository.Repo.GetDataByWhereMap(where)
}

func (obj *userPayRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	UserPayRepository.Repo.Model = &model.UserPay{}
	return UserPayRepository.Repo.GetDataListByWhereMap(where)
}
