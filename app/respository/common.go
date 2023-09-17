package respository

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
	"strings"
	"time"
)

type RepositoryInterface interface {
	Insert(model model.Model) (rowsAffected int64, e error)
	Update(model model.Model) (rowsAffected int64, e error)
	FindByLocation(model model.Model) (e error)
	DeleteByLocation(model model.Model) (rowsAffected int64, e error)
	TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error)
	SaveInRedis(model model.Model) (e error)
	FindInRedis(model model.Model) (e error)
	DeleteInRedis(model model.Model) (e error)
	SaveInRedisByKey(redisKey string, data string) (e error)
	FindInRedisByKey(redisKey string) (redisRes string, e error)
	GetDataByWhereMap(where map[string]interface{}) (e error)
	GetDataListByWhereMap(query string, args []interface{}) (list []model.User, e error)
}

type Repository struct {
	Model model.Model
	Pager *Pager
}

type Pager struct {
	PageSize    int
	Page        int
	TotalPage   int
	FieldsOrder []string // []{"id desc","name asc"}
}

func (repo *Repository) Insert(model model.Model) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.Prome.OrmWithLabelValues(model.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(model)
	if result.Error != nil {
		global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	repo.SaveInRedis(model)
	return result.RowsAffected, result.Error
}

func (repo *Repository) Update(model model.Model) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.Prome.OrmWithLabelValues(model.TableName(), "Update", e, now)
		}
	}()
	if len(model.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := model.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(model.TableName()).Where(model.Location()).Updates(updateValues)
	if result.Error != nil {
		//global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(model)
	repo.FindByLocation(model)
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *Repository) FindByLocation(model model.Model) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.Prome.OrmWithLabelValues(model.TableName(), "FindByLocation", e, now)
		}
	}()
	if len(model.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	err := repo.FindInRedis(model)
	if err != nil && err != redis.Nil {
		return err
	}
	result := global.DB.Table(model.TableName()).Where(model.Location()).First(model)
	e = result.Error
	if result.Error != nil {
		fmt.Println(model.TableName(), model.Location())
		global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	repo.SaveInRedis(model)
	return
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *Repository) DeleteByLocation(model model.Model) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.Prome.OrmWithLabelValues(model.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(model.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(model.TableName()).Where(model.Location()).Unscoped().Delete(model)
	if result.Error != nil {
		global.LOG.Error(result.Error.Error(), zap.Error(result.Error))
		return
	}
	repo.DeleteInRedis(model)
	rowsAffected = result.RowsAffected
	e = result.Error
	return
}

// 事务
func (repo *Repository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
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

func (repo *Repository) SaveInRedis(model model.Model) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = model.RedisKey()
	resByte, e := json.Marshal(model)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 999999))*time.Second)
	return nil
}

func (repo *Repository) FindInRedis(model model.Model) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = model.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), model)
	}
	return
}

func (repo *Repository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
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

func (repo *Repository) SaveInRedisByKey(redisKey string, data string) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *Repository) DeleteInRedis(model model.Model) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = model.RedisKey()
	err := global.REDIS.Del(context.Background(), redisKey).Err()
	if err != nil {
		return
	}
	return nil
}

func (repo *Repository) GetDataListByWhereMap(where map[string]interface{}) (list []model.Model, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(repo.Model.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	for _, v := range repo.Pager.FieldsOrder {
		str += strings.Replace(v, " ", "", -1)
	}
	redisKey := fmt.Sprintf("%s_list_%d_%d_%s", repo.Model.TableName(), repo.Pager.Page, repo.Pager.PageSize, str)
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(repo.Model)
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
	db := global.DB.Table(repo.Model.TableName())
	if where != nil {
		db = db.Where(where)
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

func (repo *Repository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(repo.Model.TableName(), "DeleteByLocation", e, now)
		}
	}()
	db := global.DB.Table(repo.Model.TableName()).Where(where).First(repo.Model)
	e = db.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(repo.Model)
	return
}

func (repo *Repository) Execute(db *gorm.DB, object interface{}) (e error) {
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
	return
}
