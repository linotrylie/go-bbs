package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type modlogRepository struct {
	Modlog *model.Modlog
	Pager  *Pager
	Repo   Repository
}

var ModlogRepository = newModlogRepository()

func newModlogRepository() *modlogRepository {
	return new(modlogRepository)
}

func (obj *modlogRepository) Insert(modlog model.Modlog) (rowsAffected int64, e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.Insert(&modlog)
}

func (obj *modlogRepository) Update(modlog model.Modlog) (rowsAffected int64, e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.Update(&modlog)
}

func (obj *modlogRepository) FindByLocation(modlog model.Modlog) (e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.FindByLocation(&modlog)
}

func (obj *modlogRepository) DeleteByLocation(modlog model.Modlog) (rowsAffected int64, e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.Update(&modlog)
}

func (obj *modlogRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	ModlogRepository.Repo.Model = &model.Modlog{}
	return ModlogRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *modlogRepository) SaveInRedis(modlog model.Modlog) (e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.SaveInRedis(&modlog)
}

func (obj *modlogRepository) FindInRedis(modlog model.Modlog) (e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.FindInRedis(&modlog)
}

func (obj *modlogRepository) DeleteInRedis(modlog model.Modlog) (e error) {
	ModlogRepository.Repo.Model = &modlog
	return ModlogRepository.Repo.DeleteInRedis(&modlog)
}

func (obj *modlogRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	ModlogRepository.Repo.Model = &model.Modlog{}
	return ModlogRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *modlogRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	ModlogRepository.Repo.Model = &model.Modlog{}
	return ModlogRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *modlogRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	ModlogRepository.Repo.Model = &model.Modlog{}
	return ModlogRepository.Repo.GetDataByWhereMap(where)
}

func (obj *modlogRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	ModlogRepository.Repo.Model = &model.Modlog{}
	return ModlogRepository.Repo.GetDataListByWhereMap(where)
}
