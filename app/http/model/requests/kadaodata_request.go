package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
)

type KadaoDataRequest struct {
	Username string ` json:"username,omitempty"`
	Title    string ` json:"title,omitempty"`
	Data     string ` json:"data,omitempty"`
}

func (param KadaoDataRequest) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Username,
			validation.Required.Error("用户名必填！"),
			validation.Length(1, 128).Error("用户名超出规定长度"),
		),
		validation.Field(&param.Data,
			validation.Required.Error("缺少卡刀方案数据！"),
			validation.Length(4, 10000).Error("卡刀方案数据超出规定长度"),
		),
		validation.Field(&param.Title,
			validation.Required.Error("卡刀方案标题必填！"),
			validation.Length(4, 128).Error("卡刀方案标题超出规定长度"),
		),
	)
}

type GetKaDaoDataRequest struct {
	Keyword  string ` json:"keyword,omitempty"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Order    string `json:"order"`
	Sort     string `json:"sort"`
	Dpi      string `json:"dpi"`
}

func (param GetKaDaoDataRequest) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Keyword,
			validation.Length(0, 128).Error("关键词超出规定长度"),
		),
		validation.Field(&param.Page,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			//validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			//validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.PageSize,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			//validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			//validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Order,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.In("create_time", "kid", "uid").Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Dpi,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Sort,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.In("desc", "asc").Error(exceptions.ParamInvalid.Error()),
		),
	)
}
