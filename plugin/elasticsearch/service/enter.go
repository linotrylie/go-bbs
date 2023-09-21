package service

type ESServiceGroup struct {
	ElasticsearchService
}

var ServiceGroupApp = new(ESServiceGroup)
