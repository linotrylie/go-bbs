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

type gitTagsThreadRepository struct {
	Pager *Pager
}

var GitTagsThreadRepository = newGitTagsThreadRepository()

func newGitTagsThreadRepository() *gitTagsThreadRepository {
	return new(gitTagsThreadRepository)
}

func (repo *gitTagsThreadRepository) Insert(gitTagsThread *model.GitTagsThread) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(gitTagsThread)
	e = result.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(gitTagsThread)
	return result.RowsAffected, e
}

func (repo *gitTagsThreadRepository) Update(gitTagsThread *model.GitTagsThread) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "Update", e, now)
		}
	}()
	if len(gitTagsThread.Location()) == 0 {
		return 0, errors.New("无更新条件！")
	}
	updateValues := gitTagsThread.GetChanges()
	if len(updateValues) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Model(gitTagsThread).Updates(updateValues)
	e = result.Error
	if e != nil {
		return 0, e
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(gitTagsThread)
	repo.First(gitTagsThread, []string{})
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *gitTagsThreadRepository) First(gitTagsThread *model.GitTagsThread, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "First", e, now)
		}
	}()
	if len(gitTagsThread.Location()) == 0 {
		return errors.New("无更新字段！")
	}
	//先查询redis缓存
	repo.FindInRedis(gitTagsThread)
	if gitTagsThread != nil {
		return nil
	}
	db := global.DB.Table(gitTagsThread.TableName())
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db.First(gitTagsThread)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(gitTagsThread)
	return nil
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *gitTagsThreadRepository) DeleteByLocation(gitTagsThread *model.GitTagsThread) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(gitTagsThread.Location()) == 0 {
		return 0, errors.New("无更新字段！")
	}
	result := global.DB.Table(gitTagsThread.TableName()).Unscoped().Delete(gitTagsThread)
	e = result.Error
	if e != nil {
		return 0, e
	}
	repo.DeleteInRedis(gitTagsThread)
	return result.RowsAffected, nil
}

// 事务
func (repo *gitTagsThreadRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *gitTagsThreadRepository) SaveInRedis(gitTagsThread *model.GitTagsThread) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = gitTagsThread.RedisKey()
	resByte, e := json.Marshal(gitTagsThread)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *gitTagsThreadRepository) FindInRedis(gitTagsThread *model.GitTagsThread) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = gitTagsThread.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), gitTagsThread)
	}
	return nil
}

func (repo *gitTagsThreadRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *gitTagsThreadRepository) SaveInRedisByKey(redisKey string, data string) {
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
}

func (repo *gitTagsThreadRepository) DeleteInRedis(gitTagsThread *model.GitTagsThread) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = gitTagsThread.RedisKey()
	e = global.REDIS.Del(context.Background(), redisKey).Err()
	if e != nil {
		return e
	}
	return nil
}
func (repo *gitTagsThreadRepository) GetDataListByWhereMap(query map[string]interface{}, preload []string) (list []*model.GitTagsThread, e error) {
	now := time.Now()
	gitTagsThread := &model.GitTagsThread{}
	if query == nil {
		return nil, errors.New("无查询条件！")
	}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "DeleteByLocation", e, now)
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
	redisKey := gitTagsThread.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
	val, _ := repo.FindInRedisByKey(redisKey)
	db := global.DB.Model(gitTagsThread).Where(query)
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

func (repo *gitTagsThreadRepository) GetDataListByWhere(query string, args []interface{}, preload []string) (list []*model.GitTagsThread, e error) {
	now := time.Now()
	gitTagsThread := &model.GitTagsThread{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	db := global.DB.Model(gitTagsThread)
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
	redisKey := gitTagsThread.TableName() + "_list_" + strconv.Itoa(repo.Pager.Page) + "_" + strconv.Itoa(repo.Pager.PageSize) + "_" + str
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

func (repo *gitTagsThreadRepository) GetDataByWhereMap(gitTagsThread *model.GitTagsThread, where map[string]interface{}, preload []string) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(gitTagsThread.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	repo.FindInRedis(gitTagsThread)
	if gitTagsThread != nil {
		return nil
	}
	db := global.DB.Model(gitTagsThread).Where(where)
	if preload != nil {
		for _, v := range preload {
			db.Preload(v)
		}
	}
	db = db.First(gitTagsThread)
	e = db.Error
	if e != nil {
		return e
	}
	repo.SaveInRedis(gitTagsThread)
	return nil
}

func (repo *gitTagsThreadRepository) Execute(db *gorm.DB, object interface{}) error {
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

func (repo *gitTagsThreadRepository) GetTotalPage(db *gorm.DB) (e error) {
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
