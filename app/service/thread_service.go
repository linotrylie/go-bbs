package service

import (
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
	var query string
	if fid < 1 {
		query = "fid > ?"
	} else {
		query = "fid = ?"
	}
	threadList, e = repository.ThreadRepository.GetDataListByWhere(query, []interface{}{fid}, []string{"User"})
	if e != nil {
		return nil, 0, e
	}
	return threadList, repository.ThreadRepository.Pager.TotalPage, nil
}
func (serv *threadService) Detail(fid, tid int) (*response.ThreadVo, *response.PostVo, []*response.PostVo, error) {
	thread := &model.Thread{Fid: fid, Tid: tid}
	err := repository.ThreadRepository.First(thread, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	threadVo := transform.TransformThread(thread)
	//获取文章详情
	post := &model.Post{}
	wherePost := map[string]interface{}{
		"tid":     tid,
		"deleted": 0,
		"isfirst": 1,
	}
	err = repository.PostRepository.GetDataByWhereMap(post, wherePost, []string{"LastUpdateUser", "CreateUser"})
	if err != nil {
		return nil, nil, nil, err
	}
	postVo := serv.GetPostTransform(post)
	//获取评论
	wherePostList := map[string]interface{}{
		"tid":     tid,
		"deleted": 0,
		"isfirst": 0,
	}
	repository.PostRepository.Pager = &repository.Pager{
		Page:        1,
		PageSize:    5,
		FieldsOrder: []string{"create_date asc"},
	}
	postList, err := repository.PostRepository.GetDataListByWhereMap(wherePostList, []string{"LastUpdateUser", "CreateUser"})
	var commentList []*response.PostVo
	if err == nil {
		for _, v := range postList {
			commentList = append(commentList, serv.GetPostTransform(v))
		}
	}
	go serv.After(thread)
	return threadVo, postVo, commentList, nil
}

func (serv *threadService) After(thread *model.Thread) {
	thread.SetViews(1)
	repository.ThreadRepository.Update(thread)
}

func (serv *threadService) GetPostTransform(post *model.Post) *response.PostVo {
	var group *model.Group
	pv := transform.TransformPost(post)
	if post.LastUpdateUser != nil {
		pv.LastUpdateUser = transform.TransformUser(post.LastUpdateUser)
		group, _ = ServiceGroupApp.GroupService.Detail(pv.LastUpdateUser.Gid)
		if group != nil {
			pv.LastUpdateUser.Group = group
		}
	}
	if post.CreateUser != nil {
		pv.User = transform.TransformUser(post.CreateUser)
		group, _ = ServiceGroupApp.GroupService.Detail(pv.User.Gid)
		if group != nil {
			pv.User.Group = group
		}
	}
	return pv
}

func (serv *threadService) name() {

}
