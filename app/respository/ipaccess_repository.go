package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type ipaccessRepository struct {
	Ipaccess *model.Ipaccess
	Pager    *Pager
	Repo     Repository
}

var IpaccessRepository = newIpaccessRepository()

func newIpaccessRepository() *ipaccessRepository {
	return new(ipaccessRepository)
}

func (obj *ipaccessRepository) Insert(ipaccess model.Ipaccess) (rowsAffected int64, e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.Insert(&ipaccess)
}

func (obj *ipaccessRepository) Update(ipaccess model.Ipaccess) (rowsAffected int64, e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.Update(&ipaccess)
}

func (obj *ipaccessRepository) FindByLocation(ipaccess model.Ipaccess) (e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.FindByLocation(&ipaccess)
}

func (obj *ipaccessRepository) DeleteByLocation(ipaccess model.Ipaccess) (rowsAffected int64, e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.Update(&ipaccess)
}

func (obj *ipaccessRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	IpaccessRepository.Repo.Model = &model.Ipaccess{}
	return IpaccessRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *ipaccessRepository) SaveInRedis(ipaccess model.Ipaccess) (e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.SaveInRedis(&ipaccess)
}

func (obj *ipaccessRepository) FindInRedis(ipaccess model.Ipaccess) (e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.FindInRedis(&ipaccess)
}

func (obj *ipaccessRepository) DeleteInRedis(ipaccess model.Ipaccess) (e error) {
	IpaccessRepository.Repo.Model = &ipaccess
	return IpaccessRepository.Repo.DeleteInRedis(&ipaccess)
}

func (obj *ipaccessRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	IpaccessRepository.Repo.Model = &model.Ipaccess{}
	return IpaccessRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *ipaccessRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	IpaccessRepository.Repo.Model = &model.Ipaccess{}
	return IpaccessRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *ipaccessRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	IpaccessRepository.Repo.Model = &model.Ipaccess{}
	return IpaccessRepository.Repo.GetDataByWhereMap(where)
}

func (obj *ipaccessRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	IpaccessRepository.Repo.Model = &model.Ipaccess{}
	return IpaccessRepository.Repo.GetDataListByWhereMap(where)
}
