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

type modlogRepository struct {
	Pager *Pager
}

var ModlogRepository = newModlogRepository()

func newModlogRepository() *modlogRepository {
	return new(modlogRepository)
}

func (repo *modlogRepository) Insert(modlog *model.Modlog) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(modlog)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(modlog)
	return result.RowsAffected, e
}

func (repo *modlogRepository) Update(modlog *model.Modlog) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "Update", e, now)
		}
	}()
	if len(modlog.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := modlog.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(modlog).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(modlog)
	repo.First(modlog, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *modlogRepository) First(modlog *model.Modlog, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "First", e, now)
		}
	}()
	if len(modlog.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	e = repo.FindInRedis(modlog)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(modlog.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(modlog)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(modlog)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *modlogRepository) DeleteByLocation(modlog *model.Modlog) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(modlog.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(modlog.TableName()).Unscoped().Delete(modlog)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(modlog)
	return result.RowsAffected, nil
}

// 事务
func (repo *modlogRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *modlogRepository) SaveInRedis(modlog *model.Modlog) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = modlog.RedisKey()
	resByte, e := json.Marshal(modlog)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *modlogRepository) FindInRedis(modlog *model.Modlog) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = modlog.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), modlog)
	}
	return nil
}

func (repo *modlogRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *modlogRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *modlogRepository) DeleteInRedis(modlog *model.Modlog) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = modlog.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *modlogRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.Modlog, e error) {
	now := time.Now()
	modlog := &model.Modlog{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "DeleteByLocation", e, now)
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
	redisKey := modlog.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(modlog).Where(query)
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

func (repo *modlogRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.Modlog, e error) {
	now := time.Now()
	modlog := &model.Modlog{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(modlog)
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
	redisKey := modlog.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
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

func (repo *modlogRepository) GetDataByWhereMap(modlog *model.Modlog, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(modlog.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Model(modlog).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(modlog)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(modlog)
	return nil
}

func (repo *modlogRepository) Execute(db *gorm.DB, object interface{}) error {
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

func (repo *modlogRepository) GetTotalPage(db *gorm.DB) (e error) {
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
