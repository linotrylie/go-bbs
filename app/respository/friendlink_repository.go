package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type friendlinkRepository struct {
	Friendlink *model.Friendlink
	Pager      *Pager
	Repo       Repository
}

var FriendlinkRepository = newFriendlinkRepository()

func newFriendlinkRepository() *friendlinkRepository {
	return new(friendlinkRepository)
}

func (obj *friendlinkRepository) Insert(friendlink model.Friendlink) (rowsAffected int64, e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.Insert(&friendlink)
}

func (obj *friendlinkRepository) Update(friendlink model.Friendlink) (rowsAffected int64, e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.Update(&friendlink)
}

func (obj *friendlinkRepository) FindByLocation(friendlink model.Friendlink) (e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.FindByLocation(&friendlink)
}

func (obj *friendlinkRepository) DeleteByLocation(friendlink model.Friendlink) (rowsAffected int64, e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.Update(&friendlink)
}

func (obj *friendlinkRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	FriendlinkRepository.Repo.Model = &model.Friendlink{}
	return FriendlinkRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *friendlinkRepository) SaveInRedis(friendlink model.Friendlink) (e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.SaveInRedis(&friendlink)
}

func (obj *friendlinkRepository) FindInRedis(friendlink model.Friendlink) (e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.FindInRedis(&friendlink)
}

func (obj *friendlinkRepository) DeleteInRedis(friendlink model.Friendlink) (e error) {
	FriendlinkRepository.Repo.Model = &friendlink
	return FriendlinkRepository.Repo.DeleteInRedis(&friendlink)
}

func (obj *friendlinkRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	FriendlinkRepository.Repo.Model = &model.Friendlink{}
	return FriendlinkRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *friendlinkRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	FriendlinkRepository.Repo.Model = &model.Friendlink{}
	return FriendlinkRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *friendlinkRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	FriendlinkRepository.Repo.Model = &model.Friendlink{}
	return FriendlinkRepository.Repo.GetDataByWhereMap(where)
}

func (obj *friendlinkRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	FriendlinkRepository.Repo.Model = &model.Friendlink{}
	return FriendlinkRepository.Repo.GetDataListByWhereMap(where)
}
