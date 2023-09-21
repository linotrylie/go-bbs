package service

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-bbs/plugin/elasticsearch/global"
	"go-bbs/plugin/elasticsearch/model/response"
)

type ElasticsearchService struct{}

// CreateElasticsearch 创建索引
//
//	@return type IndexResponse struct {
//		Index         string      `json:"_index,omitempty"`
//		Type          string      `json:"_type,omitempty"`
//		Id            string      `json:"_id,omitempty"`
//		Version       int64       `json:"_version,omitempty"`
//		Result        string      `json:"result,omitempty"`
//		Shards        *ShardsInfo `json:"_shards,omitempty"`
//		SeqNo         int64       `json:"_seq_no,omitempty"`
//		PrimaryTerm   int64       `json:"_primary_term,omitempty"`
//		Status        int         `json:"status,omitempty"`
//		ForcedRefresh bool        `json:"forced_refresh,omitempty"`
//	}
func (e *ElasticsearchService) CreateElasticsearch(name string, esId string, data interface{}) (err error, ret interface{}) {
	get, err := global.ESClient.Index().Index(name).Id(esId).BodyJson(data).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

// DeleteElasticsearch 删除
func (e *ElasticsearchService) DeleteElasticsearch(name string, esId string) (err error, ret interface{}) {

	get, err := global.ESClient.Delete().Index(name).Id(esId).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

// UpdateElasticsearch 更新
func (e *ElasticsearchService) UpdateElasticsearch(name string, esId string, data interface{}) (err error, ret interface{}) {

	get, err := global.ESClient.Update().Index(name).Id(esId).Doc(data).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

// GetIdElasticsearch 通过id查找Elasticsearch
func (e *ElasticsearchService) GetIdElasticsearch(name string, esId string) (err error, ret interface{}) {
	get, err := global.ESClient.Get().Index(name).Id(esId).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

// GetCountElasticsearch 当前ID索引数量
func (e *ElasticsearchService) GetCountElasticsearch(name string) (ret interface{}) {
	list, err := global.ESClient.Count(name).Do(context.Background())
	if err != nil {
		return err
	}
	return list
}

// GetQueryElasticsearch 查询Elasticsearch
func (e *ElasticsearchService) GetQueryElasticsearch(info response.ElasticSearchSearch, name, searchField string) (err error, ret interface{}) {
	size := info.PageSize
	page := info.Page
	//根据name索引查询Elasticsearch数据
	boolQ := elastic.NewQueryStringQuery(info.Title)
	boolQ = boolQ.Field(searchField)
	get, err := global.ESClient.Search(name).
		Query(boolQ). // specify the query
		//Sort("id", true). //按字段"age"排序，升序排列
		Size(size). // 分页，单页显示10条
		From((page - 1) * size).
		//FetchSourceContext(fsc).//只取对应字段
		Do(context.Background()) // 执行
	if err != nil {
		return err, ""
	}
	return err, get
}

func GetQueryMultipleElasticsearch(info response.ElasticSearchSearch, name string, fields ...string) (err error, ret interface{}) {
	size := info.PageSize
	page := info.Page
	if info.Type == 1 {
		//Elasticsearch精准搜索
		//高亮加粗
		hig := elastic.NewHighlight()
		hig = hig.Field("barCode")
		hig = hig.PreTags("<font color='red'>")
		hig = hig.PostTags("</font>")
		get, err := global.ESClient.Search(name).
			Query(elastic.NewMatchPhraseQuery("barCode", info.Title)).
			Highlight(hig).
			Size(size). // 分页，单页显示10条
			From((page - 1) * size).
			Do(context.Background())
		if err != nil {
			return err, ""
		}
		return err, get

	} else {
		// 模糊搜索多字段-需要用分词器analysis-ik
		if len(fields) > 0 {
			get, err := global.ESClient.Search(name).
				Query(elastic.NewBoolQuery().Should(elastic.NewMultiMatchQuery(info.Title, fields...).
					Fuzziness("AUTO")).MinimumShouldMatch("1")).
				From((page - 1) * size).Size(size).Do(context.Background())
			if err != nil {
				return err, ""
			}
			return err, get
		}
	}
	return err, nil
}
