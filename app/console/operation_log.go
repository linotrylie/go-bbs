package console

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bbs/app/constants"
	"go-bbs/app/http/model"
	"go-bbs/app/service"
	"go-bbs/global"
	"go.uber.org/zap"
)

func InsertOperationLog() {
	fmt.Println("操作日志消费队列开始！")
	defer fmt.Println("操作日志消费队列完成！")
	go func() {
		for {
			val, err := lpop()
			if err != nil {
				return
			}
			var operationLog model.OperationLog
			err = json.Unmarshal([]byte(val), &operationLog)
			if err != nil {
				return
			}
			err = service.OperationLogService.CreateOperationLog(&operationLog)
		}
	}()
	return
}

func lpop() (string, error) {
	// 获取队列数据
	values, err := global.REDIS.LPop(context.Background(), constants.OperationLog).Result()
	if err != nil {
		global.LOG.Error(err.Error(), zap.Error(err))
		return "", err
	}
	return values, nil
}
