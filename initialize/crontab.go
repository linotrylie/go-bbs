package initialize

import (
	"fmt"
	"go-bbs/app/console"
)

func InitCrontab() {
	fmt.Println("crontab init")
	console.InitCrontab()
}
