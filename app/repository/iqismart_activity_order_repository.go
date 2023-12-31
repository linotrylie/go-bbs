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

type iqismartActivityOrderRepository struct {
	Pager *Pager
}

var IqismartActivityOrderRepository = newIqismartActivityOrderRepository()

func newIqismartActivityOrderRepository() *iqismartActivityOrderRepository {
	return new(iqismartActivityOrderRepository)
}

func (repo *iqismartActivityOrderRepository) Insert(iqismartActivityOrder *model.IqismartActivityOrder) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(iqismartActivityOrder)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(iqismartActivityOrder)
	return result.RowsAffected, e
}

func (repo *iqismartActivityOrderRepository) Update(iqismartActivityOrder *model.IqismartActivityOrder) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "Update", e, now)
		}
	}()
	if len(iqismartActivityOrder.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := iqismartActivityOrder.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(iqismartActivityOrder).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(iqismartActivityOrder)
	repo.First(iqismartActivityOrder, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *iqismartActivityOrderRepository) First(iqismartActivityOrder *model.IqismartActivityOrder, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "First", e, now)
		}
	}()
	if len(iqismartActivityOrder.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(iqismartActivityOrder)
	if e == nil {
		return
	}
	db := global.DB.Table(iqismartActivityOrder.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(iqismartActivityOrder)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(iqismartActivityOrder)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *iqismartActivityOrderRepository) DeleteByLocation(iqismartActivityOrder *model.IqismartActivityOrder) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(iqismartActivityOrder.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(iqismartActivityOrder.TableName()).Unscoped().Delete(iqismartActivityOrder)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(iqismartActivityOrder)
	return result.RowsAffected, nil
}

// 事务
func (repo *iqismartActivityOrderRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *iqismartActivityOrderRepository) SaveInRedis(iqismartActivityOrder *model.IqismartActivityOrder) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartActivityOrder.RedisKey()
	resByte, e := json.Marshal(iqismartActivityOrder)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *iqismartActivityOrderRepository) FindInRedis(iqismartActivityOrder *model.IqismartActivityOrder) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartActivityOrder.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), iqismartActivityOrder)
	}
	return nil
}

func (repo *iqismartActivityOrderRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *iqismartActivityOrderRepository) SaveInRedisByKey(redisKey string, data string, timeout int) {
	var timeSecond time.Duration
	if timeout > 0 {
		timeSecond = time.Duration(timeout) * time.Second
	} else {
		timeSecond = time.Duration(random.RandInt(7200, 14400)) * time.Second
	}
	global.REDIS.Set(context.Background(), redisKey, data, timeSecond)
}

func (repo *iqismartActivityOrderRepository) DeleteInRedis(iqismartActivityOrder *model.IqismartActivityOrder) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartActivityOrder.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *iqismartActivityOrderRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.IqismartActivityOrder, e error) {
	now := time.Now()
	iqismartActivityOrder := &model.IqismartActivityOrder{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "DeleteByLocation", e, now)
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
	redisKey := iqismartActivityOrder.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(iqismartActivityOrder).Where(query)
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
	repo.SaveInRedisByKey(redisKey, string(marshal), 5)
	return
}

func (repo *iqismartActivityOrderRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.IqismartActivityOrder, e error) {
	now := time.Now()
	iqismartActivityOrder := &model.IqismartActivityOrder{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(iqismartActivityOrder)
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
	redisKey := iqismartActivityOrder.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
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
	repo.SaveInRedisByKey(redisKey, string(marshal), 5)
	return
}

func (repo *iqismartActivityOrderRepository) GetDataByWhereMap(iqismartActivityOrder *model.IqismartActivityOrder, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartActivityOrder.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	e = repo.FindInRedis(iqismartActivityOrder)
	if e == nil {
		return
	}
	db := global.DB.Model(iqismartActivityOrder).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(iqismartActivityOrder)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(iqismartActivityOrder)
	return nil
}

func (repo *iqismartActivityOrderRepository) Execute(db *gorm.DB, object interface{}) error {
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

func (repo *iqismartActivityOrderRepository) GetTotalPage(db *gorm.DB) (e error) {
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
		db = db.Offset((repo.Pager.Page - 1) * repo.Pager.PageSize).Limit(repo.Pager.PageSize)
	}
	return nil
}
