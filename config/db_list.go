package config

type DsnProvider interface {
	Dsn() string
}

// Embeded 结构体可以压平到上一层，从而保持 config 文件的结构和原来一样
// 见 playground: https://go.dev/play/p/KIcuhqEoxmY

// GeneralDB 也被 Pgsql 和 Mysql 原样使用

type GeneralDB struct {
	DBType       string `mapstructure:"DBType" json:"DBType" yaml:"DBType"`
	UserName     string `mapstructure:"UserName" json:"UserName" yaml:"UserName"`
	Password     string `mapstructure:"Password" json:"Password" yaml:"Password"`
	Host         string `mapstructure:"Host" json:"Host" yaml:"Host"`
	DBName       string `mapstructure:"DBName" json:"DBName" yaml:"DBName"`
	TablePrefix  string `mapstructure:"TablePrefix" json:"TablePrefix" yaml:"TablePrefix"`
	Charset      string `mapstructure:"Charset" json:"Charset" yaml:"Charset"`
	ParseTime    bool   `mapstructure:"ParseTime" json:"ParseTime" yaml:"ParseTime"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns" json:"MaxIdleConns" yaml:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns" json:"MaxOpenConns" yaml:"MaxOpenConns"`
}

type Database struct {
	Master GeneralDB `mapstructure:"Master" json:"Master" yaml:"Master"`
	Slave  GeneralDB `mapstructure:"Slave" json:"Slave" yaml:"Slave"`
	LogZap bool      `mapstructure:"LogZap" json:"LogZap" yaml:"LogZap"`
}

func (d Database) GetLogMode() string {
	return "error"
}
