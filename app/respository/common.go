package respository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/redis/go-redis/v9"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	model model.Model
}
type Pager struct {
	PageSize    int
	Page        int
	TotalPage   int
	FieldsOrder []string // []{"id desc","name asc"}
}

func Insert(model model.Model) (rowsAffected int64, err error) {
	result := global.DB.Create(model)
	if result.Error != nil {
		global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	SaveInRedis(model)
	return result.RowsAffected, result.Error
}

func Update(model model.Model) (rowsAffected int64, e error) {
	if len(model.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := model.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(model.TableName()).Where(model.Location()).Updates(updateValues)
	if result.Error != nil {
		//global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	//更新完成后，重新缓存
	FindByLocation(model)
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func FindByLocation(model model.Model) (e error) {
	if len(model.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	err := FindInRedis(model)
	if err == nil {
		return err
	}
	result := global.DB.Table(model.TableName()).Where(model.Location()).First(model)
	if result.Error != nil {
		global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	SaveInRedis(model)
	e = result.Error
	return
}

// DeleteByLocation 此方法为硬删除 慎用
func DeleteByLocation(model model.Model) (rowsAffected int64, e error) {
	if len(model.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(model.TableName()).Where(model.Location()).Unscoped().Delete(model)
	if result.Error != nil {
		global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	DeleteInRedis(model)
	rowsAffected = result.RowsAffected
	e = result.Error
	return
}

func (p *Pager) Execute(db *gorm.DB, object interface{}) (e error) {

	if p.Page != 0 && p.PageSize != 0 {
		var count64 int64
		e = db.Count(&count64).Error
		count := int(count64)
		if e != nil {
			return
		}
		if count != 0 {
			//Calculate the length of the pagination
			if count%p.PageSize == 0 {
				p.TotalPage = count / p.PageSize
			} else {
				p.TotalPage = count/p.PageSize + 1
			}
		}
		db = db.Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize)
	}

	orderValue := p.FieldsOrder
	if len(orderValue) > 0 {
		for _, v := range orderValue {
			db = db.Order(v)
		}
	}

	resultDB := db.Find(object)
	if resultDB.Error != nil {
		return resultDB.Error
	}
	return
}

// 事务
func TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	return global.DB.Transaction(func(tx *gorm.DB) (e error) {
		defer func() {
			if err := recover(); err != nil {
				e = errors.New(fmt.Sprint(err))
				global.LOG.Error(e.Error(), zap.Error(e))
			}
		}()
		e = fun()
		return
	}, opts...)
}

//////////////Redis///////////////////////////

func SaveInRedis(model model.Model) (e error) {
	var redisKey string
	redisKey = model.RedisKey()
	resByte, e := json.Marshal(model)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 999999))*time.Second)
	return nil
}

func FindInRedis(model model.Model) (e error) {
	var redisKey string
	redisKey = model.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil {
		global.LOG.Error(e.Error(), zap.Error(e))
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), model)
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			return
		}
	}
	return
}

func DeleteInRedis(model model.Model) (e error) {
	var redisKey string
	redisKey = model.RedisKey()
	err := global.REDIS.Del(context.Background(), redisKey).Err()
	if err != nil {
		global.LOG.Error(err.Error(), zap.Error(err))
		return
	}
	return nil
}
