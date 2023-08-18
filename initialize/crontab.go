package initialize

import (
	"GoFreeBns/app/console"
	"fmt"
)

func InitCrontab() {
	fmt.Println("crontab init")
	console.InitCrontab()
}
