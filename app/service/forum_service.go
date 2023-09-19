package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/app/repository"
	"go-bbs/app/transform"
)

type forumService struct {
}

var ForumService = newForumService()

func newForumService() *forumService {
	return new(forumService)
}

func (serv *forumService) ThreadList(fid, page, pageSize int, order, sort string) (*model.Forum, []*response.ThreadVo, int64, error) {
	forumModel := &model.Forum{}
	if fid <= 0 {
		forumModel.Fid = 0
		forumModel.Name = "全部"
	} else {
		forum := &model.Forum{Fid: fid}
		err := repository.ForumRepository.First(forum, nil)
		if err != nil {
			return nil, nil, 0, err
		}
		forumModel = forum
	}
	list, totalPage, err := ServiceGroupApp.ThreadService.List(fid, page, pageSize, order, sort)
	if err != nil {
		return nil, nil, 0, err
	}
	var threadVoList []*response.ThreadVo
	for _, v := range list {
		threadVo := transform.TransformThread(v)
		threadVo.User = transform.TransformUser(&v.User)
		group, _ := ServiceGroupApp.GroupService.Detail(threadVo.User.Gid)
		threadVo.User.Group = group
		threadVoList = append(threadVoList, threadVo)
	}
	return forumModel, threadVoList, totalPage, err
}

func (serv *forumService) List() ([]*model.Forum, error) {
	repository.ForumRepository.Pager = &repository.Pager{Page: 0, PageSize: 0}
	list, err := repository.ForumRepository.GetDataListByWhereMap(nil, nil)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (serv *forumService) name() {

}
