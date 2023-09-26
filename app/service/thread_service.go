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
func (serv *threadService) List(fid, page, pageSize int, order, sort string) (threadList []*model.Thread, totalPage int64, e error) {
	threadRepo.Pager = &repository.Pager{Page: page, PageSize: pageSize, FieldsOrder: []string{order + " " + sort}}
	var query string
	if fid < 1 {
		query = "fid > ?"
	} else {
		query = "fid = ?"
	}
	threadList, e = threadRepo.GetDataListByWhere(query, []interface{}{fid}, []string{"User"})
	if e != nil {
		return nil, 0, e
	}
	return threadList, threadRepo.Pager.TotalPage, nil
}
func (serv *threadService) Detail(fid, tid int) (*response.ThreadVo, *response.PostVo, error) {
	thread := &model.Thread{Fid: fid, Tid: tid}
	err := repository.ThreadRepository.First(thread, nil)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}
	postVo := PostService.GetPostTransform(post)
	go serv.After(thread)
	return threadVo, postVo, nil
}

func (serv *threadService) After(thread *model.Thread) {
	thread.SetViews(1)
	repository.ThreadRepository.Update(thread)
}

func (serv *threadService) name() {

}
