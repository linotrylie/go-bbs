package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/app/repository"
	"go-bbs/app/transform"
)

type postService struct {
}

var PostService = newPostService()

func newPostService() *postService {
	return new(postService)
}

func (serv *postService) CommentList(tid, page, pageSize int, order, sort string) ([]*response.PostVo, int64, error) {
	postRepo.Pager = &repository.Pager{
		Page:        page,
		PageSize:    pageSize,
		FieldsOrder: []string{order + " " + sort},
	}
	commentList, err := postRepo.
		GetDataListByWhere("tid = ? and isfirst = ? and deleted = ?", []interface{}{tid, 0, 0}, []string{"CreateUser", "LastUpdateUser"})
	if err != nil {
		return nil, 0, err
	}
	var postVoList []*response.PostVo
	for _, v := range commentList {
		serv.GetPostTransform(v)
		postVoList = append(postVoList, serv.GetPostTransform(v))
	}
	return postVoList, postRepo.Pager.TotalPage, nil
}

func (serv *postService) GetPostTransform(post *model.Post) *response.PostVo {
	var group *model.Group
	pv := transform.TransformPost(post)
	if post.LastUpdateUser != nil {
		pv.LastUpdateUser = transform.TransformUser(post.LastUpdateUser)
		group, _ = GroupService.Detail(pv.LastUpdateUser.Gid)
		if group != nil {
			pv.LastUpdateUser.Group = group
		}
	}
	if post.CreateUser != nil {
		pv.User = transform.TransformUser(post.CreateUser)
		group, _ = GroupService.Detail(pv.User.Gid)
		if group != nil {
			pv.User.Group = group
		}
	}
	return pv
}

func (serv *postService) name() {

}
