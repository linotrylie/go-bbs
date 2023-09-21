package converter

func GetRepositoryTemplate() string {
	return `
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

type %sRepository struct {
	Pager *Pager
}
var %sRepository = new%sRepository()

func new%sRepository() *%sRepository {
	return new(%sRepository)
}

func (repo *%sRepository) Insert(%s *model.%s) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(%s)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(%s)
	return result.RowsAffected, e
}

func (repo *%sRepository) Update(%s *model.%s) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "Update", e, now)
		}
	}()
	if len(%s.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := %s.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(%s).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(%s)
	repo.First(%s,[]string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *%sRepository) First(%s *model.%s, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "First", e, now)
		}
	}()
	if len(%s.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	repo.FindInRedis(%s)
	if %s != nil {
		return nil
	}
	db := global.DB.Table(%s.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(%s)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(%s)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *%sRepository) DeleteByLocation(%s *model.%s) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(%s.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(%s.TableName()).Unscoped().Delete(%s)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(%s)
	return result.RowsAffected,nil
}

// 事务
func (repo *%sRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *%sRepository) SaveInRedis(%s *model.%s) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = %s.RedisKey()
	resByte, e := json.Marshal(%s)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *%sRepository) FindInRedis(%s *model.%s) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = %s.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), %s)
	}
	return nil
}

func (repo *%sRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *%sRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *%sRepository) DeleteInRedis(%s *model.%s) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = %s.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *%sRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.%s, e error) {
	now := time.Now()
	%s := &model.%s{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	for k,vv := range query {
		str += k + Strval(vv)
	}
	redisKey := %s.TableName() + "_list_"+ strconv.Itoa( repo.Pager.Page ) +"_"+ strconv.Itoa( repo.Pager.PageSize ) +"_" + str;
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(%s).Where(query)
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

func (repo *%sRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.%s, e error) {
	now := time.Now()
	%s := &model.%s{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(%s)
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
	redisKey := %s.TableName() + "_list_"+ strconv.Itoa( repo.Pager.Page ) +"_"+ strconv.Itoa( repo.Pager.PageSize ) +"_" + str;
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

func (repo *%sRepository) GetDataByWhereMap(%s *model.%s,where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	repo.FindInRedis(%s)
	if %s != nil {
		return nil
	}
	db := global.DB.Model(%s).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(%s)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(%s)
	return nil
}

func (repo *%sRepository) Execute(db *gorm.DB, object interface{}) error {
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

func (repo *%sRepository) GetTotalPage(db *gorm.DB) (e error) {
		if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
			var count64 int64
			e = db.Count(&count64).Error
			count := int(count64)
			if e != nil {
				return e
			}
			if count != 0 {
				//Calculate the length of the pagination
				if count % repo.Pager.PageSize == 0 {
					repo.Pager.TotalPage = int64(count / repo.Pager.PageSize)
				} else {
					repo.Pager.TotalPage = int64(count / repo.Pager.PageSize + 1)
				}
			}
		}
		return nil
}

`
}
