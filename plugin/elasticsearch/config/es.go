package config

type Elasticsearch struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 哪个数据库
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 服务器地址:端口
	User     string `mapstructure:"user" json:"user" yaml:"user"`             // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Enable   bool   `mapstructure:"enable" json:"enable" yaml:"enable"`       // 是否启用
}
