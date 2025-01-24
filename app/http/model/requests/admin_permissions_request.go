package requests

type AdminPermissionsRequest struct {
	Id         int    ` json:"id"`
	Name       string ` json:"name"`
	Slug       string ` json:"slug"`
	HttpMethod string ` json:"httpmethod"`
	HttpPath   string ` json:"httppath"`
	Order      int    ` json:"order"`
	ParentId   int    ` json:"parentid"`
	CreatedAt  string ` json:"createdat"`
	UpdatedAt  string ` json:"updatedat"`
}
