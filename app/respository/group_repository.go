package respository

import (
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"sync"
)

type GroupRepository struct {
	mu     sync.Mutex
	Group  *model.Group
	Pager  *Pager
	IsLock bool
}

// Insert 保存
func (obj *GroupRepository) Insert() (effectedRow int64, err error) {
	effectedRow, err = Insert(obj.Group)
	if err != nil {
		return
	}
	return
}

// Update 更新
func (obj *GroupRepository) Update() (effectedRow int64, err error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	effectedRow, err = Update(obj.Group)
	if err != nil {
		return
	}
	return
}

// First 查询单条
func (obj *GroupRepository) First() (err error) {
	err = FindByLocation(obj.Group)
	if err != nil {
		return
	}
	return
}

// Delete 此方法为硬删除 慎用
func (obj *GroupRepository) Delete() (rowsAffected int64, e error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	rowsAffected, e = DeleteByLocation(obj.Group)
	return
}

// FindByWhere 批量查询 带分页
func (obj *GroupRepository) FindByWhere(query string, args []interface{}) (list []model.Group, e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	db := global.DB.Table(obj.Group.TableName())
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
