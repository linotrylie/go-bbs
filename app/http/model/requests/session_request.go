package requests

type SessionRequest struct {
	Sid       string ` json:"sid"`
	Uid       int    ` json:"uid"`
	Fid       int    ` json:"fid"`
	Url       string ` json:"url"`
	Ip        int    ` json:"ip"`
	Useragent string ` json:"useragent"`
	Data      string ` json:"data"`
	Bigdata   int    ` json:"bigdata"`
	LastDate  int    ` json:"lastdate"`
}
