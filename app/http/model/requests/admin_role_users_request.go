package requests

type AdminRoleUsersRequest struct {
	RoleId    int    ` json:"roleid"`
	UserId    int    ` json:"userid"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
