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
var forumRepository = repository.ForumRepository

func newForumService() *forumService {
	return new(forumService)
}
func (serv *forumService) ThreadList(fid, page, pageSize int, order, sort string) (map[string]interface{}, error) {
	forumModel := &model.Forum{}
	if fid <= 0 {
		forumModel.Fid = 0
		forumModel.Name = "全部"
	} else {
		forum := &model.Forum{Fid: fid}
		err := repository.ForumRepository.First(forum)
		if err != nil {
			return nil, err
		}
		forumModel = forum
	}
	list, totalPage, err := ServiceGroupApp.ThreadService.List(fid, page, pageSize, order, sort)
	if err != nil {
		return nil, err
	}
	var threadVoList []*response.ThreadVo
	for _, v := range list {
		threadVo := transform.TransformThread(v)
		threadVo.User = transform.TransformUser(&v.User)
		group, _ := ServiceGroupApp.GroupService.Detail(threadVo.User.Gid)
		threadVo.User.Group = *group
		threadVoList = append(threadVoList, threadVo)
	}
	mapRes := make(map[string]interface{})
	mapRes["forum"] = forumModel
	mapRes["thread"] = map[string]interface{}{
		"list":       threadVoList,
		"page":       page,
		"page_size":  pageSize,
		"total_page": totalPage,
	}
	return mapRes, err
}

func (serv *forumService) List() ([]*model.Forum, error) {
	repository.ForumRepository.Pager = &repository.Pager{Page: 0, PageSize: 0}
	list, err := repository.ForumRepository.GetDataListByWhereMap(nil)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (serv *forumService) name() {

}
