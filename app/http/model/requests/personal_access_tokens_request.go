package requests

type PersonalAccessTokensRequest struct {
	Id            int    ` json:"id"`
	TokenableType string ` json:"tokenabletype"`
	TokenableId   int    ` json:"tokenableid"`
	Name          string ` json:"name"`
	Token         string ` json:"token"`
	Abilities     string ` json:"abilities"`
	LastUsedAt    string ` json:"lastusedat"`
	ExpiresAt     string ` json:"expiresat"`
	CreatedAt     string ` json:"createdat"`
	UpdatedAt     string ` json:"updatedat"`
}
