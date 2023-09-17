package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type tableDayRepository struct {
	TableDay *model.TableDay
	Pager    *Pager
	Repo     Repository
}

var TableDayRepository = newTableDayRepository()

func newTableDayRepository() *tableDayRepository {
	return new(tableDayRepository)
}

func (obj *tableDayRepository) Insert(tableDay model.TableDay) (rowsAffected int64, e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.Insert(&tableDay)
}

func (obj *tableDayRepository) Update(tableDay model.TableDay) (rowsAffected int64, e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.Update(&tableDay)
}

func (obj *tableDayRepository) FindByLocation(tableDay model.TableDay) (e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.FindByLocation(&tableDay)
}

func (obj *tableDayRepository) DeleteByLocation(tableDay model.TableDay) (rowsAffected int64, e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.Update(&tableDay)
}

func (obj *tableDayRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	TableDayRepository.Repo.Model = &model.TableDay{}
	return TableDayRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *tableDayRepository) SaveInRedis(tableDay model.TableDay) (e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.SaveInRedis(&tableDay)
}

func (obj *tableDayRepository) FindInRedis(tableDay model.TableDay) (e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.FindInRedis(&tableDay)
}

func (obj *tableDayRepository) DeleteInRedis(tableDay model.TableDay) (e error) {
	TableDayRepository.Repo.Model = &tableDay
	return TableDayRepository.Repo.DeleteInRedis(&tableDay)
}

func (obj *tableDayRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	TableDayRepository.Repo.Model = &model.TableDay{}
	return TableDayRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *tableDayRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	TableDayRepository.Repo.Model = &model.TableDay{}
	return TableDayRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *tableDayRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	TableDayRepository.Repo.Model = &model.TableDay{}
	return TableDayRepository.Repo.GetDataByWhereMap(where)
}

func (obj *tableDayRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	TableDayRepository.Repo.Model = &model.TableDay{}
	return TableDayRepository.Repo.GetDataListByWhereMap(where)
}
