package respository

import (
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"sync"
)

type IqismartActivityOrderRepository struct {
	mu                    sync.Mutex
	IqismartActivityOrder *model.IqismartActivityOrder
	Pager                 *Pager
	IsLock                bool
}

// Insert 保存
func (obj *IqismartActivityOrderRepository) Insert() (effectedRow int64, err error) {
	effectedRow, err = Insert(obj.IqismartActivityOrder)
	if err != nil {
		return
	}
	return
}

// Update 更新
func (obj *IqismartActivityOrderRepository) Update() (effectedRow int64, err error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	effectedRow, err = Update(obj.IqismartActivityOrder)
	if err != nil {
		return
	}
	return
}

// First 查询单条
func (obj *IqismartActivityOrderRepository) First() (err error) {
	err = FindByLocation(obj.IqismartActivityOrder)
	if err != nil {
		return
	}
	return
}

// Delete 此方法为硬删除 慎用
func (obj *IqismartActivityOrderRepository) Delete() (rowsAffected int64, e error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	rowsAffected, e = DeleteByLocation(obj.IqismartActivityOrder)
	return
}

// FindByWhere 批量查询 带分页
func (obj *IqismartActivityOrderRepository) FindByWhere(query string, args []interface{}) (list []model.IqismartActivityOrder, e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	db := global.DB.Table(obj.IqismartActivityOrder.TableName())
	if query != "" {
		db = db.Where(query, args...)
	}
	e = obj.Pager.Execute(db, &list)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, exceptions.NotFoundData
	}
	return
}
