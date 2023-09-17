package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type userRepository struct {
	User  *model.User
	Pager *Pager
	Repo  Repository
}

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return new(userRepository)
}

func (obj *userRepository) Insert(user model.User) (rowsAffected int64, e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.Insert(&user)
}

func (obj *userRepository) Update(user model.User) (rowsAffected int64, e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.Update(&user)
}

func (obj *userRepository) FindByLocation(user model.User) (e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.FindByLocation(&user)
}

func (obj *userRepository) DeleteByLocation(user model.User) (rowsAffected int64, e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.Update(&user)
}

func (obj *userRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	UserRepository.Repo.Model = &model.User{}
	return UserRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *userRepository) SaveInRedis(user model.User) (e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.SaveInRedis(&user)
}

func (obj *userRepository) FindInRedis(user model.User) (e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.FindInRedis(&user)
}

func (obj *userRepository) DeleteInRedis(user model.User) (e error) {
	UserRepository.Repo.Model = &user
	return UserRepository.Repo.DeleteInRedis(&user)
}

func (obj *userRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	UserRepository.Repo.Model = &model.User{}
	return UserRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *userRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	UserRepository.Repo.Model = &model.User{}
	return UserRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *userRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	UserRepository.Repo.Model = &model.User{}
	return UserRepository.Repo.GetDataByWhereMap(where)
}

func (obj *userRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	UserRepository.Repo.Model = &model.User{}
	return UserRepository.Repo.GetDataListByWhereMap(where)
}
