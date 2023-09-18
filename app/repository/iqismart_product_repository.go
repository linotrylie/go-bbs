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
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(iqismartProduct)
	return result.RowsAffected, result.Error
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
		return 0, errors.New("location cannot be empty")
	}
	updateValues := iqismartProduct.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(iqismartProduct.TableName()).Where(iqismartProduct.Location()).Updates(updateValues)
	if result.Error != nil {
		return
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(iqismartProduct)
	repo.First(iqismartProduct)
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *iqismartProductRepository) First(iqismartProduct *model.IqismartProduct) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "First", e, now)
		}
	}()
	if len(iqismartProduct.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	err := repo.FindInRedis(iqismartProduct)
	if err != nil && err != redis.Nil {
		return err
	}
	result := global.DB.Table(iqismartProduct.TableName()).Where(iqismartProduct.Location()).First(iqismartProduct)
	e = result.Error
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(iqismartProduct)
	return
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
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(iqismartProduct.TableName()).Where(iqismartProduct.Location()).Unscoped().Delete(iqismartProduct)
	if result.Error != nil {
		return
	}
	repo.DeleteInRedis(iqismartProduct)
	rowsAffected = result.RowsAffected
	e = result.Error
	return
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
	return
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

func (repo *iqismartProductRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *iqismartProductRepository) DeleteInRedis(iqismartProduct *model.IqismartProduct) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = iqismartProduct.RedisKey()
	err := global.REDIS.Del(context.Background(), redisKey).Err()
	if err != nil {
		return
	}
	return nil
}
func (repo *iqismartProductRepository) GetDataListByWhereMap(query map[string]interface{}) (list []*model.IqismartProduct, e error) {
	now := time.Now()
	iqismartProduct := &model.IqismartProduct{}
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
	redisKey := iqismartProduct.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(iqismartProduct)
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
	db := global.DB.Table(iqismartProduct.TableName()).Where(query)
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

func (repo *iqismartProductRepository) GetDataListByWhere(query string, args []interface{}) (list []*model.IqismartProduct, e error) {
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
	redisKey := iqismartProduct.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(iqismartProduct)
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
	db := global.DB.Table(iqismartProduct.TableName())
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

func (repo *iqismartProductRepository) GetDataByWhereMap(iqismartProduct *model.IqismartProduct, where map[string]interface{}) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(iqismartProduct.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Table(iqismartProduct.TableName()).Where(where).First(iqismartProduct)
	e = db.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(iqismartProduct)
	return
}

func (repo *iqismartProductRepository) Execute(db *gorm.DB, object interface{}) error {
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
