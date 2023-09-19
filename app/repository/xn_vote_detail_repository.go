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

type xnVoteDetailRepository struct {
	Pager *Pager
}

var XnVoteDetailRepository = newXnVoteDetailRepository()

func newXnVoteDetailRepository() *xnVoteDetailRepository {
	return new(xnVoteDetailRepository)
}

func (repo *xnVoteDetailRepository) Insert(xnVoteDetail *model.XnVoteDetail) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(xnVoteDetail)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(xnVoteDetail)
	return result.RowsAffected, e
}

func (repo *xnVoteDetailRepository) Update(xnVoteDetail *model.XnVoteDetail) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "Update", e, now)
		}
	}()
	if len(xnVoteDetail.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := xnVoteDetail.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(xnVoteDetail).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(xnVoteDetail)
	repo.First(xnVoteDetail, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *xnVoteDetailRepository) First(xnVoteDetail *model.XnVoteDetail, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "First", e, now)
		}
	}()
	if len(xnVoteDetail.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(xnVoteDetail)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(xnVoteDetail.TableName())
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db.First(xnVoteDetail)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(xnVoteDetail)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *xnVoteDetailRepository) DeleteByLocation(xnVoteDetail *model.XnVoteDetail) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(xnVoteDetail.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(xnVoteDetail.TableName()).Unscoped().Delete(xnVoteDetail)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(xnVoteDetail)
	return result.RowsAffected, nil
}

// 事务
func (repo *xnVoteDetailRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *xnVoteDetailRepository) SaveInRedis(xnVoteDetail *model.XnVoteDetail) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = xnVoteDetail.RedisKey()
	resByte, e := json.Marshal(xnVoteDetail)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *xnVoteDetailRepository) FindInRedis(xnVoteDetail *model.XnVoteDetail) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = xnVoteDetail.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), xnVoteDetail)
	}
	return nil
}

func (repo *xnVoteDetailRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *xnVoteDetailRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *xnVoteDetailRepository) DeleteInRedis(xnVoteDetail *model.XnVoteDetail) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = xnVoteDetail.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *xnVoteDetailRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.XnVoteDetail, e error) {
	now := time.Now()
	xnVoteDetail := &model.XnVoteDetail{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "DeleteByLocation", e, now)
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
	redisKey := xnVoteDetail.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(xnVoteDetail).Where(query)
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

func (repo *xnVoteDetailRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.XnVoteDetail, e error) {
	now := time.Now()
	xnVoteDetail := &model.XnVoteDetail{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(xnVoteDetail)
	if query != "" {
		db = db.Where(query, args...)
		for _, vv := range args {
			str += Strval(vv)
		}
	}
	redisKey := xnVoteDetail.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
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

func (repo *xnVoteDetailRepository) GetDataByWhereMap(xnVoteDetail *model.XnVoteDetail, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(xnVoteDetail.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Model(xnVoteDetail).Where(where)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db = db.First(xnVoteDetail)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(xnVoteDetail)
	return nil
}

func (repo *xnVoteDetailRepository) Execute(db *gorm.DB, object interface{}) error {
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
