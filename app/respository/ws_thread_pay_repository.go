package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type wsThreadPayRepository struct {
	WsThreadPay *model.WsThreadPay
	Pager       *Pager
	Repo        Repository
}

var WsThreadPayRepository = newWsThreadPayRepository()

func newWsThreadPayRepository() *wsThreadPayRepository {
	return new(wsThreadPayRepository)
}

func (obj *wsThreadPayRepository) Insert(wsThreadPay model.WsThreadPay) (rowsAffected int64, e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.Insert(&wsThreadPay)
}

func (obj *wsThreadPayRepository) Update(wsThreadPay model.WsThreadPay) (rowsAffected int64, e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.Update(&wsThreadPay)
}

func (obj *wsThreadPayRepository) FindByLocation(wsThreadPay model.WsThreadPay) (e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.FindByLocation(&wsThreadPay)
}

func (obj *wsThreadPayRepository) DeleteByLocation(wsThreadPay model.WsThreadPay) (rowsAffected int64, e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.Update(&wsThreadPay)
}

func (obj *wsThreadPayRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	WsThreadPayRepository.Repo.Model = &model.WsThreadPay{}
	return WsThreadPayRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *wsThreadPayRepository) SaveInRedis(wsThreadPay model.WsThreadPay) (e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.SaveInRedis(&wsThreadPay)
}

func (obj *wsThreadPayRepository) FindInRedis(wsThreadPay model.WsThreadPay) (e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.FindInRedis(&wsThreadPay)
}

func (obj *wsThreadPayRepository) DeleteInRedis(wsThreadPay model.WsThreadPay) (e error) {
	WsThreadPayRepository.Repo.Model = &wsThreadPay
	return WsThreadPayRepository.Repo.DeleteInRedis(&wsThreadPay)
}

func (obj *wsThreadPayRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	WsThreadPayRepository.Repo.Model = &model.WsThreadPay{}
	return WsThreadPayRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *wsThreadPayRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	WsThreadPayRepository.Repo.Model = &model.WsThreadPay{}
	return WsThreadPayRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *wsThreadPayRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	WsThreadPayRepository.Repo.Model = &model.WsThreadPay{}
	return WsThreadPayRepository.Repo.GetDataByWhereMap(where)
}

func (obj *wsThreadPayRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	WsThreadPayRepository.Repo.Model = &model.WsThreadPay{}
	return WsThreadPayRepository.Repo.GetDataListByWhereMap(where)
}
