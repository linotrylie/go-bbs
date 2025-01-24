package requests

type UserRaffleLogRequest struct {
	Id         int ` json:"id"`
	Uid        int ` json:"uid"`
	Tid        int ` json:"tid"`
	CreateDate int ` json:"createdate"` // 创建时间
}
