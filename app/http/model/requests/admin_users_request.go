package requests

type AdminUsersRequest struct {
	Id            int    ` json:"id"`
	Username      string ` json:"username"`
	Password      string ` json:"password"`
	Name          string ` json:"name"`
	Avatar        string ` json:"avatar"`
	RememberToken string ` json:"remembertoken"`
	CreatedAt     string ` json:"createdat"`
	UpdatedAt     string ` json:"updatedat"`
}
