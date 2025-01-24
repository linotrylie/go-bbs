package requests

type FailedJobsRequest struct {
	Id         int    ` json:"id"`
	Uuid       string ` json:"uuid"`
	Connection string ` json:"connection"`
	Queue      string ` json:"queue"`
	Payload    string ` json:"payload"`
	Exception  string ` json:"exception"`
	FailedAt   string ` json:"failedat"`
}
