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

type postUpdateLogRepository struct {
	Pager *Pager
}

var PostUpdateLogRepository = newPostUpdateLogRepository()

func newPostUpdateLogRepository() *postUpdateLogRepository {
	return new(postUpdateLogRepository)
}

func (repo *postUpdateLogRepository) Insert(postUpdateLog *model.PostUpdateLog) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(postUpdateLog)
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(postUpdateLog)
	return result.RowsAffected, result.Error
}

func (repo *postUpdateLogRepository) Update(postUpdateLog *model.PostUpdateLog) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "Update", e, now)
		}
	}()
	if len(postUpdateLog.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := postUpdateLog.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(postUpdateLog.TableName()).Where(postUpdateLog.Location()).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(postUpdateLog)
	repo.First(postUpdateLog, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *postUpdateLogRepository) First(postUpdateLog *model.PostUpdateLog, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "First", e, now)
		}
	}()
	if len(postUpdateLog.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	e = repo.FindInRedis(postUpdateLog)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(postUpdateLog.TableName()).Where(postUpdateLog.Location())
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db.First(postUpdateLog)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(postUpdateLog)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *postUpdateLogRepository) DeleteByLocation(postUpdateLog *model.PostUpdateLog) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(postUpdateLog.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(postUpdateLog.TableName()).Where(postUpdateLog.Location()).Unscoped().Delete(postUpdateLog)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(postUpdateLog)
	return result.RowsAffected, nil
}

// 事务
func (repo *postUpdateLogRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *postUpdateLogRepository) SaveInRedis(postUpdateLog *model.PostUpdateLog) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = postUpdateLog.RedisKey()
	resByte, e := json.Marshal(postUpdateLog)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *postUpdateLogRepository) FindInRedis(postUpdateLog *model.PostUpdateLog) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = postUpdateLog.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), postUpdateLog)
	}
	return nil
}

func (repo *postUpdateLogRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *postUpdateLogRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *postUpdateLogRepository) DeleteInRedis(postUpdateLog *model.PostUpdateLog) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = postUpdateLog.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *postUpdateLogRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.PostUpdateLog, e error) {
	now := time.Now()
	postUpdateLog := &model.PostUpdateLog{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := postUpdateLog.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(postUpdateLog)
		if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
			var count64 int64
			e = db.Count(&count64).Error
			count := int(count64)
			if e != nil {
				return nil, e
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
		return list, e
	}
	db := global.DB.Table(postUpdateLog.TableName()).Where(query)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
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
	return list, nil
}

func (repo *postUpdateLogRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.PostUpdateLog, e error) {
	now := time.Now()
	postUpdateLog := &model.PostUpdateLog{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := postUpdateLog.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(postUpdateLog)
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
		return list, e
	}
	db := global.DB.Table(postUpdateLog.TableName())
	if query != "" {
		db = db.Where(query, args...)
	}
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
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
	return list, nil
}

func (repo *postUpdateLogRepository) GetDataByWhereMap(postUpdateLog *model.PostUpdateLog, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(postUpdateLog.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Table(postUpdateLog.TableName()).Where(where)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db = db.First(postUpdateLog)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(postUpdateLog)
	return nil
}

func (repo *postUpdateLogRepository) Execute(db *gorm.DB, object interface{}) error {
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
