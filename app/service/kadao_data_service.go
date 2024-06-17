package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/repository"
	"go-bbs/global"
)

type kaDaoDataService struct {
}

var KaDaoDataService = newKaDaoDataService()

func newKaDaoDataService() *kaDaoDataService {
	return new(kaDaoDataService)
}

func (serv *kaDaoDataService) GetKaDaoDataList(keyword, dpi, sort, order string, page, pageSize int) (kaDaoDataList []*model.KadaoData, totalPage int64, e error) {
	//先检查是否存在相同用户名的用户
	user := model.User{Username: keyword}
	hasUser := UserService.IsHasUserByUsername(keyword, &user)
	query := ""
	kadaoDataRepo.Pager = &repository.Pager{Page: page, PageSize: pageSize, FieldsOrder: []string{order + " " + sort}}
	if !hasUser {
		query = "title like %?% and is_share = 1 and dpi = ?"
		kaDaoDataList, e = kadaoDataRepo.GetDataListByWhere(query, []interface{}{keyword, dpi}, nil)
	} else {
		if global.User.Uid == user.Uid {
			query = "uid = ?  and dpi = ?"
		} else {
			query = "uid = ? and is_share = 1 and dpi = ?"
		}
		kaDaoDataList, e = kadaoDataRepo.GetDataListByWhere(query, []interface{}{user.Uid, dpi}, nil)
	}
	if e != nil {
		return nil, 0, e
	}
	return kaDaoDataList, kadaoDataRepo.Pager.TotalPage, nil
}
