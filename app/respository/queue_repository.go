package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type queueRepository struct {
	Queue *model.Queue
	Pager *Pager
	Repo  Repository
}

var QueueRepository = newQueueRepository()

func newQueueRepository() *queueRepository {
	return new(queueRepository)
}

func (obj *queueRepository) Insert(queue model.Queue) (rowsAffected int64, e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.Insert(&queue)
}

func (obj *queueRepository) Update(queue model.Queue) (rowsAffected int64, e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.Update(&queue)
}

func (obj *queueRepository) FindByLocation(queue model.Queue) (e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.FindByLocation(&queue)
}

func (obj *queueRepository) DeleteByLocation(queue model.Queue) (rowsAffected int64, e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.Update(&queue)
}

func (obj *queueRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	QueueRepository.Repo.Model = &model.Queue{}
	return QueueRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *queueRepository) SaveInRedis(queue model.Queue) (e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.SaveInRedis(&queue)
}

func (obj *queueRepository) FindInRedis(queue model.Queue) (e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.FindInRedis(&queue)
}

func (obj *queueRepository) DeleteInRedis(queue model.Queue) (e error) {
	QueueRepository.Repo.Model = &queue
	return QueueRepository.Repo.DeleteInRedis(&queue)
}

func (obj *queueRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	QueueRepository.Repo.Model = &model.Queue{}
	return QueueRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *queueRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	QueueRepository.Repo.Model = &model.Queue{}
	return QueueRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *queueRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	QueueRepository.Repo.Model = &model.Queue{}
	return QueueRepository.Repo.GetDataByWhereMap(where)
}

func (obj *queueRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	QueueRepository.Repo.Model = &model.Queue{}
	return QueueRepository.Repo.GetDataListByWhereMap(where)
}
