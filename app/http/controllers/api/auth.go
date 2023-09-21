package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bbs/app/entity"
	"go-bbs/app/http/model"
	"go-bbs/app/repository"
	"net/http"
	"strconv"
)

type AuthController struct {
}

func (c *AuthController) Index(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	var user = &model.User{}
	user.Uid = 1
	user.Update("username", "freebns1")
	_, err := repository.UserRepository.Update(user)
	if err != nil {
		return
	}
	var group = &model.Group{}
	group.Gid = user.Gid
	repository.GroupRepository.First(group, nil)
	user.SetCredits(1).SetGolds(1)
	fmt.Println(user)
	userEntity := entity.UserEntity{User: user, Group: group}

	pager := repository.Pager{Page: page, PageSize: 5}
	args := make([]interface{}, 1)
	args[0] = 1
	repository.UserRepository.Pager = &pager
	list, _ := repository.UserRepository.GetDataListByWhere("uid > ?", args, nil)
	var result = make(map[string]interface{})
	result["list"] = list
	result["page"] = pager.Page
	result["page_size"] = pager.PageSize
	result["total_page"] = pager.TotalPage
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
		"user":    userEntity,
		"res":     result,
	})
}

func (controller *AuthController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
	})
}

func (controller *AuthController) Login(ctx *gin.Context) {

}
