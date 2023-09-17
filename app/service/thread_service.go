package service

import (
	"fmt"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/app/respository"
	"go-bbs/app/transform"
)

type threadService struct {
}

var ThreadService = newThreadService()

func newThreadService() *threadService {
	return new(threadService)
}
func (serv *threadService) List(fid, page, pageSize int, order, sort string) (threadList []*model.Thread, totalPage int, e error) {
	respository.ThreadRepository.Pager = &respository.Pager{Page: page, PageSize: pageSize, FieldsOrder: []string{order + " " + sort}}
	threadList, e = respository.ThreadRepository.ThreadList(fid)
	if e != nil {
		return nil, 0, e
	}
	return threadList, respository.ThreadRepository.Pager.TotalPage, nil
}
func (serv *threadService) Detail(fid, tid int) (err error) {
	respository.ThreadRepository.Thread = &model.Thread{Fid: fid, Tid: tid}
	err = respository.ThreadRepository.First()
	if err != nil {
		return
	}
	threadVo := transform.TransformThread(respository.ThreadRepository.Thread)
	//threadVo.User = transform.TransformUser(&respository.ThreadRepository.Thread.User)
	//group, _ := ServiceGroupApp.GroupService.Detail(threadVo.User.Gid)
	//threadVo.User.Group = *group

	where := map[string]interface{}{
		"tid":     tid,
		"deleted": 0,
	}
	postList, err := respository.PostRepository.GetPostListByWhere(where)
	if err != nil {
		return err
	}
	var postVo *response.PostVo
	var commentList []*response.PostVo
	for _, v := range postList {
		pv := transform.TransformPost(v)
		pv.LastUpdateUser = transform.TransformUser(&respository.PostRepository.Post.LastUpdateUser)
		group, _ := ServiceGroupApp.GroupService.Detail(pv.LastUpdateUser.Gid)
		pv.LastUpdateUser.Group = *group
		pv.User = transform.TransformUser(&respository.PostRepository.Post.CreateUser)
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
