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

type ggFavoriteThreadRepository struct {
	Pager *Pager
}

var GgFavoriteThreadRepository = newGgFavoriteThreadRepository()

func newGgFavoriteThreadRepository() *ggFavoriteThreadRepository {
	return new(ggFavoriteThreadRepository)
}

func (repo *ggFavoriteThreadRepository) Insert(ggFavoriteThread *model.GgFavoriteThread) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(ggFavoriteThread)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(ggFavoriteThread)
	return result.RowsAffected, e
}

func (repo *ggFavoriteThreadRepository) Update(ggFavoriteThread *model.GgFavoriteThread) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "Update", e, now)
		}
	}()
	if len(ggFavoriteThread.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := ggFavoriteThread.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(ggFavoriteThread).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(ggFavoriteThread)
	repo.First(ggFavoriteThread, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *ggFavoriteThreadRepository) First(ggFavoriteThread *model.GgFavoriteThread, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "First", e, now)
		}
	}()
	if len(ggFavoriteThread.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(ggFavoriteThread)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(ggFavoriteThread.TableName())
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db.First(ggFavoriteThread)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(ggFavoriteThread)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *ggFavoriteThreadRepository) DeleteByLocation(ggFavoriteThread *model.GgFavoriteThread) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(ggFavoriteThread.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(ggFavoriteThread.TableName()).Unscoped().Delete(ggFavoriteThread)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(ggFavoriteThread)
	return result.RowsAffected, nil
}

// 事务
func (repo *ggFavoriteThreadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *ggFavoriteThreadRepository) SaveInRedis(ggFavoriteThread *model.GgFavoriteThread) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = ggFavoriteThread.RedisKey()
	resByte, e := json.Marshal(ggFavoriteThread)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *ggFavoriteThreadRepository) FindInRedis(ggFavoriteThread *model.GgFavoriteThread) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = ggFavoriteThread.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), ggFavoriteThread)
	}
	return nil
}

func (repo *ggFavoriteThreadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *ggFavoriteThreadRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *ggFavoriteThreadRepository) DeleteInRedis(ggFavoriteThread *model.GgFavoriteThread) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = ggFavoriteThread.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *ggFavoriteThreadRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.GgFavoriteThread, e error) {
	now := time.Now()
	ggFavoriteThread := &model.GgFavoriteThread{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	for k, vv := range query {
		str += k + Strval(vv)
	}
	redisKey := ggFavoriteThread.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(ggFavoriteThread).Where(query)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	if val != "" {
		e = json.Unmarshal([]byte(val), &list)
		if e != nil {
			return nil, e
		}
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
					repo.Pager.TotalPage = int64(count / repo.Pager.PageSize)
				} else {
					repo.Pager.TotalPage = int64(count/repo.Pager.PageSize + 1)
				}
			}
		}
		return list, e
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

func (repo *ggFavoriteThreadRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.GgFavoriteThread, e error) {
	now := time.Now()
	ggFavoriteThread := &model.GgFavoriteThread{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(ggFavoriteThread)
	if query != "" {
		db = db.Where(query, args...)
		for _, vv := range args {
			str += Strval(vv)
		}
	}
	redisKey := ggFavoriteThread.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	if val != "" {
		e = json.Unmarshal([]byte(val), &list)
		if e != nil {
			return nil, e
		}

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
					repo.Pager.TotalPage = int64(count / repo.Pager.PageSize)
				} else {
					repo.Pager.TotalPage = int64(count/repo.Pager.PageSize + 1)
				}
			}
		}
		return list, e
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

func (repo *ggFavoriteThreadRepository) GetDataByWhereMap(ggFavoriteThread *model.GgFavoriteThread, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(ggFavoriteThread.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Model(ggFavoriteThread).Where(where)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db = db.First(ggFavoriteThread)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(ggFavoriteThread)
	return nil
}

func (repo *ggFavoriteThreadRepository) Execute(db *gorm.DB, object interface{}) error {
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
				repo.Pager.TotalPage = int64(count / repo.Pager.PageSize)
			} else {
				repo.Pager.TotalPage = int64(count/repo.Pager.PageSize + 1)
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
