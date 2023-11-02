package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total_page"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
