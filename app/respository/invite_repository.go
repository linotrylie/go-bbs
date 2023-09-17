package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type inviteRepository struct {
	Invite *model.Invite
	Pager  *Pager
	Repo   Repository
}

var InviteRepository = newInviteRepository()

func newInviteRepository() *inviteRepository {
	return new(inviteRepository)
}

func (obj *inviteRepository) Insert(invite model.Invite) (rowsAffected int64, e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.Insert(&invite)
}

func (obj *inviteRepository) Update(invite model.Invite) (rowsAffected int64, e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.Update(&invite)
}

func (obj *inviteRepository) FindByLocation(invite model.Invite) (e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.FindByLocation(&invite)
}

func (obj *inviteRepository) DeleteByLocation(invite model.Invite) (rowsAffected int64, e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.Update(&invite)
}

func (obj *inviteRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	InviteRepository.Repo.Model = &model.Invite{}
	return InviteRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *inviteRepository) SaveInRedis(invite model.Invite) (e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.SaveInRedis(&invite)
}

func (obj *inviteRepository) FindInRedis(invite model.Invite) (e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.FindInRedis(&invite)
}

func (obj *inviteRepository) DeleteInRedis(invite model.Invite) (e error) {
	InviteRepository.Repo.Model = &invite
	return InviteRepository.Repo.DeleteInRedis(&invite)
}

func (obj *inviteRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	InviteRepository.Repo.Model = &model.Invite{}
	return InviteRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *inviteRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	InviteRepository.Repo.Model = &model.Invite{}
	return InviteRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *inviteRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	InviteRepository.Repo.Model = &model.Invite{}
	return InviteRepository.Repo.GetDataByWhereMap(where)
}

func (obj *inviteRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	InviteRepository.Repo.Model = &model.Invite{}
	return InviteRepository.Repo.GetDataListByWhereMap(where)
}
