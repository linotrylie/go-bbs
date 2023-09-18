package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/redis/go-redis/v9"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type attachRepository struct {
	Pager *Pager
}

var AttachRepository = newAttachRepository()

func newAttachRepository() *attachRepository {
	return new(attachRepository)
}

func (repo *attachRepository) Insert(attach *model.Attach) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(attach)
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(attach)
	return result.RowsAffected, result.Error
}

func (repo *attachRepository) Update(attach *model.Attach) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "Update", e, now)
		}
	}()
	if len(attach.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := attach.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(attach.TableName()).Where(attach.Location()).Updates(updateValues)
	if result.Error != nil {
		return
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(attach)
	repo.First(attach)
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *attachRepository) First(attach *model.Attach) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "First", e, now)
		}
	}()
	if len(attach.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	err := repo.FindInRedis(attach)
	if err != nil && err != redis.Nil {
		return err
	}
	result := global.DB.Table(attach.TableName()).Where(attach.Location()).First(attach)
	e = result.Error
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(attach)
	return
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *attachRepository) DeleteByLocation(attach *model.Attach) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(attach.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(attach.TableName()).Where(attach.Location()).Unscoped().Delete(attach)
	if result.Error != nil {
		return
	}
	repo.DeleteInRedis(attach)
	rowsAffected = result.RowsAffected
	e = result.Error
	return
}

// 事务
func (repo *attachRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *attachRepository) SaveInRedis(attach *model.Attach) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = attach.RedisKey()
	resByte, e := json.Marshal(attach)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *attachRepository) FindInRedis(attach *model.Attach) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = attach.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), attach)
	}
	return
}

func (repo *attachRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	redisRes, e = global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		return
	}
	return
}

func (repo *attachRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *attachRepository) DeleteInRedis(attach *model.Attach) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = attach.RedisKey()
	err := global.REDIS.Del(context.Background(), redisKey).Err()
	if err != nil {
		return
	}
	return nil
}
func (repo *attachRepository) GetDataListByWhereMap(query map[string]interface{}) (list []*model.Attach, e error) {
	now := time.Now()
	attach := &model.Attach{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := attach.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(attach)
		if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
			var count64 int64
			e = db.Count(&count64).Error
			count := int(count64)
			if e != nil {
				return
			}
			if count != 0 {
				//Calculate the length of the pagination
				if count%repo.Pager.PageSize == 0 {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 0
				} else {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
				}
			}
		}
		return
	}
	db := global.DB.Table(attach.TableName()).Where(query)
	e = repo.Execute(db, &list)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, exceptions.NotFoundData
	}
	marshal, e := json.Marshal(list)
	if e != nil {
		return nil, e
	}
	repo.SaveInRedisByKey(redisKey, string(marshal))
	return
}

func (repo *attachRepository) GetDataListByWhere(query string, args []interface{}) (list []*model.Attach, e error) {
	now := time.Now()
	attach := &model.Attach{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := attach.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(attach)
		if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
			var count64 int64
			e = db.Count(&count64).Error
			count := int(count64)
			if e != nil {
				return
			}
			if count != 0 {
				//Calculate the length of the pagination
				if count%repo.Pager.PageSize == 0 {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 0
				} else {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
				}
			}
		}
		return
	}
	db := global.DB.Table(attach.TableName())
	if query != "" {
		db = db.Where(query, args...)
	}
	e = repo.Execute(db, &list)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, exceptions.NotFoundData
	}
	marshal, e := json.Marshal(list)
	if e != nil {
		return nil, e
	}
	repo.SaveInRedisByKey(redisKey, string(marshal))
	return
}

func (repo *attachRepository) GetDataByWhereMap(attach *model.Attach, where map[string]interface{}) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(attach.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Table(attach.TableName()).Where(where).First(attach)
	e = db.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(attach)
	return
}

func (repo *attachRepository) Execute(db *gorm.DB, object interface{}) error {
	if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
		var count64 int64
		e := db.Count(&count64).Error
		count := int(count64)
		if e != nil {
			return e
		}
		if count != 0 {
			//Calculate the length of the pagination
			if count%repo.Pager.PageSize == 0 {
				repo.Pager.TotalPage = count / repo.Pager.PageSize
			} else {
				repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
			}
		}
		db = db.Offset((repo.Pager.Page - 1) * repo.Pager.PageSize).Limit(repo.Pager.PageSize)
	}
	orderValue := repo.Pager.FieldsOrder
	if len(orderValue) > 0 {
		for _, v := range orderValue {
			db = db.Order(v)
		}
	}
	resultDB := db.Find(object)
	if resultDB.Error != nil {
		return resultDB.Error
	}
	return nil
}
