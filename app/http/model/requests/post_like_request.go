package requests

type PostLikeRequest struct {
	Tid        int ` json:"tid"`        // 帖子ID
	Pid        int ` json:"pid"`        // 回帖ID
	Uid        int ` json:"uid"`        // 用户ID
	CreateDate int ` json:"createdate"` // 添加时间
	CreateIp   int ` json:"createip"`   // 添加IP
}
