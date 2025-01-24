package requests

type SgSignRequest struct {
	Id           int    ` json:"id"`           // ID
	Uid          int    ` json:"uid"`          // 用户ID
	Stime        int    ` json:"stime"`        // 最后签到时间
	Credits      int    ` json:"credits"`      // 签到积分
	Todaycredits int    ` json:"todaycredits"` // 今日积分
	Counts       int    ` json:"counts"`       // 签到天数
	Keepdays     int    ` json:"keepdays"`     // 连续签到
	User         string ` json:"user"`         // 签到用户
}
