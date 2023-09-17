package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type groupRepository struct {
	Group *model.Group
	Pager *Pager
	Repo  Repository
}

var GroupRepository = newGroupRepository()

func newGroupRepository() *groupRepository {
	return new(groupRepository)
}

func (obj *groupRepository) Insert(group model.Group) (rowsAffected int64, e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.Insert(&group)
}

func (obj *groupRepository) Update(group model.Group) (rowsAffected int64, e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.Update(&group)
}

func (obj *groupRepository) FindByLocation(group model.Group) (e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.FindByLocation(&group)
}

func (obj *groupRepository) DeleteByLocation(group model.Group) (rowsAffected int64, e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.Update(&group)
}

func (obj *groupRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	GroupRepository.Repo.Model = &model.Group{}
	return GroupRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *groupRepository) SaveInRedis(group model.Group) (e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.SaveInRedis(&group)
}

func (obj *groupRepository) FindInRedis(group model.Group) (e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.FindInRedis(&group)
}

func (obj *groupRepository) DeleteInRedis(group model.Group) (e error) {
	GroupRepository.Repo.Model = &group
	return GroupRepository.Repo.DeleteInRedis(&group)
}

func (obj *groupRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	GroupRepository.Repo.Model = &model.Group{}
	return GroupRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *groupRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	GroupRepository.Repo.Model = &model.Group{}
	return GroupRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *groupRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	GroupRepository.Repo.Model = &model.Group{}
	return GroupRepository.Repo.GetDataByWhereMap(where)
}

func (obj *groupRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	GroupRepository.Repo.Model = &model.Group{}
	return GroupRepository.Repo.GetDataListByWhereMap(where)
}
