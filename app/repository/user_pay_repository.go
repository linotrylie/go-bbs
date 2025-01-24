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

type UserPayRepository struct {
	Pager *Pager
}

var userPayRepository = newUserPayRepository()

func newUserPayRepository() *UserPayRepository {
	return new(UserPayRepository)
}

func (repo *UserPayRepository) Insert(m *model.UserPay) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(m)
	e = result.Error
	if e != nil {
		return
	}
	return result.RowsAffected, e
}

func (repo *UserPayRepository) Update(m *model.UserPay) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "Update", e, now)
		}
	}()
	if len(m.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := m.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(m).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(m)
	repo.First(m, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *UserPayRepository) First(m *model.UserPay, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "First", e, now)
		}
	}()
	if len(m.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(m)
	if e == nil {
		return
	}
	db := global.DB.Table(m.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(m)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(m)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *UserPayRepository) DeleteByLocation(m *model.UserPay) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(m.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(m.TableName()).Unscoped().Delete(m)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(m)
	return result.RowsAffected, nil
}

// 事务
func (repo *UserPayRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *UserPayRepository) SaveInRedis(m *model.UserPay) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	if m.IsCache() {
		return nil
	}
	var redisKey string
	redisKey = m.RedisKey()
	resByte, e := json.Marshal(m)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *UserPayRepository) FindInRedis(m *model.UserPay) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	if m.IsCache() {
		return nil
	}
	var redisKey string
	redisKey = m.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil || e != redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), m)
	}
	return nil
}

func (repo *UserPayRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	redisRes, e = global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil || e != redis.Nil {
		return "", nil
	} else {
		return "", nil
	}
	return "", nil
}

func (repo *UserPayRepository) SaveInRedisByKey(redisKey string, data string, timeout int) {
	var timeSecond time.Duration
	if timeout > 0 {
		timeSecond = time.Duration(timeout) * time.Second
	} else {
		timeSecond = time.Duration(random.RandInt(7200, 14400)) * time.Second
	}
	global.REDIS.Set(context.Background(), redisKey, data, timeSecond)
}

func (repo *UserPayRepository) DeleteInRedis(m *model.UserPay) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	if m.IsCache() {
		return nil
	}
	var redisKey string
	redisKey = m.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *UserPayRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.UserPay, e error) {
	m := &model.UserPay{}
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "GetDataListByWhereMap", e, now)
		}
	}()
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	for kk, vv := range query {
		str += kk + Strval(vv)
	}
	var redisKey string
	redisKey = m.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	//如果模型是缓存类则先查询缓存内的数据
	if m.IsCache() {
		val, _ := repo.FindInRedisByKey(redisKey)
		if val != "" {
			e = json.Unmarshal([]byte(val), &list)
			if e != nil {
				return nil, e
			}
		}
	}
	//缓存内没有则查询数据库
	db := global.DB.Model(m).Where(query)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	e = repo.Execute(db, &list)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, exceptions.NotFoundData
	}
	//将本次查询结果缓存起来
	marshal, e := json.Marshal(list)
	if e != nil {
		return nil, e
	}
	repo.SaveInRedisByKey(redisKey, string(marshal), 5)
	return
}

func (repo *UserPayRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.UserPay, e error) {
	now := time.Now()
	m := &model.UserPay{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	var redisKey string
	if m.IsCache() {
		if repo.Pager.FieldsOrder != nil {
			for _, v := range repo.Pager.FieldsOrder {
				str += strings.Replace(v, " ", "", -1)
			}
		}
		if query != "" {
			for _, vv := range args {
				str += Strval(vv)
			}
		}
		redisKey := m.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
		val, _ := repo.FindInRedisByKey(redisKey)
		if val != "" {
			e = json.Unmarshal([]byte(val), &list)
			if e != nil {
				return nil, e
			}
			return
		}
	}
	db := global.DB.Model(m)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
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
	repo.SaveInRedisByKey(redisKey, string(marshal), 5)
	return
}

func (repo *UserPayRepository) GetDataByWhereMap(m *model.UserPay, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(m.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	if m.IsCache() {
		e = repo.FindInRedis(m)
		if e == nil {
			return
		}
	}
	db := global.DB.Model(m).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(m)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(m)
	return nil
}

func (repo *UserPayRepository) Execute(db *gorm.DB, object interface{}) error {
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

func (repo *UserPayRepository) GetTotalPage(db *gorm.DB) (e error) {
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
