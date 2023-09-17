package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type mypostRepository struct {
	Mypost *model.Mypost
	Pager  *Pager
	Repo   Repository
}

var MypostRepository = newMypostRepository()

func newMypostRepository() *mypostRepository {
	return new(mypostRepository)
}

func (obj *mypostRepository) Insert(mypost model.Mypost) (rowsAffected int64, e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.Insert(&mypost)
}

func (obj *mypostRepository) Update(mypost model.Mypost) (rowsAffected int64, e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.Update(&mypost)
}

func (obj *mypostRepository) FindByLocation(mypost model.Mypost) (e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.FindByLocation(&mypost)
}

func (obj *mypostRepository) DeleteByLocation(mypost model.Mypost) (rowsAffected int64, e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.Update(&mypost)
}

func (obj *mypostRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	MypostRepository.Repo.Model = &model.Mypost{}
	return MypostRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *mypostRepository) SaveInRedis(mypost model.Mypost) (e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.SaveInRedis(&mypost)
}

func (obj *mypostRepository) FindInRedis(mypost model.Mypost) (e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.FindInRedis(&mypost)
}

func (obj *mypostRepository) DeleteInRedis(mypost model.Mypost) (e error) {
	MypostRepository.Repo.Model = &mypost
	return MypostRepository.Repo.DeleteInRedis(&mypost)
}

func (obj *mypostRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	MypostRepository.Repo.Model = &model.Mypost{}
	return MypostRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *mypostRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	MypostRepository.Repo.Model = &model.Mypost{}
	return MypostRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *mypostRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	MypostRepository.Repo.Model = &model.Mypost{}
	return MypostRepository.Repo.GetDataByWhereMap(where)
}

func (obj *mypostRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	MypostRepository.Repo.Model = &model.Mypost{}
	return MypostRepository.Repo.GetDataListByWhereMap(where)
}
