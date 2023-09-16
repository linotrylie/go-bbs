package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/app/respository"
	"go-bbs/app/transform"
)

type ForumService struct {
	ForumRepo respository.ForumRepository
}

func (serv *ForumService) ThreadList(fid, page, pageSize int) (map[string]interface{}, error) {
	forumModel := &model.Forum{}
	if fid <= 0 {
		forumModel.Fid = 0
		forumModel.Name = "全部"
	} else {
		serv.ForumRepo.Forum = &model.Forum{Fid: fid}
		err := serv.ForumRepo.First()
		if err != nil {
			return nil, err
		}
		forumModel = serv.ForumRepo.Forum
	}
	ServiceGroupApp.ThreadService.ForumRepo.Pager = &respository.Pager{Page: page, PageSize: pageSize, FieldsOrder: []string{"create_date desc"}}
	list, err := ServiceGroupApp.ThreadService.List(fid)
	if err != nil {
		return nil, err
	}
	var threadVoList []response.ThreadVo
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
		"total_page": ServiceGroupApp.ThreadService.ForumRepo.Pager.TotalPage,
	}
	return mapRes, err
}

func (serv *ForumService) List() ([]model.Forum, error) {
	serv.ForumRepo.Pager = &respository.Pager{Page: 1, PageSize: 100}
	list, err := serv.ForumRepo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (serv *ForumService) name() {

}
