package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/repository"
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
	group := &model.Group{Gid: gid}
	err := repository.GroupRepository.First(group, nil)
	if err != nil {
		return nil, err
	}
	return group, nil
}
