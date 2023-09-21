package initialize

import (
	"github.com/olivere/elastic/v7"
	"go-bbs/global"
	es_global "go-bbs/plugin/elasticsearch/global"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func ConnectElasticsearch() {
	elasticsearch()
}

// 初始化elasticsearch
func elasticsearch() {
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(es_global.ElasticSearch.Host + ":" + es_global.ElasticSearch.Port), // 设置elasticsearch地址
		elastic.SetSniff(false),                                                  // 关闭嗅探
		elastic.SetHealthcheckInterval(20 * time.Second),                         //每20秒检查一次
		elastic.SetGzip(false),                                                   // 关闭Gzip压缩
		elastic.SetErrorLog(log.New(os.Stderr, "elastic_search", log.LstdFlags)), // Error日志
		elastic.SetInfoLog(log.New(os.Stderr, "elastic_search", log.LstdFlags)),
	}
	var err error
	es_global.ESClient, err = elastic.NewClient(options...)
	if err != nil {
		global.LOG.Error("elasticsearch 连接失败", zap.Error(err))
		//panic("elasticsearch 连接失败：" + err.Error())
	}
}
