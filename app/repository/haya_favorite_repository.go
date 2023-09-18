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

type hayaFavoriteRepository struct {
	Pager *Pager
}

var HayaFavoriteRepository = newHayaFavoriteRepository()

func newHayaFavoriteRepository() *hayaFavoriteRepository {
	return new(hayaFavoriteRepository)
}

func (repo *hayaFavoriteRepository) Insert(hayaFavorite *model.HayaFavorite) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(hayaFavorite)
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(hayaFavorite)
	return result.RowsAffected, result.Error
}

func (repo *hayaFavoriteRepository) Update(hayaFavorite *model.HayaFavorite) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "Update", e, now)
		}
	}()
	if len(hayaFavorite.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := hayaFavorite.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(hayaFavorite.TableName()).Where(hayaFavorite.Location()).Updates(updateValues)
	if result.Error != nil {
		return
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(hayaFavorite)
	repo.First(hayaFavorite)
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *hayaFavoriteRepository) First(hayaFavorite *model.HayaFavorite) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "First", e, now)
		}
	}()
	if len(hayaFavorite.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	err := repo.FindInRedis(hayaFavorite)
	if err != nil && err != redis.Nil {
		return err
	}
	result := global.DB.Table(hayaFavorite.TableName()).Where(hayaFavorite.Location()).First(hayaFavorite)
	e = result.Error
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(hayaFavorite)
	return
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *hayaFavoriteRepository) DeleteByLocation(hayaFavorite *model.HayaFavorite) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(hayaFavorite.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(hayaFavorite.TableName()).Where(hayaFavorite.Location()).Unscoped().Delete(hayaFavorite)
	if result.Error != nil {
		return
	}
	repo.DeleteInRedis(hayaFavorite)
	rowsAffected = result.RowsAffected
	e = result.Error
	return
}

// 事务
func (repo *hayaFavoriteRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *hayaFavoriteRepository) SaveInRedis(hayaFavorite *model.HayaFavorite) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = hayaFavorite.RedisKey()
	resByte, e := json.Marshal(hayaFavorite)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *hayaFavoriteRepository) FindInRedis(hayaFavorite *model.HayaFavorite) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = hayaFavorite.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), hayaFavorite)
	}
	return
}

func (repo *hayaFavoriteRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *hayaFavoriteRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *hayaFavoriteRepository) DeleteInRedis(hayaFavorite *model.HayaFavorite) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = hayaFavorite.RedisKey()
	err := global.REDIS.Del(context.Background(), redisKey).Err()
	if err != nil {
		return
	}
	return nil
}
func (repo *hayaFavoriteRepository) GetDataListByWhereMap(query map[string]interface{}) (list []*model.HayaFavorite, e error) {
	now := time.Now()
	hayaFavorite := &model.HayaFavorite{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := hayaFavorite.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(hayaFavorite)
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
	db := global.DB.Table(hayaFavorite.TableName()).Where(query)
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

func (repo *hayaFavoriteRepository) GetDataListByWhere(query string, args []interface{}) (list []*model.HayaFavorite, e error) {
	now := time.Now()
	hayaFavorite := &model.HayaFavorite{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := hayaFavorite.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(hayaFavorite)
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
	db := global.DB.Table(hayaFavorite.TableName())
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

func (repo *hayaFavoriteRepository) GetDataByWhereMap(hayaFavorite *model.HayaFavorite, where map[string]interface{}) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(hayaFavorite.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Table(hayaFavorite.TableName()).Where(where).First(hayaFavorite)
	e = db.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(hayaFavorite)
	return
}

func (repo *hayaFavoriteRepository) Execute(db *gorm.DB, object interface{}) error {
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
