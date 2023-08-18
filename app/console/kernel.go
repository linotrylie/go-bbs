package console

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func InitCrontab() {
	c := newWithSeconds()
	var err error
	_, err = c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	if err != nil {
		return
	}
	_, err = c.AddFunc("1 * * * * *", Test)
	if err != nil {
		return
	}
	c.Start()
}
