package requests

type AdminRoleMenuRequest struct {
	RoleId    int    ` json:"roleid"`
	MenuId    int    ` json:"menuid"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
