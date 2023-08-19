package requests

type AttachRequest struct {
	Aid         int    ` json:"aid"`
	Tid         int    ` json:"tid"`
	Pid         int    ` json:"pid"`
	Uid         int    ` json:"uid"`
	Filesize    int    ` json:"filesize"`
	Width       int    ` json:"width"`
	Height      int    ` json:"height"`
	Filename    string ` json:"filename"`
	Orgfilename string ` json:"orgfilename"`
	Filetype    string ` json:"filetype"`
	CreateDate  int    ` json:"createdate"`
	Comment     string ` json:"comment"`
	Downloads   int    ` json:"downloads"`
	Credits     int    ` json:"credits"`
	Golds       int    ` json:"golds"`
	Rmbs        int    ` json:"rmbs"`
	Isimage     int    ` json:"isimage"`
}
