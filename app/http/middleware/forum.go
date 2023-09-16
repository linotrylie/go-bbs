package middleware

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/app/respository"
	"go-bbs/global"
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
		forumId, ok := ctx.GetQuery("forum")
		if !ok {
			response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
			ctx.Abort()
			return
		}
		fid, _ := strconv.Atoi(forumId)
		if fid <= 0 {
			ctx.Next()
			return
		}
		forumAccess := &model.ForumAccess{Gid: gid, Fid: fid}
		err := respository.FindByLocation(forumAccess)
		if err != nil {
			response.FailWithMessage(err.Error(), ctx)
			ctx.Abort()
			return
		}
		if forumAccess.Allowread == 0 {
			response.FailWithMessage(exceptions.NotAuth.Error(), ctx)
			ctx.Abort()
			return
		}
		ctx.Set("forum_access", forumAccess)
		ctx.Next()
	}
}
