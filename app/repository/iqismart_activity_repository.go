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

type iqismartActivityRepository struct {
	Pager *Pager
}

var IqismartActivityRepository = newIqismartActivityRepository()

func newIqismartActivityRepository() *iqismartActivityRepository {
	return new(iqismartActivityRepository)
}

func (repo *iqismartActivityRepository) Insert(iqismartActivity *model.IqismartActivity) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(iqismartActivity)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(iqismartActivity)
	return result.RowsAffected, e
}

func (repo *iqismartActivityRepository) Update(iqismartActivity *model.IqismartActivity) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "Update", e, now)
		}
	}()
	if len(iqismartActivity.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := iqismartActivity.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(iqismartActivity).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(iqismartActivity)
	repo.First(iqismartActivity, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *iqismartActivityRepository) First(iqismartActivity *model.IqismartActivity, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "First", e, now)
		}
	}()
	if len(iqismartActivity.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(iqismartActivity)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(iqismartActivity.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(iqismartActivity)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(iqismartActivity)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *iqismartActivityRepository) DeleteByLocation(iqismartActivity *model.IqismartActivity) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(iqismartActivity.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(iqismartActivity.TableName()).Unscoped().Delete(iqismartActivity)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(iqismartActivity)
	return result.RowsAffected, nil
}

// 事务
func (repo *iqismartActivityRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *iqismartActivityRepository) SaveInRedis(iqismartActivity *model.IqismartActivity) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartActivity.RedisKey()
	resByte, e := json.Marshal(iqismartActivity)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *iqismartActivityRepository) FindInRedis(iqismartActivity *model.IqismartActivity) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartActivity.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), iqismartActivity)
	}
	return nil
}

func (repo *iqismartActivityRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *iqismartActivityRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *iqismartActivityRepository) DeleteInRedis(iqismartActivity *model.IqismartActivity) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartActivity.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *iqismartActivityRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.IqismartActivity, e error) {
	now := time.Now()
	iqismartActivity := &model.IqismartActivity{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "DeleteByLocation", e, now)
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
	redisKey := iqismartActivity.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(iqismartActivity).Where(query)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	if val != "" {
		e = json.Unmarshal([]byte(val), &list)
		if e != nil {
			return nil, e
		}
		e = repo.GetTotalPage(db)
		if e != nil {
			return nil, e
		}
		return
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

func (repo *iqismartActivityRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.IqismartActivity, e error) {
	now := time.Now()
	iqismartActivity := &model.IqismartActivity{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(iqismartActivity)
	if query != "" {
		db = db.Where(query, args...)
		for _, vv := range args {
			str += Strval(vv)
		}
	}
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	redisKey := iqismartActivity.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	if val != "" {
		e = json.Unmarshal([]byte(val), &list)
		if e != nil {
			return nil, e
		}
		e = repo.GetTotalPage(db)
		if e != nil {
			return nil, e
		}
		return
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

func (repo *iqismartActivityRepository) GetDataByWhereMap(iqismartActivity *model.IqismartActivity, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivity.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Model(iqismartActivity).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(iqismartActivity)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(iqismartActivity)
	return nil
}

func (repo *iqismartActivityRepository) Execute(db *gorm.DB, object interface{}) error {
	e := repo.GetTotalPage(db)
	if e != nil {
		return e
	}
	orderValue := repo.Pager.FieldsOrder
	if len(orderValue) > 0 {
		for _, v := range orderValue {
			db.Order(v)
		}
	}
	resultDB := db.Find(object)
	if resultDB.Error != nil {
		return resultDB.Error
	}
	return nil
}

func (repo *iqismartActivityRepository) GetTotalPage(db *gorm.DB) (e error) {
	if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
		var count64 int64
		e = db.Count(&count64).Error
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
	}
	return nil
}
