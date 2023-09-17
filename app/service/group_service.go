package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/respository"
)

type groupService struct {
}

var GroupService = newGroupService()

func newGroupService() *groupService {
	return new(groupService)
}
func (serv *groupService) name() {

}

func (serv *groupService) Detail(gid int) (*model.Group, error) {
	respository.GroupRepository.Group = &model.Group{Gid: gid}
	err := respository.GroupRepository.First()
	if err != nil {
		return nil, err
	}
	return respository.GroupRepository.Group, nil
}
