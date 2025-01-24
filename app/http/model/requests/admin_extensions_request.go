package requests

type AdminExtensionsRequest struct {
	Id        int    ` json:"id"`
	Name      string ` json:"name"`
	Version   string ` json:"version"`
	IsEnabled int    ` json:"isenabled"`
	Options   string ` json:"options"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
