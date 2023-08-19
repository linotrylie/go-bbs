package requests

type TableDayRequest struct {
	Year       int    ` json:"year"`       // 年
	Month      int    ` json:"month"`      // 月
	Day        int    ` json:"day"`        // 日
	CreateDate int    ` json:"createdate"` // 时间戳
	Table      string ` json:"table"`      // 表名
	Maxid      int    ` json:"maxid"`      // 最大ID
	Count      int    ` json:"count"`      // 总数
}
