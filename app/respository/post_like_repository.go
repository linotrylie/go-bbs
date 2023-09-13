package respository

import (
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"sync"
)

type PostLikeRepository struct {
	mu       sync.Mutex
	PostLike *model.PostLike
	Pager    *Pager
	IsLock   bool
}

// Insert 保存
func (obj *PostLikeRepository) Insert() (effectedRow int64, err error) {
	effectedRow, err = Insert(obj.PostLike)
	if err != nil {
		return
	}
	return
}

// Update 更新
func (obj *PostLikeRepository) Update() (effectedRow int64, err error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	effectedRow, err = Update(obj.PostLike)
	if err != nil {
		return
	}
	return
}

// First 查询单条
func (obj *PostLikeRepository) First() (err error) {
	err = FindByLocation(obj.PostLike)
	if err != nil {
		return
	}
	return
}

// Delete 此方法为硬删除 慎用
func (obj *PostLikeRepository) Delete() (rowsAffected int64, e error) {
	if obj.IsLock {
		obj.mu.Lock()
		defer obj.mu.Unlock()
	}
	rowsAffected, e = DeleteByLocation(obj.PostLike)
	return
}

// FindByWhere 批量查询 带分页
func (obj *PostLikeRepository) FindByWhere(query string, args []interface{}) (list []model.PostLike, e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	db := global.DB.Table(obj.PostLike.TableName())
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
