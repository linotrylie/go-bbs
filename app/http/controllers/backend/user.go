package backend

import (
	"GoFreeBns/app/http/model"
	"GoFreeBns/app/http/model/response"
	"GoFreeBns/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (controller *UserController) Index(ctx *gin.Context) {
	result := map[string]interface{}{}
	user := model.User{Uid: 3}
	global.DB.Table("bbs_user").Where("uid = ?", 1).Take(&result)
	global.DB.First(&user)
	fmt.Println(user)
	response.OkWithDetailed(user, "success", ctx)
}
