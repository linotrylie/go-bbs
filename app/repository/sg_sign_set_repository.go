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

type sgSignSetRepository struct {
	Pager *Pager
}

var SgSignSetRepository = newSgSignSetRepository()

func newSgSignSetRepository() *sgSignSetRepository {
	return new(sgSignSetRepository)
}

func (repo *sgSignSetRepository) Insert(sgSignSet *model.SgSignSet) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(sgSignSet)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(sgSignSet)
	return result.RowsAffected, e
}

func (repo *sgSignSetRepository) Update(sgSignSet *model.SgSignSet) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "Update", e, now)
		}
	}()
	if len(sgSignSet.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := sgSignSet.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(sgSignSet).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(sgSignSet)
	repo.First(sgSignSet, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *sgSignSetRepository) First(sgSignSet *model.SgSignSet, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "First", e, now)
		}
	}()
	if len(sgSignSet.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(sgSignSet)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(sgSignSet.TableName())
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db.First(sgSignSet)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(sgSignSet)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *sgSignSetRepository) DeleteByLocation(sgSignSet *model.SgSignSet) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(sgSignSet.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(sgSignSet.TableName()).Unscoped().Delete(sgSignSet)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(sgSignSet)
	return result.RowsAffected, nil
}

// 事务
func (repo *sgSignSetRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *sgSignSetRepository) SaveInRedis(sgSignSet *model.SgSignSet) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = sgSignSet.RedisKey()
	resByte, e := json.Marshal(sgSignSet)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *sgSignSetRepository) FindInRedis(sgSignSet *model.SgSignSet) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = sgSignSet.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), sgSignSet)
	}
	return nil
}

func (repo *sgSignSetRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *sgSignSetRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *sgSignSetRepository) DeleteInRedis(sgSignSet *model.SgSignSet) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = sgSignSet.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *sgSignSetRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.SgSignSet, e error) {
	now := time.Now()
	sgSignSet := &model.SgSignSet{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "DeleteByLocation", e, now)
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
	redisKey := sgSignSet.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(sgSignSet).Where(query)
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

func (repo *sgSignSetRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.SgSignSet, e error) {
	now := time.Now()
	sgSignSet := &model.SgSignSet{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(sgSignSet)
	if query != "" {
		db = db.Where(query, args...)
		for _, vv := range args {
			str += Strval(vv)
		}
	}
	redisKey := sgSignSet.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
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

func (repo *sgSignSetRepository) GetDataByWhereMap(sgSignSet *model.SgSignSet, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(sgSignSet.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Model(sgSignSet).Where(where)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db = db.First(sgSignSet)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(sgSignSet)
	return nil
}

func (repo *sgSignSetRepository) Execute(db *gorm.DB, object interface{}) error {
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
