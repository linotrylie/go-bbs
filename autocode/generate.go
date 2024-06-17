package main

import (
	"go-bbs/autocode/converter"
)

func main() {
	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		SeperatFile: true,
	})
	// 开始迁移转换
	//生成model
	GenerateModel(t2t)
	//生成entity
	//GenerateEntity(t2t)
	//生成repository
	//GenerateRepository(t2t)
	//生成requests
	//GenerateRequests(t2t)
}

func GenerateModel(t2t *converter.Table2Struct) error {
	return t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table("kadao_data").
		// 表前缀
		Prefix("bbs_").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("model").
		// tag字段的key值,默认是orm
		TagKey("gorm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath("./app/http/model/").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn("root:root@tcp(localhost:3306)/freebns?charset=utf8").
		// 执行
		Run()
}

func GenerateEntity(t2t *converter.Table2Struct) error {
	return t2t.
		Table("operation_log").
		Prefix("bbs_").PackageName("entity").RealNameMethod("TableName").SavePath("./app/entity/").
		Dsn("root:root@tcp(localhost:3306)/freebns?charset=utf8").
		RunEntity()
}

func GenerateRepository(t2t *converter.Table2Struct) error {
	return t2t.
		//Table("*").
		Prefix("bbs_").PackageName("repository").RealNameMethod("TableName").SavePath("./app/repository/").
		Dsn("root:root@tcp(localhost:3306)/freebns?charset=utf8").
		RunRepository()
}

func GenerateRequests(t2t *converter.Table2Struct) error {
	return t2t.
		Table("kadao_data").
		Prefix("bbs_").PackageName("requests").RealNameMethod("TableName").SavePath("./app./http/model/requests/").
		Dsn("root:root@tcp(localhost:3306)/freebns?charset=utf8").
		RunRequest()
}
