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
	Tid int `form:"tid"`
	//Pid      int    `form:"pid" binding:"pid"`
	Page     int    `form:"page" `
	PageSize int    `form:"page_size"`
	Order    string `form:"order"`
	Sort     string `form:"sort"`
}

func (param *PostList) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Tid,
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		//validation.Field(&param.Pid,
		//	validation.Required.Error(exceptions.ParamInvalid.Error()),
		//	validation.Min(1).Error(exceptions.ParamInvalid.Error()),
		//	validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		//),
		validation.Field(&param.Page,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.PageSize,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Order,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.In("create_date", "last_date", "posts", "views").Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Sort,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.In("desc", "asc").Error(exceptions.ParamInvalid.Error()),
		),
	)
}
