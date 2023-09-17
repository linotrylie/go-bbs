package response

type PostVo struct {
	Tid              int         `json:"tid"`
	Pid              int         `json:"pid"`
	Uid              int         `json:"uid"`
	Isfirst          int         `json:"isfirst"`
	CreateDate       string      `json:"create_date"`
	Userip           string      `json:"userip"`
	Images           int         `json:"images"`
	Files            int         `json:"files"`
	Doctype          int         `json:"doctype"`
	Quotepid         int         `json:"quotepid"`
	Message          string      `json:"message"`
	MessageFmt       string      `json:"message_fmt"`
	LocationPost     string      `json:"location_post"`
	Likes            int         `json:"likes"` // 点赞数
	Deleted          int         `json:"deleted"`
	Updates          int         `json:"updates"`
	LastUpdateDate   string      `json:"last_update_date"`
	LastUpdateUid    int         `json:"last_update_uid"`
	LastUpdateReason string      `json:"last_update_reason"`
	ReplyHide        int         `json:"reply_hide"`
	LastUpdateUser   *UserVo     `json:"last_update_user"`
	User             *UserVo     `json:"user"`
	AttachList       []*AttachVo `json:"attach_list"`
}
