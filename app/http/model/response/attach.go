package response

type AttachVo struct {
	Aid         int    `json:"aid"`
	Tid         int    `json:"tid"`
	Pid         int    `json:"pid"`
	Uid         int    `json:"uid"`
	Filesize    int64  `json:"filesize"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Filename    string `json:"filename"`
	Orgfilename string `json:"orgfilename"`
	Filetype    string `json:"filetype"`
	CreateDate  string `json:"create_date"`
	Comment     string `json:"comment"`
	Downloads   int    `json:"downloads"`
	Credits     int    `json:"credits"`
	Golds       int    `json:"golds"`
	Rmbs        int    `json:"rmbs"`
	Isimage     int    `json:"isimage"`
}
