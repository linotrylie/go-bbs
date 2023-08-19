package requests

type ForumAccessRequest struct {
	Fid         int ` json:"fid"`
	Gid         int ` json:"gid"`
	Allowread   int ` json:"allowread"`
	Allowthread int ` json:"allowthread"`
	Allowpost   int ` json:"allowpost"`
	Allowattach int ` json:"allowattach"`
	Allowdown   int ` json:"allowdown"`
}
