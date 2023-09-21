package global

import (
	"github.com/olivere/elastic/v7"
	"go-bbs/plugin/elasticsearch/config"
)

var (
	ElasticSearch = new(config.Elasticsearch)
	ESClient      *elastic.Client
)
