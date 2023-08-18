package requests

type ThreadTopRequest struct {
	Fid int ` json:"fid"`
	Tid int ` json:"tid"`
	Top int ` json:"top"`
}
