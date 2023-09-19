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

type gitTagsRepository struct {
	Pager *Pager
}

var GitTagsRepository = newGitTagsRepository()

func newGitTagsRepository() *gitTagsRepository {
	return new(gitTagsRepository)
}

func (repo *gitTagsRepository) Insert(gitTags *model.GitTags) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(gitTags)
	if result.Error != nil {

		return
	}
	repo.SaveInRedis(gitTags)
	return result.RowsAffected, result.Error
}

func (repo *gitTagsRepository) Update(gitTags *model.GitTags) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "Update", e, now)
		}
	}()
	if len(gitTags.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := gitTags.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Model(gitTags).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(gitTags)
	repo.First(gitTags, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *gitTagsRepository) First(gitTags *model.GitTags, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "First", e, now)
		}
	}()
	if len(gitTags.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	e = repo.FindInRedis(gitTags)
	if e != nil && e != redis.Nil {
		return e
	}
	db := global.DB.Table(gitTags.TableName())
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db.First(gitTags)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(gitTags)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *gitTagsRepository) DeleteByLocation(gitTags *model.GitTags) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(gitTags.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(gitTags.TableName()).Unscoped().Delete(gitTags)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(gitTags)
	return result.RowsAffected, nil
}

// 事务
func (repo *gitTagsRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *gitTagsRepository) SaveInRedis(gitTags *model.GitTags) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = gitTags.RedisKey()
	resByte, e := json.Marshal(gitTags)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *gitTagsRepository) FindInRedis(gitTags *model.GitTags) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = gitTags.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), gitTags)
	}
	return nil
}

func (repo *gitTagsRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *gitTagsRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *gitTagsRepository) DeleteInRedis(gitTags *model.GitTags) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = gitTags.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *gitTagsRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.GitTags, e error) {
	now := time.Now()
	gitTags := &model.GitTags{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := gitTags.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(gitTags)
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
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 0
				} else {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
				}
			}
		}
		return list, e
	}
	db := global.DB.Model(gitTags).Where(query)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
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

func (repo *gitTagsRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.GitTags, e error) {
	now := time.Now()
	gitTags := &model.GitTags{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := gitTags.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(gitTags)
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
		return list, e
	}
	db := global.DB.Model(gitTags)
	if query != "" {
		db = db.Where(query, args...)
	}
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
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

func (repo *gitTagsRepository) GetDataByWhereMap(gitTags *model.GitTags, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTags.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Model(gitTags).Where(where)
	if preload != nil {
		for _, v := range preload {
			db = db.Preload(v)
		}
	}
	db = db.First(gitTags)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(gitTags)
	return nil
}

func (repo *gitTagsRepository) Execute(db *gorm.DB, object interface{}) error {
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
