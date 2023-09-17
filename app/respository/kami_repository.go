package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type kamiRepository struct {
	Kami  *model.Kami
	Pager *Pager
	Repo  Repository
}

var KamiRepository = newKamiRepository()

func newKamiRepository() *kamiRepository {
	return new(kamiRepository)
}

func (obj *kamiRepository) Insert(kami model.Kami) (rowsAffected int64, e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.Insert(&kami)
}

func (obj *kamiRepository) Update(kami model.Kami) (rowsAffected int64, e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.Update(&kami)
}

func (obj *kamiRepository) FindByLocation(kami model.Kami) (e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.FindByLocation(&kami)
}

func (obj *kamiRepository) DeleteByLocation(kami model.Kami) (rowsAffected int64, e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.Update(&kami)
}

func (obj *kamiRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	KamiRepository.Repo.Model = &model.Kami{}
	return KamiRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *kamiRepository) SaveInRedis(kami model.Kami) (e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.SaveInRedis(&kami)
}

func (obj *kamiRepository) FindInRedis(kami model.Kami) (e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.FindInRedis(&kami)
}

func (obj *kamiRepository) DeleteInRedis(kami model.Kami) (e error) {
	KamiRepository.Repo.Model = &kami
	return KamiRepository.Repo.DeleteInRedis(&kami)
}

func (obj *kamiRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	KamiRepository.Repo.Model = &model.Kami{}
	return KamiRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *kamiRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	KamiRepository.Repo.Model = &model.Kami{}
	return KamiRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *kamiRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	KamiRepository.Repo.Model = &model.Kami{}
	return KamiRepository.Repo.GetDataByWhereMap(where)
}

func (obj *kamiRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	KamiRepository.Repo.Model = &model.Kami{}
	return KamiRepository.Repo.GetDataListByWhereMap(where)
}
