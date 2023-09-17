package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type hayaFavoriteRepository struct {
	HayaFavorite *model.HayaFavorite
	Pager        *Pager
	Repo         Repository
}

var HayaFavoriteRepository = newHayaFavoriteRepository()

func newHayaFavoriteRepository() *hayaFavoriteRepository {
	return new(hayaFavoriteRepository)
}

func (obj *hayaFavoriteRepository) Insert(hayaFavorite model.HayaFavorite) (rowsAffected int64, e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.Insert(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) Update(hayaFavorite model.HayaFavorite) (rowsAffected int64, e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.Update(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) FindByLocation(hayaFavorite model.HayaFavorite) (e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.FindByLocation(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) DeleteByLocation(hayaFavorite model.HayaFavorite) (rowsAffected int64, e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.Update(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	HayaFavoriteRepository.Repo.Model = &model.HayaFavorite{}
	return HayaFavoriteRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *hayaFavoriteRepository) SaveInRedis(hayaFavorite model.HayaFavorite) (e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.SaveInRedis(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) FindInRedis(hayaFavorite model.HayaFavorite) (e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.FindInRedis(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) DeleteInRedis(hayaFavorite model.HayaFavorite) (e error) {
	HayaFavoriteRepository.Repo.Model = &hayaFavorite
	return HayaFavoriteRepository.Repo.DeleteInRedis(&hayaFavorite)
}

func (obj *hayaFavoriteRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	HayaFavoriteRepository.Repo.Model = &model.HayaFavorite{}
	return HayaFavoriteRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *hayaFavoriteRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	HayaFavoriteRepository.Repo.Model = &model.HayaFavorite{}
	return HayaFavoriteRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *hayaFavoriteRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	HayaFavoriteRepository.Repo.Model = &model.HayaFavorite{}
	return HayaFavoriteRepository.Repo.GetDataByWhereMap(where)
}

func (obj *hayaFavoriteRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	HayaFavoriteRepository.Repo.Model = &model.HayaFavorite{}
	return HayaFavoriteRepository.Repo.GetDataListByWhereMap(where)
}
