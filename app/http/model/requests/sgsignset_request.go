package requests

type SgSignSetRequest struct {
	Id        int    ` json:"id"`        // id
	SgSignnum int    ` json:"sgsignnum"` // 签到总数
	SgSign    int    ` json:"sgsign"`    // 今日签到人数
	SgSignOne string ` json:"sgsignone"` // 今日第一
	SgSignTop string ` json:"sgsigntop"` // 今日前十
	Time      int    ` json:"time"`      // 最后签到时间
}
