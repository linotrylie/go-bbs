package service

import (
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/repository"
)

type kaDaoDataService struct {
}

var KaDaoDataService = newKaDaoDataService()

func newKaDaoDataService() *kaDaoDataService {
	return new(kaDaoDataService)
}

func (serv *kaDaoDataService) GetKaDaoDataList(username, sort, order string, page, pageSize int) (kaDaoDataList []*model.KadaoData, totalPage int64, e error) {
	//先检查是否存在相同用户名的用户
	user := model.User{Username: username}
	hasUser := UserService.IsHasUserByUsername(username, &user)
	if !hasUser {
		return nil, 0, exceptions.UserNotFound
	}
	kadaoDataRepo.Pager = &repository.Pager{Page: page, PageSize: pageSize, FieldsOrder: []string{order + " " + sort}}
	kaDaoDataList, e = kadaoDataRepo.GetDataListByWhere("uid = ?", []interface{}{user.Uid}, nil)
	if e != nil {
		return nil, 0, e
	}
	return kaDaoDataList, kadaoDataRepo.Pager.TotalPage, nil
}
