package requests

type AdminRolePermissionsRequest struct {
	RoleId       int    ` json:"roleid"`
	PermissionId int    ` json:"permissionid"`
	CreatedAt    string ` json:"createdat"`
	UpdatedAt    string ` json:"updatedat"`
}
