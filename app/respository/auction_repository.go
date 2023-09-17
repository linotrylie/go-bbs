package respository

import (
	"database/sql"
	"go-bbs/app/http/model"
)

type auctionRepository struct {
	Auction *model.Auction
	Pager   *Pager
	Repo    Repository
}

var AuctionRepository = newAuctionRepository()

func newAuctionRepository() *auctionRepository {
	return new(auctionRepository)
}

func (obj *auctionRepository) Insert(auction model.Auction) (rowsAffected int64, e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.Insert(&auction)
}

func (obj *auctionRepository) Update(auction model.Auction) (rowsAffected int64, e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.Update(&auction)
}

func (obj *auctionRepository) FindByLocation(auction model.Auction) (e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.FindByLocation(&auction)
}

func (obj *auctionRepository) DeleteByLocation(auction model.Auction) (rowsAffected int64, e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.Update(&auction)
}

func (obj *auctionRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	AuctionRepository.Repo.Model = &model.Auction{}
	return AuctionRepository.Repo.TransactionExecute(fun, opts...)
}

func (obj *auctionRepository) SaveInRedis(auction model.Auction) (e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.SaveInRedis(&auction)
}

func (obj *auctionRepository) FindInRedis(auction model.Auction) (e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.FindInRedis(&auction)
}

func (obj *auctionRepository) DeleteInRedis(auction model.Auction) (e error) {
	AuctionRepository.Repo.Model = &auction
	return AuctionRepository.Repo.DeleteInRedis(&auction)
}

func (obj *auctionRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	AuctionRepository.Repo.Model = &model.Auction{}
	return AuctionRepository.Repo.SaveInRedisByKey(redisKey, data)
}

func (obj *auctionRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	AuctionRepository.Repo.Model = &model.Auction{}
	return AuctionRepository.Repo.FindInRedisByKey(redisKey)
}

func (obj *auctionRepository) GetDataByWhereMap(where map[string]interface{}) (e error) {
	AuctionRepository.Repo.Model = &model.Auction{}
	return AuctionRepository.Repo.GetDataByWhereMap(where)
}

func (obj *auctionRepository) GetDataListByWhereMap(where map[string]interface{}) ([]model.Model, error) {
	AuctionRepository.Repo.Model = &model.Auction{}
	return AuctionRepository.Repo.GetDataListByWhereMap(where)
}
