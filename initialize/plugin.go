package initialize

import (
	"github.com/gin-gonic/gin"
	"go-bbs/global"
	"go-bbs/plugin/elasticsearch"
	"go-bbs/plugin/elasticsearch/initialize"
	"go-bbs/plugin/email"
	"go-bbs/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")

	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	EmailPlugin := email.CreateEmailPlug(
		global.CONFIG.Email.To,
		global.CONFIG.Email.From,
		global.CONFIG.Email.Host,
		global.CONFIG.Email.Secret,
		global.CONFIG.Email.Nickname,
		global.CONFIG.Email.Port,
		global.CONFIG.Email.IsSSL,
	)

	ESPlugin := elasticsearch.CreateESPlugin(
		global.CONFIG.ElasticSearch.Host,
		global.CONFIG.ElasticSearch.Port,
		global.CONFIG.ElasticSearch.User,
		global.CONFIG.ElasticSearch.Password,
	)

	PluginInit(PublicGroup, EmailPlugin, ESPlugin)
	//连接ES服务器
	if global.CONFIG.ElasticSearch.Enable {
		initialize.ConnectElasticsearch()
	}
}
