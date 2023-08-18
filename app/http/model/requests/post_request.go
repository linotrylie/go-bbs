package requests

type PostRequest struct {
	Tid              int    ` json:"tid"`
	Pid              int    ` json:"pid"`
	Uid              int    ` json:"uid"`
	Isfirst          int    ` json:"isfirst"`
	CreateDate       int    ` json:"createdate"`
	Userip           int    ` json:"userip"`
	Images           int    ` json:"images"`
	Files            int    ` json:"files"`
	Doctype          int    ` json:"doctype"`
	Quotepid         int    ` json:"quotepid"`
	Message          string ` json:"message"`
	MessageFmt       string ` json:"messagefmt"`
	LocationPost     string ` json:"locationpost"`
	Likes            int    ` json:"likes"` // 点赞数
	Deleted          int    ` json:"deleted"`
	Updates          int    ` json:"updates"`
	LastUpdateDate   int    ` json:"lastupdatedate"`
	LastUpdateUid    int    ` json:"lastupdateuid"`
	LastUpdateReason string ` json:"lastupdatereason"`
}
