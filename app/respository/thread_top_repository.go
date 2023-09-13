package respository

import (
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"sync"
)

type ThreadTopRepository struct {
	mu        sync.Mutex
	ThreadTop *model.ThreadTop
	Pager     *Pager
	IsLock    bool
}

// Insert 保存
func (obj *ThreadTopRepository) Insert() (effectedRow int64, err error) {
	effectedRow, err = Insert(obj.ThreadTop)
	if err != nil {
		return
	}
	return
}

// Update 更新
func (obj *ThreadTopRepository) Update() (effectedRow int64, err error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	effectedRow, err = Update(obj.ThreadTop)
	if err != nil {
		return
	}
	return
}

// First 查询单条
func (obj *ThreadTopRepository) First() (err error) {
	err = FindByLocation(obj.ThreadTop)
	if err != nil {
		return
	}
	return
}

// Delete 此方法为硬删除 慎用
func (obj *ThreadTopRepository) Delete() (rowsAffected int64, e error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	rowsAffected, e = DeleteByLocation(obj.ThreadTop)
	return
}

// FindByWhere 批量查询 带分页
func (obj *ThreadTopRepository) FindByWhere(query string, args []interface{}) (list []model.ThreadTop, e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	db := global.DB.Table(obj.ThreadTop.TableName())
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
