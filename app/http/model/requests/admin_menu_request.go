package requests

type AdminMenuRequest struct {
	Id        int    ` json:"id"`
	ParentId  int    ` json:"parentid"`
	Order     int    ` json:"order"`
	Title     string ` json:"title"`
	Icon      string ` json:"icon"`
	Uri       string ` json:"uri"`
	Extension string ` json:"extension"`
	Show      int    ` json:"show"`
	CreatedAt string ` json:"createdat"`
	UpdatedAt string ` json:"updatedat"`
}
