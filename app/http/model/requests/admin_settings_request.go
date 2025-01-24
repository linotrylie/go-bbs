package requests

type AdminSettingsRequest struct {
	Slug      string ` json:"slug"`
	Value     string ` json:"value"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
