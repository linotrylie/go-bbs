package requests

type MigrationsRequest struct {
	Id        int    ` json:"id"`
	Migration string ` json:"migration"`
	Batch     int    ` json:"batch"`
}
