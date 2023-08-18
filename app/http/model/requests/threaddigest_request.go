package requests

type ThreadDigestRequest struct {
	Fid    int ` json:"fid"`
	Tid    int ` json:"tid"`
	Uid    int ` json:"uid"`
	Digest int ` json:"digest"`
}
