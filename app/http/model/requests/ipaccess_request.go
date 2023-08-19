package requests

type IpaccessRequest struct {
	Ip          int ` json:"ip"`
	Mails       int ` json:"mails"`
	Users       int ` json:"users"`
	Logins      int ` json:"logins"`
	Threads     int ` json:"threads"`
	Posts       int ` json:"posts"`
	Attachs     int ` json:"attachs"`
	Attachsizes int ` json:"attachsizes"`
	LastDate    int ` json:"lastdate"`
	Actions     int ` json:"actions"`
	Action1     int ` json:"action1"`
	Action2     int ` json:"action2"`
	Action3     int ` json:"action3"`
	Action4     int ` json:"action4"`
}
