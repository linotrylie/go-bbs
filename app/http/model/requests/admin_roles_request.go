package requests

type AdminRolesRequest struct {
	Id        int    ` json:"id"`
	Name      string ` json:"name"`
	Slug      string ` json:"slug"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
