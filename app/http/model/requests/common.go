package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
	"math"
)

type Pager struct {
	Page     int    `form:"page" binding:"page"`
	PageSize int    `form:"page_size" binding:"page_size"`
	Order    string `form:"order" binding:"order"`
	Sort     string `form:"sort" binding:"sort"`
}

func (param *Pager) Validate() error {
	return validation.ValidateStruct(param,
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

type GenerateFormToken struct {
	Method   string `json:"method"`
	Username int    `json:"username"`
}
