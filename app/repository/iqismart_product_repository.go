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

type iqismartProductRepository struct {
	Pager *Pager
}

var IqismartProductRepository = newIqismartProductRepository()

func newIqismartProductRepository() *iqismartProductRepository {
	return new(iqismartProductRepository)
}

func (repo *iqismartProductRepository) Insert(iqismartProduct *model.IqismartProduct) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(iqismartProduct)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(iqismartProduct)
	return result.RowsAffected, e
}

func (repo *iqismartProductRepository) Update(iqismartProduct *model.IqismartProduct) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "Update", e, now)
		}
	}()
	if len(iqismartProduct.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := iqismartProduct.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(iqismartProduct).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(iqismartProduct)
	repo.First(iqismartProduct, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *iqismartProductRepository) First(iqismartProduct *model.IqismartProduct, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "First", e, now)
		}
	}()
	if len(iqismartProduct.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(iqismartProduct)
	if e == nil {
		return
	}
	db := global.DB.Table(iqismartProduct.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(iqismartProduct)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(iqismartProduct)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *iqismartProductRepository) DeleteByLocation(iqismartProduct *model.IqismartProduct) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(iqismartProduct.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(iqismartProduct.TableName()).Unscoped().Delete(iqismartProduct)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(iqismartProduct)
	return result.RowsAffected, nil
}

// 事务
func (repo *iqismartProductRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *iqismartProductRepository) SaveInRedis(iqismartProduct *model.IqismartProduct) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartProduct.RedisKey()
	resByte, e := json.Marshal(iqismartProduct)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *iqismartProductRepository) FindInRedis(iqismartProduct *model.IqismartProduct) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartProduct.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), iqismartProduct)
	}
	return nil
}

func (repo *iqismartProductRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *iqismartProductRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *iqismartProductRepository) DeleteInRedis(iqismartProduct *model.IqismartProduct) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartProduct.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *iqismartProductRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.IqismartProduct, e error) {
	now := time.Now()
	iqismartProduct := &model.IqismartProduct{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "DeleteByLocation", e, now)
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
	redisKey := iqismartProduct.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(iqismartProduct).Where(query)
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

func (repo *iqismartProductRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.IqismartProduct, e error) {
	now := time.Now()
	iqismartProduct := &model.IqismartProduct{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(iqismartProduct)
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
	redisKey := iqismartProduct.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
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

func (repo *iqismartProductRepository) GetDataByWhereMap(iqismartProduct *model.IqismartProduct, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	e = repo.FindInRedis(iqismartProduct)
	if e == nil {
		return
	}
	db := global.DB.Model(iqismartProduct).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(iqismartProduct)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(iqismartProduct)
	return nil
}

func (repo *iqismartProductRepository) Execute(db *gorm.DB, object interface{}) error {
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

func (repo *iqismartProductRepository) GetTotalPage(db *gorm.DB) (e error) {
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
