package service

import (
	"fmt"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/app/repository"
	"go-bbs/app/transform"
)

type threadService struct {
}

var ThreadService = newThreadService()

func newThreadService() *threadService {
	return new(threadService)
}
func (serv *threadService) List(fid, page, pageSize int, order, sort string) (threadList []*model.Thread, totalPage int, e error) {
	repository.ThreadRepository.Pager = &repository.Pager{Page: page, PageSize: pageSize, FieldsOrder: []string{order + " " + sort}}
	where := make(map[string]interface{})
	where["fid"] = fid
	threadList, e = repository.ThreadRepository.GetDataListByWhereMap(where)
	if e != nil {
		return nil, 0, e
	}
	return threadList, repository.ThreadRepository.Pager.TotalPage, nil
}
func (serv *threadService) Detail(fid, tid int) (err error) {
	thread := &model.Thread{Fid: fid, Tid: tid}
	err = repository.ThreadRepository.First(thread)
	if err != nil {
		return
	}
	threadVo := transform.TransformThread(thread)
	where := map[string]interface{}{
		"tid":     tid,
		"deleted": 0,
	}
	postList, err := repository.PostRepository.GetDataListByWhereMap(where)
	if err != nil {
		return err
	}
	var postVo *response.PostVo
	var commentList []*response.PostVo
	for _, v := range postList {
		pv := transform.TransformPost(v)
		pv.LastUpdateUser = transform.TransformUser(&v.LastUpdateUser)
		group, _ := ServiceGroupApp.GroupService.Detail(pv.LastUpdateUser.Gid)
		pv.LastUpdateUser.Group = *group
		pv.User = transform.TransformUser(&v.CreateUser)
		group, _ = ServiceGroupApp.GroupService.Detail(pv.User.Gid)
		pv.User.Group = *group
		if v.Isfirst == 1 {
			postVo = pv
		} else {
			commentList = append(commentList, pv)
		}
	}
	fmt.Println(threadVo, postVo, commentList)
	return
}
func (serv *threadService) name() {

}
