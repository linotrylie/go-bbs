package elasticsearch

import (
	"github.com/gin-gonic/gin"
	"go-bbs/plugin/elasticsearch/global"
	"go-bbs/plugin/elasticsearch/router"
)

type ESPlugin struct {
}

func CreateESPlugin(Host, Port, User, Password string) *ESPlugin {
	global.ElasticSearch.User = User
	global.ElasticSearch.Host = Host
	global.ElasticSearch.Port = Port
	global.ElasticSearch.Password = Password
	return &ESPlugin{}
}
func (*ESPlugin) Register(group *gin.RouterGroup) {
	router.ESRouterGroupApp.InitEsRouter(group)
}

func (*ESPlugin) RouterPath() string {
	return "es"
}
