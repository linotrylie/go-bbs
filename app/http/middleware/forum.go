package middleware

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	repository "go-bbs/app/repository"
	"go-bbs/global"
	"math"
	"strconv"
)

func AuthForum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var gid int
		if global.User == nil {
			gid = 0
		} else {
			gid = global.User.Gid
		}
		forumId := ctx.DefaultQuery("fid", "0")
		fid, _ := strconv.Atoi(forumId)
		err := validation.Validate(fid,
			validation.Min(0).Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		)
		if err != nil {
			response.FailWithMessage(err.Error(), ctx)
			ctx.Abort()
			return
		}
		if fid > 0 {
			forumAccess := &model.ForumAccess{Gid: gid, Fid: fid}
			repository.ForumAccessRepository.First(forumAccess, nil)
			if forumAccess.Fid != 0 && forumAccess.Allowread == 0 {
				response.FailWithMessage(exceptions.NotAuth.Error(), ctx)
				ctx.Abort()
				return
			}
			ctx.Set("forum_access", forumAccess)
		}
		ctx.Next()
	}
}
