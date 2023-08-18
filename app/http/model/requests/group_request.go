package requests

type GroupRequest struct {
	Gid             int    ` json:"gid"`
	Name            string ` json:"name"`
	Creditsfrom     int    ` json:"creditsfrom"`
	Creditsto       int    ` json:"creditsto"`
	Allowread       int    ` json:"allowread"`
	Allowthread     int    ` json:"allowthread"`
	Allowpost       int    ` json:"allowpost"`
	Allowattach     int    ` json:"allowattach"`
	Allowdown       int    ` json:"allowdown"`
	Allowtop        int    ` json:"allowtop"`
	Allowupdate     int    ` json:"allowupdate"`
	Allowdelete     int    ` json:"allowdelete"`
	Allowmove       int    ` json:"allowmove"`
	Allowbanuser    int    ` json:"allowbanuser"`
	Allowdeleteuser int    ` json:"allowdeleteuser"`
	Allowviewip     int    ` json:"allowviewip"`
	Allowharddelete int    ` json:"allowharddelete"`
	Readp           int    ` json:"readp"`
	AllowPostRead   int    ` json:"allowpostread"`
	Allowsell       int    ` json:"allowsell"`
	AllowOffer      int    ` json:"allowoffer"`
}
