package requests

type ChatgptDataRequest struct {
	Id           int    ` json:"id"`
	Uid          int    ` json:"uid"`
	Prompt       string ` json:"prompt"` // 提问问题
	Answer       string ` json:"answer"` // 回答记录
	RequestData  ` json:"requestdata"`
	ResponseData ` json:"responsedata"`
	Ip           string ` json:"ip"`
	ModelId      int    ` json:"modelid"`
	ModelName    string ` json:"modelname"`
	CreatedAt    string ` json:"createdat"`
	UpdatedAt    string ` json:"updatedat"`
	Golds        int    ` json:"golds"`
}
