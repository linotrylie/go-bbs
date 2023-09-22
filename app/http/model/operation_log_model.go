package model

import (
	"fmt"
)

type OperationLog struct {
	changes      map[string]interface{}
	Id           int    `gorm:"primaryKey;column:id" json:"id"`
	Ip           uint32 `gorm:"column:ip" json:"ip"`
	Method       string `gorm:"column:method" json:"method"`   // 请求方法
	Path         string `gorm:"column:path" json:"path"`       // 请求路径
	Status       int    `gorm:"column:status" json:"status"`   // 请求状态
	Latency      string `gorm:"column:latency" json:"latency"` // 延迟
	Agent        string `gorm:"column:agent" json:"agent"`
	ErrorMessage string `gorm:"column:error_message" json:"error_message"`
	Body         string `gorm:"column:body" json:"body"`
	Resp         string `gorm:"column:resp" json:"resp"`
	Uid          int    `gorm:"column:uid" json:"uid"`
}

func (*OperationLog) TableName() string {
	return "bbs_operation_log"
}

// Location .
func (obj *OperationLog) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *OperationLog) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
}

// GetChanges .
func (obj *OperationLog) GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *OperationLog) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *OperationLog) SetId(val int) *OperationLog {
	obj.Id += val
	obj.Update("id", obj.Id)
	return obj
}
func (obj *OperationLog) SetIp(val uint32) *OperationLog {
	obj.Ip += val
	obj.Update("ip", obj.Ip)
	return obj
}
func (obj *OperationLog) SetMethod(val string) *OperationLog {
	obj.Method = val
	obj.Update("method", obj.Method)
	return obj
}
func (obj *OperationLog) SetPath(val string) *OperationLog {
	obj.Path = val
	obj.Update("path", obj.Path)
	return obj
}
func (obj *OperationLog) SetStatus(val int) *OperationLog {
	obj.Status += val
	obj.Update("status", obj.Status)
	return obj
}
func (obj *OperationLog) SetLatency(val string) *OperationLog {
	obj.Latency = val
	obj.Update("latency", obj.Latency)
	return obj
}
func (obj *OperationLog) SetAgent(val string) *OperationLog {
	obj.Agent = val
	obj.Update("agent", obj.Agent)
	return obj
}
func (obj *OperationLog) SetErrorMessage(val string) *OperationLog {
	obj.ErrorMessage = val
	obj.Update("error_message", obj.ErrorMessage)
	return obj
}
func (obj *OperationLog) SetBody(val string) *OperationLog {
	obj.Body = val
	obj.Update("body", obj.Body)
	return obj
}
func (obj *OperationLog) SetResp(val string) *OperationLog {
	obj.Resp = val
	obj.Update("resp", obj.Resp)
	return obj
}
func (obj *OperationLog) SetUid(val int) *OperationLog {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
