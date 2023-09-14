package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/respository"
)

type GroupService struct {
	GroupRepo respository.GroupRepository
}

func (serv *GroupService) name() {

}

func (serv *GroupService) Detail(gid int) (*model.Group, error) {
	serv.GroupRepo.Group = &model.Group{Gid: gid}
	err := serv.GroupRepo.First()
	if err != nil {
		return nil, err
	}
	return serv.GroupRepo.Group, nil
}
