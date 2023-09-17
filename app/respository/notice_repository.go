package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type noticeRepository struct {
	Notice *model.Notice
	Pager  *Pager
	Repo   Repository
}

var NoticeRepository = newNoticeRepository()

func newNoticeRepository() *noticeRepository {
	return new(noticeRepository)
}

func (obj *noticeRepository) Insert(notice model.Notice) (rowsAffected int64, e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.Insert(&notice)
}

func (obj *noticeRepository) Update(notice model.Notice) (rowsAffected int64, e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.Update(&notice)
}

func (obj *noticeRepository) FindByLocation(notice model.Notice) (e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.FindByLocation(&notice)
}

func (obj *noticeRepository) DeleteByLocation(notice model.Notice) (rowsAffected int64, e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.Update(&notice)
}

func (obj *noticeRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	NoticeRepository.Repo.Model = &model.Notice{}
	return NoticeRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *noticeRepository) SaveInRedis(notice model.Notice) (e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.SaveInRedis(&notice)
}

func (obj *noticeRepository) FindInRedis(notice model.Notice) (e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.FindInRedis(&notice)
}

func (obj *noticeRepository) DeleteInRedis(notice model.Notice) (e error) {
	NoticeRepository.Repo.Model = &notice
	return NoticeRepository.Repo.DeleteInRedis(&notice)
}

func (obj *noticeRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	NoticeRepository.Repo.Model = &model.Notice{}
	return NoticeRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *noticeRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	NoticeRepository.Repo.Model = &model.Notice{}
	return NoticeRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *noticeRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	NoticeRepository.Repo.Model = &model.Notice{}
	return NoticeRepository.Repo.GetDataByWhereMap(where)
}

func (obj *noticeRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	NoticeRepository.Repo.Model = &model.Notice{}
	return NoticeRepository.Repo.GetDataListByWhereMap(where)
}
