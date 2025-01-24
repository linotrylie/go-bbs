package requests

type AdminPermissionMenuRequest struct {
	PermissionId int    ` json:"permissionid"`
	MenuId       int    ` json:"menuid"`
	CreatedAt    string ` json:"createdat"`
	UpdatedAt    string ` json:"updatedat"`
}
