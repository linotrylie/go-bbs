package requests

type ChatgptRequest struct {
	Id          int     ` json:"id"`
	Name        string  ` json:"name"`
	ModelType   string  ` json:"modeltype"`
	From        string  ` json:"from"` // 来源
	Golds       int     ` json:"golds"`
	CreatedAt   string  ` json:"createdat"`
	UpdatedAt   string  ` json:"updatedat"`
	MaxTokens   int     ` json:"maxtokens"`
	Temperature float64 ` json:"temperature"`
	DeletedAt   string  ` json:"deletedat"`
}
