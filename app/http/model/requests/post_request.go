package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
	"math"
)

type PostRequest struct {
	Tid              int    ` json:"tid"`
	Pid              int    ` json:"pid"`
	Uid              int    ` json:"uid"`
	Isfirst          int    ` json:"isfirst"`
	CreateDate       int    ` json:"createdate"`
	Userip           int    ` json:"userip"`
	Images           int    ` json:"images"`
	Files            int    ` json:"files"`
	Doctype          int    ` json:"doctype"`
	Quotepid         int    ` json:"quotepid"`
	Message          string ` json:"message"`
	MessageFmt       string ` json:"messagefmt"`
	LocationPost     string ` json:"locationpost"`
	Likes            int    ` json:"likes"` // 点赞数
	Deleted          int    ` json:"deleted"`
	Updates          int    ` json:"updates"`
	LastUpdateDate   int    ` json:"lastupdatedate"`
	LastUpdateUid    int    ` json:"lastupdateuid"`
	LastUpdateReason string ` json:"lastupdatereason"`
}

type PostList struct {
	Tid int `form:"tid" binding:"tid"`
	Pid int `form:"pid" binding:"pid"`
	Pager
}

func (param *PostList) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Tid,
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Pid,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Pager),
	)
}
