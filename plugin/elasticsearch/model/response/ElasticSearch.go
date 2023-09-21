package response

import "go-bbs/app/http/model/requests"

type ElasticSearchSearch struct {
	Title string `json:"title" form:"title"`
	Type  uint   `json:"type" form:"type"`
	requests.Pager
}
