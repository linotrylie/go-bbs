package service

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/app/repository"
	"go-bbs/app/transform"
	"go-bbs/utils"
	"strconv"
	"time"
)

type postService struct {
}

var PostService = newPostService()

func newPostService() *postService {
	return new(postService)
}

func (serv *postService) CommentList(tid, page, pageSize int, order, sort string) ([]*response.PostVo, int64, error) {
	repository.PostRepository.Pager = &repository.Pager{
		Page:        page,
		PageSize:    pageSize,
		FieldsOrder: []string{order + " " + sort},
	}
	commentList, err := repository.PostRepository.
		GetDataListByWhere("tid = ? and isfirst = ? and deleted = ?", []interface{}{tid, 0, 0}, []string{"CreateUser", "LastUpdateUser"})
	if err != nil {
		return nil, 0, err
	}
	var postVoList []*response.PostVo
	for _, v := range commentList {
		serv.GetPostTransform(v)
		postVoList = append(postVoList, serv.GetPostTransform(v))
	}
	return postVoList, repository.PostRepository.Pager.TotalPage, nil
}

// CommentCreate 创建评论
func (serv *postService) CommentCreate(commentCreate requests.PostCommentCreate, ctx *gin.Context) error {
	var post *model.Post
	//先检查评论的帖子是否还存在
	isExist := ThreadService.ExistThread(commentCreate.Tid)
	if !isExist {
		return exceptions.ThreadIsValid
	}
	quote, err := serv.PostQuote(post.Quotepid)
	if err != nil {
		return nil
	}
	//<blockquote class="blockquote">
	//		<a href="user-14.htm" class="text-small text-muted user">
	//			<img class="avatar-1" src="view/img/avatar.png">
	//			xiaohui
	//		</a>
	//		天啊  两个月 35 ！？
	//		</blockquote>上去打2个怪就下，不想跑，完全玩不起来呢
	post.SetUid(commentCreate.Uid).SetMessage(utils.TrimHtml(commentCreate.Message)).
		SetMessageFmt(quote + post.Message).SetCreateDate(time.Now().Unix()).
		SetIsfirst(0).SetUserip(utils.Ip2long(ctx.ClientIP())).
		SetDoctype(0)
	return nil
}

func (serv *postService) PostQuote(pid int) (string, error) {
	var comment *model.Post
	comment.Pid = pid
	err := repository.PostRepository.First(comment, nil)
	if err != nil {
		return "", err
	}
	if comment.Message == "" {
		return "", exceptions.NotFoundData
	}
	var user *model.User
	user.Uid = comment.Uid
	err = repository.UserRepository.First(user, nil)
	if err != nil {
		return "", err
	}
	if user.Username == "" {
		return "", exceptions.NotFoundData
	}
	userVo := transform.TransformUser(user)
	r := `<blockquote class="blockquote">
	<a href="user-` + strconv.Itoa(userVo.Uid) + `" class="text-small text-muted user">
	<img class="avatar-1" src="` + userVo.Avatar + `">
		` + userVo.Username + `
	</a>
	` + comment.Message + `
	</blockquote>`
	return r, err
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
