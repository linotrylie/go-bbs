package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type slideRepository struct {
	Slide *model.Slide
	Pager *Pager
	Repo  Repository
}

var SlideRepository = newSlideRepository()

func newSlideRepository() *slideRepository {
	return new(slideRepository)
}

func (obj *slideRepository) Insert(slide model.Slide) (rowsAffected int64, e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.Insert(&slide)
}

func (obj *slideRepository) Update(slide model.Slide) (rowsAffected int64, e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.Update(&slide)
}

func (obj *slideRepository) FindByLocation(slide model.Slide) (e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.FindByLocation(&slide)
}

func (obj *slideRepository) DeleteByLocation(slide model.Slide) (rowsAffected int64, e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.Update(&slide)
}

func (obj *slideRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	SlideRepository.Repo.Model = &model.Slide{}
	return SlideRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *slideRepository) SaveInRedis(slide model.Slide) (e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.SaveInRedis(&slide)
}

func (obj *slideRepository) FindInRedis(slide model.Slide) (e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.FindInRedis(&slide)
}

func (obj *slideRepository) DeleteInRedis(slide model.Slide) (e error) {
	SlideRepository.Repo.Model = &slide
	return SlideRepository.Repo.DeleteInRedis(&slide)
}

func (obj *slideRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	SlideRepository.Repo.Model = &model.Slide{}
	return SlideRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *slideRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	SlideRepository.Repo.Model = &model.Slide{}
	return SlideRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *slideRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	SlideRepository.Repo.Model = &model.Slide{}
	return SlideRepository.Repo.GetDataByWhereMap(where)
}

func (obj *slideRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	SlideRepository.Repo.Model = &model.Slide{}
	return SlideRepository.Repo.GetDataListByWhereMap(where)
}
