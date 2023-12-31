package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
	"math"
)

type ThreadRequest struct {
	Fid            int    ` json:"fid"`
	Tid            int    ` json:"tid"`
	Top            int    ` json:"top"`
	Uid            int    ` json:"uid"`
	Userip         int    ` json:"userip"`
	Subject        string ` json:"subject"`
	CreateDate     int    ` json:"createdate"`
	LastDate       int    ` json:"lastdate"`
	Views          int    ` json:"views"`
	Posts          int    ` json:"posts"`
	Images         int    ` json:"images"`
	Files          int    ` json:"files"`
	Mods           int    ` json:"mods"`
	Closed         int    ` json:"closed"`
	Firstpid       int    ` json:"firstpid"`
	Lastuid        int    ` json:"lastuid"`
	Lastpid        int    ` json:"lastpid"`
	LocationTr     string ` json:"locationtr"`
	Favorites      int    ` json:"favorites"` // 收藏数
	Likes          int    ` json:"likes"`     // 点赞数
	Highlight      int    ` json:"highlight"`
	ContentBuy     int    ` json:"contentbuy"`
	ContentBuyType int    ` json:"contentbuytype"`
	Digest         int    ` json:"digest"`
	Deleted        int    ` json:"deleted"`
	Readp          int    ` json:"readp"`
	OfferNum       int    ` json:"offernum"`
	OfferStatus    int    ` json:"offerstatus"`
	Tagids         string ` json:"tagids"`
	TagidsTime     int    ` json:"tagidstime"`
}

type ThreadList struct {
	Fid int `form:"fid"`
	Tid int `form:"tid"`
}

func (param *ThreadList) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Tid,
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
		validation.Field(&param.Fid,
			validation.Required.Error(exceptions.ParamInvalid.Error()),
			validation.Min(1).Error(exceptions.ParamInvalid.Error()),
			validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
		),
	)
}
