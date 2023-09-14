package respository

import (
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"time"
)

type UserRepository struct {
	User  *model.User
	Pager *Pager
}

// Insert 保存
func (obj *UserRepository) Insert() (effectedRow int64, err error) {
	effectedRow, err = Insert(obj.User)
	if err != nil {
		return
	}
	return
}

// Update 更新
func (obj *UserRepository) Update() (effectedRow int64, err error) {
	effectedRow, err = Update(obj.User)
	if err != nil {
		return
	}
	return
}

// First 查询单条
func (obj *UserRepository) First() (err error) {
	err = FindByLocation(obj.User)
	if err != nil {
		return
	}
	return
}

// Delete 此方法为硬删除 慎用
func (obj *UserRepository) Delete() (rowsAffected int64, e error) {
	rowsAffected, e = DeleteByLocation(obj.User)
	return
}

// FindByWhere 批量查询 带分页
func (obj *UserRepository) FindByWhere(query string, args []interface{}) (list []model.User, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prometheus.OrmWithLabelValues(obj.User.TableName(), "DeleteByLocation", e, now)
		}
	}()
	db := global.DB.Table(obj.User.TableName())
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

func (obj *UserRepository) FindUserByMap(where map[string]interface{}) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	db := global.DB.Table(obj.User.TableName()).Where(where).Find(obj.User)
	SaveInRedis(obj.User)
	e = db.Error
	return
}
