package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
)

type UserController struct {
}

func (controller *UserController) name(ctx *gin.Context) {

}
func (controller *UserController) Index(ctx *gin.Context) {
	result := map[string]interface{}{}
	user := model.User{Uid: 3}
	global.DB.Table("bbs_user").Where("uid = ?", 1).Take(&result)
	global.DB.First(&user)
	fmt.Println(user)
	response.OkWithDetailed(user, "success", ctx)
}
