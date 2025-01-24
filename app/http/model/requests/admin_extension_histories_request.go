package requests

type AdminExtensionHistoriesRequest struct {
	Id        int    ` json:"id"`
	Name      string ` json:"name"`
	Type      int    ` json:"type"`
	Version   string ` json:"version"`
	Detail    string ` json:"detail"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
