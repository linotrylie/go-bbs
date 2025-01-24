package requests

type HayaFavoriteRequest struct {
	Tid        int ` json:"tid"`        // 帖子ID
	Uid        int ` json:"uid"`        // 用户ID
	CreateDate int ` json:"createdate"` // 添加时间
	CreateIp   int ` json:"createip"`   // 添加IP
}
