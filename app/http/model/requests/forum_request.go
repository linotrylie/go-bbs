package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
	"math"
)

type ForumRequest struct {
	Fid           int    ` json:"fid"`
	Name          string ` json:"name"`
	Rank          int    ` json:"rank"`
	Threads       int    ` json:"threads"`
	Todayposts    int    ` json:"todayposts"`
	Todaythreads  int    ` json:"todaythreads"`
	Brief         string ` json:"brief"`
	Announcement  string ` json:"announcement"`
	Accesson      int    ` json:"accesson"`
	Orderby       int    ` json:"orderby"`
	CreateDate    int    ` json:"createdate"`
	Icon          int    ` json:"icon"`
	Moduids       string ` json:"moduids"`
	SeoTitle      string ` json:"seotitle"`
	SeoKeywords   string ` json:"seokeywords"`
	Digests       int    ` json:"digests"`
	CreateCredits int    ` json:"createcredits"`
	CreateGolds   int    ` json:"creategolds"`
	PostCredits   int    ` json:"postcredits"`
	PostGolds     int    ` json:"postgolds"`
	AllowOffer    int    ` json:"allowoffer"`
}
type ThreadList struct {
	Fid      int    `form:"fid"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Order    string `form:"order"`
	Sort     string `form:"sort"`
}

func (param *ThreadList) Validate() error {
	return validation.ValidateStruct(param,
		//validation.Field(&param.Fid,
		//	validation.Min(-1).Error(exceptions.ParamInvalid.Error()),
		//	validation.Required.Error(exceptions.ParamInvalid.Error()),
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
