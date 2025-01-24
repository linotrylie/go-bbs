package requests

type AdsRequest struct {
	Id        int    ` json:"id"`
	Title     string ` json:"title"`
	Content   string ` json:"content"`
	PicUrl    string ` json:"picurl"`
	Position  int    ` json:"position"` // 0 卡刀页背景图 1卡刀页左边 2 卡刀页右边 3 用户页头图
	Link      string ` json:"link"`
	ExpiredAt string ` json:"expiredat"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
	DeletedAt string ` json:"deletedat"`
}
