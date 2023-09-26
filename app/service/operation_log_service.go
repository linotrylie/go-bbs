package service

import (
	"go-bbs/app/http/model"
)

type operationLogService struct {
}

var OperationLogService = newOperationLogService()

func newOperationLogService() *operationLogService {
	return new(operationLogService)
}

func (serv *operationLogService) CreateOperationLog(log *model.OperationLog) error {
	_, err := operationLogRepo.Insert(log)
	if err != nil {
		return err
	}
	return nil
}
