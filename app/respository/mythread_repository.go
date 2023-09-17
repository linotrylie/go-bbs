package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type mythreadRepository struct {
	Mythread *model.Mythread
	Pager    *Pager
	Repo     Repository
}

var MythreadRepository = newMythreadRepository()

func newMythreadRepository() *mythreadRepository {
	return new(mythreadRepository)
}

func (obj *mythreadRepository) Insert(mythread model.Mythread) (rowsAffected int64, e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.Insert(&mythread)
}

func (obj *mythreadRepository) Update(mythread model.Mythread) (rowsAffected int64, e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.Update(&mythread)
}

func (obj *mythreadRepository) FindByLocation(mythread model.Mythread) (e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.FindByLocation(&mythread)
}

func (obj *mythreadRepository) DeleteByLocation(mythread model.Mythread) (rowsAffected int64, e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.Update(&mythread)
}

func (obj *mythreadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	MythreadRepository.Repo.Model = &model.Mythread{}
	return MythreadRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *mythreadRepository) SaveInRedis(mythread model.Mythread) (e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.SaveInRedis(&mythread)
}

func (obj *mythreadRepository) FindInRedis(mythread model.Mythread) (e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.FindInRedis(&mythread)
}

func (obj *mythreadRepository) DeleteInRedis(mythread model.Mythread) (e error) {
	MythreadRepository.Repo.Model = &mythread
	return MythreadRepository.Repo.DeleteInRedis(&mythread)
}

func (obj *mythreadRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	MythreadRepository.Repo.Model = &model.Mythread{}
	return MythreadRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *mythreadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	MythreadRepository.Repo.Model = &model.Mythread{}
	return MythreadRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *mythreadRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	MythreadRepository.Repo.Model = &model.Mythread{}
	return MythreadRepository.Repo.GetDataByWhereMap(where)
}

func (obj *mythreadRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	MythreadRepository.Repo.Model = &model.Mythread{}
	return MythreadRepository.Repo.GetDataListByWhereMap(where)
}
