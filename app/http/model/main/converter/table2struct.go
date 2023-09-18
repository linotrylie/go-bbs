package converter

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"strings"
)

// map for converting mysql type to golang types
var typeForMysqlToGo = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "string", // time.Time
	"datetime":           "string", // time.Time
	"timestamp":          "string", // time.Time
	"time":               "string", // time.Time
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

type Table2Struct struct {
	dsn            string
	savePath       string
	db             *sql.DB
	table          string
	prefix         string
	config         *T2tConfig
	err            error
	realNameMethod string
	enableJsonTag  bool   // 是否添加json的tag, 默认不添加
	packageName    string // 生成struct的包名(默认为空的话, 则取名为: package model)
	tagKey         string // tag字段的key值,默认是orm
}

type T2tConfig struct {
	RmTagIfUcFirsted bool // 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
	TagToLower       bool // tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
	UcFirstOnly      bool // 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
	SeperatFile      bool // 每个struct放入单独的文件,默认false,放入同一个文件
}

func NewTable2Struct() *Table2Struct {
	return &Table2Struct{}
}

func (t *Table2Struct) Dsn(d string) *Table2Struct {
	t.dsn = d
	return t
}

func (t *Table2Struct) TagKey(r string) *Table2Struct {
	t.tagKey = r
	return t
}

func (t *Table2Struct) PackageName(r string) *Table2Struct {
	t.packageName = r
	return t
}

func (t *Table2Struct) RealNameMethod(r string) *Table2Struct {
	t.realNameMethod = r
	return t
}

func (t *Table2Struct) SavePath(p string) *Table2Struct {
	t.savePath = p
	return t
}

func (t *Table2Struct) DB(d *sql.DB) *Table2Struct {
	t.db = d
	return t
}

func (t *Table2Struct) Table(tab string) *Table2Struct {
	t.table = tab
	return t
}

func (t *Table2Struct) Prefix(p string) *Table2Struct {
	t.prefix = p
	return t
}

func (t *Table2Struct) EnableJsonTag(p bool) *Table2Struct {
	t.enableJsonTag = p
	return t
}

func (t *Table2Struct) Config(c *T2tConfig) *Table2Struct {
	t.config = c
	return t
}

func (t *Table2Struct) RunRepository() error {
	formatStr := `
package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/redis/go-redis/v9"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type %sRepository struct {
	Pager *Pager
}
var %sRepository = new%sRepository()

func new%sRepository() *%sRepository {
	return new(%sRepository)
}

func (repo *%sRepository) Insert(%s *model.%s) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "Insert", e, now)
		}
	}()
	result := global.DB.Create(%s)
	if result.Error != nil {
		
		return
	}
	repo.SaveInRedis(%s)
	return result.RowsAffected, result.Error
}

func (repo *%sRepository) Update(%s *model.%s) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "Update", e, now)
		}
	}()
	if len(%s.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := %s.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}
	result := global.DB.Table(%s.TableName()).Where(%s.Location()).Updates(updateValues)
	if result.Error != nil {
		return
	}
	//更新完成后，重新缓存
	repo.DeleteInRedis(%s)
	repo.First(%s)
	e = result.Error
	rowsAffected = result.RowsAffected
	return
}

func (repo *%sRepository) First(%s *model.%s) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "First", e, now)
		}
	}()
	if len(%s.Location()) == 0 {
		return errors.New("location cannot be empty")
	}
	//先查询redis缓存
	err := repo.FindInRedis(%s)
	if err != nil && err != redis.Nil {
		return err
	}
	result := global.DB.Table(%s.TableName()).Where(%s.Location()).First(%s)
	e = result.Error
	if result.Error != nil {
		
		return
	}
	repo.SaveInRedis(%s)
	return
}

// DeleteByLocation 此方法为硬删除 慎用
func (repo *%sRepository) DeleteByLocation(%s *model.%s) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "DeleteByLocation", e, now)
		}
	}()
	if len(%s.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	result := global.DB.Table(%s.TableName()).Where(%s.Location()).Unscoped().Delete(%s)
	if result.Error != nil {
		return
	}
	repo.DeleteInRedis(%s)
	rowsAffected = result.RowsAffected
	e = result.Error
	return
}

// 事务
func (repo *%sRepository) TransactionExecute(fun func() error, opts ...*sql.TxOptions) (e error) {
	return global.DB.Transaction(func(tx *gorm.DB) (e error) {
		defer func() {
			if err := recover(); err != nil {
				e = errors.New(fmt.Sprint(err))
				global.LOG.Error(e.Error(), zap.Error(e))
			}
		}()
		e = fun()
		return
	}, opts...)
}

//////////////Redis///////////////////////////

func (repo *%sRepository) SaveInRedis(%s *model.%s) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = %s.RedisKey()
	resByte, e := json.Marshal(%s)
	if e != nil {
		return e
	}
	resStr := string(resByte)
	global.REDIS.Set(context.Background(), redisKey, resStr, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *%sRepository) FindInRedis(%s *model.%s) (e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = %s.RedisKey()
	redisRes, e := global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		e = json.Unmarshal([]byte(redisRes), %s)
	}
	return
}

func (repo *%sRepository) FindInRedisByKey(redisKey string) (redisRes string, e error) {
	defer func() {
		if e != nil && e != redis.Nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	redisRes, e = global.REDIS.Get(context.Background(), redisKey).Result()
	if e != nil && e != redis.Nil {
		return
	} else if e == redis.Nil {
		return
	} else {
		return
	}
	return
}

func (repo *%sRepository) SaveInRedisByKey(redisKey string, data string) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	global.REDIS.Set(context.Background(), redisKey, data, time.Duration(random.RandInt(7200, 14400))*time.Second)
	return nil
}

func (repo *%sRepository) DeleteInRedis(%s *model.%s) (e error) {
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
		}
	}()
	var redisKey string
	redisKey = %s.RedisKey()
	err := global.REDIS.Del(context.Background(), redisKey).Err()
	if err != nil {
		return
	}
	return nil
}
func (repo *%sRepository) GetDataListByWhereMap(query map[string]interface{}) (list []*model.%s, e error) {
	now := time.Now()
	%s := &model.%s{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "DeleteByLocation", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := %s.TableName() + "_list_"+ strconv.Itoa( repo.Pager.Page ) +"_"+ strconv.Itoa( repo.Pager.PageSize ) +"_" + str;
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(%s)
		if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
			var count64 int64
			e = db.Count(&count64).Error
			count := int(count64)
			if e != nil {
				return
			}
			if count != 0 {
				//Calculate the length of the pagination
				if count %` + ` repo.Pager.PageSize == 0 {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 0
				} else {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
				}
			}
		}
		return
	}
	db := global.DB.Table(%s.TableName()).Where(query)
	e = repo.Execute(db, &list)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, exceptions.NotFoundData
	}
	marshal, e := json.Marshal(list)
	if e != nil {
		return nil, e
	}
	repo.SaveInRedisByKey(redisKey, string(marshal))
	return
}

func (repo *%sRepository) GetDataListByWhere(query string, args []interface{}) (list []*model.%s, e error) {
	now := time.Now()
	%s := &model.%s{}
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "GetDataListByWhere", e, now)
		}
	}()
	var str string
	if repo.Pager.FieldsOrder != nil {
		for _, v := range repo.Pager.FieldsOrder {
			str += strings.Replace(v, " ", "", -1)
		}
	}
	redisKey := %s.TableName() + "_list_"+ strconv.Itoa( repo.Pager.Page ) +"_"+ strconv.Itoa( repo.Pager.PageSize ) +"_" + str;
	key, _ := repo.FindInRedisByKey(redisKey)
	if key != "" {
		e = json.Unmarshal([]byte(key), &list)
		if e != nil {
			return nil, e
		}
		db := global.DB.Model(%s)
		if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
			var count64 int64
			e = db.Count(&count64).Error
			count := int(count64)
			if e != nil {
				return
			}
			if count != 0 {
				//Calculate the length of the pagination
				if count %` + ` repo.Pager.PageSize == 0 {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 0
				} else {
					repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
				}
			}
		}
		return
	}
	db := global.DB.Table(%s.TableName())
	if query != "" {
		db = db.Where(query, args...)
	}
	e = repo.Execute(db, &list)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, exceptions.NotFoundData
	}
	marshal, e := json.Marshal(list)
	if e != nil {
		return nil, e
	}
	repo.SaveInRedisByKey(redisKey, string(marshal))
	return
}

func (repo *%sRepository) GetDataByWhereMap(%s *model.%s,where map[string]interface{}) (e error) {
	now := time.Now()
	defer func() {
		if e != nil {
			global.LOG.Error(e.Error(), zap.Error(e))
			global.Prome.OrmWithLabelValues(%s.TableName(), "GetDataByWhereMap", e, now)
		}
	}()
	db := global.DB.Table(%s.TableName()).Where(where).First(%s)
	e = db.Error
	if e != nil {
		return
	}
	repo.SaveInRedis(%s)
	return
}

func (repo *%sRepository) Execute(db *gorm.DB, object interface{}) error {
	if repo.Pager.Page != 0 && repo.Pager.PageSize != 0 {
		var count64 int64
		e := db.Count(&count64).Error
		count := int(count64)
		if e != nil {
			return e
		}
		if count != 0 {
			//Calculate the length of the pagination
			if count %` + ` repo.Pager.PageSize == 0 {
				repo.Pager.TotalPage = count / repo.Pager.PageSize
			} else {
				repo.Pager.TotalPage = count/repo.Pager.PageSize + 1
			}
		}
		db = db.Offset((repo.Pager.Page - 1) * repo.Pager.PageSize).Limit(repo.Pager.PageSize)
	}
	orderValue := repo.Pager.FieldsOrder
	if len(orderValue) > 0 {
		for _, v := range orderValue {
			db = db.Order(v)
		}
	}
	resultDB := db.Find(object)
	if resultDB.Error != nil {
		return resultDB.Error
	}
	return nil
}
`
	if t.config == nil {
		t.config = new(T2tConfig)
	}
	// 链接mysql, 获取db对象
	t.dialMysql()
	if t.err != nil {
		return t.err
	}
	// 获取表和字段的shcema
	tableColumns, err := t.getColumns()
	if err != nil {
		return err
	}
	// 组装struct
	for tableRealName, _ := range tableColumns {
		// 去除前缀
		if t.prefix != "" {
			tableRealName = tableRealName[len(t.prefix):]
		}
		tableName := strutil.CamelCase(tableRealName)
		switch len(tableName) {
		case 0:
		case 1:
			tableName = strings.ToUpper(tableName[0:1])
		default:
			// 字符长度大于1时
			tableName = strings.ToUpper(tableName[0:1]) + tableName[1:]
		}
		//fmt.Println(tableName)
		/*if tableRealName == "user" || tableRealName == "group" || tableRealName == "forum" || tableRealName == "thread" {
			continue
		}*/
		// newTableName forumAccess tableName ForumAccess tableRealName forum_access
		newTableName := strings.ToLower(tableName[0:1]) + tableName[1:]
		big := tableName
		small := newTableName
		sprintf := fmt.Sprintf(formatStr, newTableName, tableName, tableName, tableName, newTableName, newTableName,
			newTableName, newTableName, tableName, newTableName, newTableName, newTableName,
			newTableName, newTableName, tableName, newTableName, newTableName, newTableName, newTableName, newTableName, newTableName, newTableName,
			newTableName, newTableName, tableName, newTableName, newTableName, newTableName, newTableName, newTableName, newTableName, newTableName,
			newTableName, newTableName, tableName, newTableName, newTableName, newTableName, newTableName, newTableName, newTableName,
			newTableName, newTableName, small, tableName, newTableName, newTableName,
			newTableName, newTableName, big,
			newTableName, newTableName, small, small, small, small, tableName, newTableName, newTableName, tableName,
			newTableName /*63*/, tableName, newTableName, newTableName, newTableName, "", newTableName, newTableName, tableName,
			newTableName, tableName, newTableName, newTableName, newTableName, "" /*77*/, newTableName, newTableName,
			newTableName, tableName, newTableName, newTableName, newTableName, newTableName,
			newTableName, "",
		)
		// 写入文件struct
		var savePath = t.savePath
		// 是否指定保存路径
		//if savePath == "" {
		savePath += strings.ToLower(tableRealName) + "_repository.go"
		//}
		filePath := fmt.Sprintf("%s", savePath)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Can not write file")
			return err
		}
		defer f.Close()
		f.WriteString(sprintf)
		cmd := exec.Command("gofmt", "-w", filePath)
		cmd.Run()
	}
	return nil
}

func (t *Table2Struct) RunRequest() error {
	if t.config == nil {
		t.config = new(T2tConfig)
	}
	// 链接mysql, 获取db对象
	t.dialMysql()
	if t.err != nil {
		return t.err
	}

	// 获取表和字段的shcema
	tableColumns, err := t.getColumns()
	if err != nil {
		return err
	}

	//fmt.Println(tableColumns)

	// 包名
	var packageName string
	if t.packageName == "" {
		packageName = "package model\n\n"
	} else {
		packageName = fmt.Sprintf("package %s\n\n", t.packageName)
	}
	// 组装struct
	var structContent string
	for tableRealName, item := range tableColumns {
		// 去除前缀
		if t.prefix != "" {
			tableRealName = tableRealName[len(t.prefix):]
		}
		tableName := strutil.CamelCase(tableRealName)
		switch len(tableName) {
		case 0:
		case 1:
			tableName = strings.ToUpper(tableName[0:1])
		default:
			// 字符长度大于1时
			tableName = strings.ToUpper(tableName[0:1]) + tableName[1:]
		}
		depth := 1
		structContent += "type " + tableName + "Request struct {\n"
		for _, v := range item {
			v.Tag = "`" + ` json:"` + strings.ToLower(v.ColumnName) + `"` + "`"
			// 字段注释
			var clumnComment string
			if v.ColumnComment != "" {
				clumnComment = fmt.Sprintf(" // %s", v.ColumnComment)
			}
			structContent += fmt.Sprintf("%s%s %s %s%s\n",
				tab(depth), v.ColumnName, v.Type, v.Tag, clumnComment)
		}
		structContent += tab(depth-1) + "}\n\n"
		// 写入文件struct
		var savePath = t.savePath
		// 是否指定保存路径
		//if savePath == "" {
		savePath += strings.ToLower(tableName) + "_request.go"
		//}
		filePath := fmt.Sprintf("%s", savePath)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Can not write file")
			return err
		}
		defer f.Close()
		f.WriteString(packageName + structContent)
		cmd := exec.Command("gofmt", "-w", filePath)
		cmd.Run()
		structContent = ""
	}
	return nil
}

func (t *Table2Struct) RunEntity() error {
	if t.config == nil {
		t.config = new(T2tConfig)
	}
	// 链接mysql, 获取db对象
	t.dialMysql()
	if t.err != nil {
		return t.err
	}

	// 获取表和字段的shcema
	tableColumns, err := t.getColumns()
	if err != nil {
		return err
	}

	//fmt.Println(tableColumns)

	// 包名
	var packageName string
	if t.packageName == "" {
		packageName = "package entity\n\n"
	} else {
		packageName = fmt.Sprintf("package %s\n\n", t.packageName)
	}
	packageName += `import "go-bbs/app/http/model"`
	packageName += "\n"
	// 组装struct
	var structContent string
	for tableRealName, _ := range tableColumns {
		// 去除前缀
		if t.prefix != "" {
			tableRealName = tableRealName[len(t.prefix):]
		}
		tableName := strutil.CamelCase(tableRealName)
		switch len(tableName) {
		case 0:
		case 1:
			tableName = strings.ToUpper(tableName[0:1])
		default:
			// 字符长度大于1时
			tableName = strings.ToUpper(tableName[0:1]) + tableName[1:]
		}
		structContent += "type " + tableName + "Entity struct {\n"
		structContent += fmt.Sprintf("*model.%s\n", tableName)
		structContent += "}"
		// 写入文件struct
		var savePath = t.savePath
		// 是否指定保存路径
		//if savePath == "" {
		savePath += strings.ToLower(tableRealName) + "_entity.go"
		//}
		filePath := fmt.Sprintf("%s", savePath)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Can not write file")
			return err
		}
		defer f.Close()
		f.WriteString(packageName + structContent)
		cmd := exec.Command("gofmt", "-w", filePath)
		cmd.Run()
		structContent = ""
	}
	return nil
}

func (t *Table2Struct) Run() error {
	if t.config == nil {
		t.config = new(T2tConfig)
	}
	// 链接mysql, 获取db对象
	t.dialMysql()
	if t.err != nil {
		return t.err
	}

	// 获取表和字段的shcema
	tableColumns, err := t.getColumns()
	if err != nil {
		return err
	}

	//fmt.Println(tableColumns)

	// 包名
	var packageName string
	if t.packageName == "" {
		packageName = "package model\n\n"
	} else {
		packageName = fmt.Sprintf("package %s\n\n", t.packageName)
	}

	// 组装struct
	var structContent string
	for tableRealName, item := range tableColumns {
		// 去除前缀
		if t.prefix != "" {
			tableRealName = tableRealName[len(t.prefix):]
		}
		tableName := strutil.CamelCase(tableRealName)
		primaryStructMap := map[string]string{}
		switch len(tableName) {
		case 0:
		case 1:
			tableName = strings.ToUpper(tableName[0:1])
		default:
			// 字符长度大于1时
			tableName = strings.ToUpper(tableName[0:1]) + tableName[1:]
		}
		depth := 1
		structContent += "type " + tableName + " struct {\n"
		structContent += "changes   map[string]interface{}\n"
		for _, v := range item {
			//structContent += tab(depth) + v.ColumnName + " " + v.Type + " " + v.Json + "\n"
			//column := v.Tag
			if v.Primary == "PRI" {
				primaryStructMap[v.ColumnName] = v.ColumnName
				v.Tag = "`" + `gorm:"primaryKey;column:` + strings.ToLower(v.Tag) + `"` + ` json:"` + strings.ToLower(v.Tag) + `"` + "`"
			} else {
				v.Tag = "`" + `gorm:"column:` + strings.ToLower(v.Tag) + `"` + ` json:"` + strings.ToLower(v.Tag) + `"` + "`"
			}
			// 字段注释
			var clumnComment string
			if v.ColumnComment != "" {
				clumnComment = fmt.Sprintf(" // %s", v.ColumnComment)
			}
			structContent += fmt.Sprintf("%s%s %s %s%s\n",
				tab(depth), v.ColumnName, v.Type, v.Tag, clumnComment)
		}
		structContent += tab(depth-1) + "}\n\n"

		// 添加 method 获取真实表名
		if t.realNameMethod != "" {
			structContent += fmt.Sprintf("func (*%s) %s() string {\n",
				tableName, t.realNameMethod)
			structContent += fmt.Sprintf("%sreturn \"%s%s\"\n",
				tab(depth), t.prefix, tableRealName)
			structContent += "}\n\n"
		}

		structContent += "// Location" + " .\n"
		structContent += fmt.Sprintf("func (obj *%s) Location() map[string]interface{} {\n", tableName)

		locationStr := ""
		for k, v := range primaryStructMap {
			if locationStr != "" {
				locationStr += ","
			}
			locationStr += fmt.Sprintf(`"%s":obj.%s`, strings.ToLower(k), v)
		}
		structContent += fmt.Sprintf(`return map[string]interface{}{%s}`+"\n", locationStr)
		structContent += "}\n"

		structContent += "// Redis Key .\n"
		structContent += fmt.Sprintf("func (obj *%s) RedisKey() string {\n", tableName)
		redisKeyStr := " obj.TableName() "
		if len(primaryStructMap) == 0 {
			redisKeyStr += ` + "_" + ` + `fmt.Sprintf("%v",time.Now().Unix())`
		} else {
			for _, v := range primaryStructMap {
				if redisKeyStr != "" {
					redisKeyStr += ""
				}
				redisKeyStr += ` + "_" + ` + `fmt.Sprintf("%v"` + fmt.Sprintf(`,obj.%s)`, v)
			}
		}
		redisKeyStr += "\n"
		structContent += fmt.Sprintf("return %s", redisKeyStr)
		structContent += "}\n"

		structContent += fmt.Sprintf("// GetChanges .\nfunc (obj *%s) GetChanges() map[string]interface{} {\n\tif obj.changes == nil {\n\t\treturn nil\n\t}\n\tresult := make(map[string]interface{})\n\tfor k, v := range obj.changes {\n\t\tresult[k] = v\n\t}\n\tobj.changes = nil\n\treturn result\n}\n\n// Update .\nfunc (obj *%s) Update(name string, value interface{}) {\n\tif obj.changes == nil {\n\t\tobj.changes = make(map[string]interface{})\n\t}\n\tobj.changes[name] = value\n}", tableName, tableName)
		fmt.Println(structContent)
		//return nil
		// 如果有引入 time.Time, 则需要引入 time 包
		var importContent string

		if strings.Contains(structContent, "time.Time") {
			importContent = "import (\"fmt\"\n\"time\"\n)\n"
		} else if strings.Contains(structContent, "time.Now()") {
			importContent = "import (\"fmt\"\n\"time\"\n)\n"
		} else {
			importContent = "import (\"fmt\"\n)\n"
		}

		// 写入文件struct
		var savePath = t.savePath
		// 是否指定保存路径
		//if savePath == "" {
		savePath += strings.ToLower(tableRealName) + "_model.go"
		//}
		filePath := fmt.Sprintf("%s", savePath)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Can not write file")
			return err
		}
		defer f.Close()
		f.WriteString(packageName + importContent + structContent)
		cmd := exec.Command("gofmt", "-w", filePath)
		cmd.Run()
		structContent = ""
		importContent = ""
	}

	/*// 如果有引入 time.Time, 则需要引入 time 包
	var importContent string
	if strings.Contains(structContent, "time.Time") {
		importContent = "import \"time\"\n\n"
	}

	// 写入文件struct
	var savePath = t.savePath
	// 是否指定保存路径
	if savePath == "" {
		savePath = "model.go"
	}
	filePath := fmt.Sprintf("%s", savePath)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Can not write file")
		return err
	}
	defer f.Close()

	f.WriteString(packageName + importContent + structContent)

	cmd := exec.Command("gofmt", "-w", filePath)
	cmd.Run()*/

	return nil
}

func (t *Table2Struct) dialMysql() {
	if t.db == nil {
		if t.dsn == "" {
			t.err = errors.New("dsn数据库配置缺失")
			return
		}
		t.db, t.err = sql.Open("mysql", t.dsn)
	}
	return
}

type column struct {
	ColumnName    string
	Type          string
	Nullable      string
	TableName     string
	ColumnComment string
	Tag           string
	Primary       string
}

// Function for fetching schema definition of passed table
func (t *Table2Struct) getColumns(table ...string) (tableColumns map[string][]column, err error) {
	tableColumns = make(map[string][]column)
	// sql
	var sqlStr = `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT,COLUMN_KEY
		FROM information_schema.COLUMNS 
		WHERE table_schema = DATABASE()`
	// 是否指定了具体的table
	if t.table != "" {
		sqlStr += fmt.Sprintf(" AND TABLE_NAME = '%s'", t.prefix+t.table)
	}
	// sql排序
	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"

	rows, err := t.db.Query(sqlStr)
	if err != nil {
		fmt.Println("Error reading table information: ", err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		col := column{}
		err = rows.Scan(&col.ColumnName, &col.Type, &col.Nullable, &col.TableName, &col.ColumnComment, &col.Primary)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//col.Json = strings.ToLower(col.ColumnName)
		col.Tag = col.ColumnName
		col.ColumnComment = col.ColumnComment
		col.ColumnName = t.camelCase(col.ColumnName)
		col.Type = typeForMysqlToGo[col.Type]
		// 字段首字母本身大写, 是否需要删除tag
		if t.config.RmTagIfUcFirsted &&
			col.ColumnName[0:1] == strings.ToUpper(col.ColumnName[0:1]) {
			col.Tag = "-"
		} else {
			// 是否需要将tag转换成小写
			if t.config.TagToLower {
				col.Tag = strings.ToLower(col.Tag)
			}
			//if col.Nullable == "YES" {
			//	col.Json = fmt.Sprintf("`json:\"%s,omitempty\"`", col.Json)
			//} else {
			//}
		}
		/*if t.enableJsonTag {
			//col.Json = fmt.Sprintf("`json:\"%s\" %s:\"%s\"`", col.Json, t.config.TagKey, col.Json)
			col.Tag = fmt.Sprintf("`%s:\"%s\" json:\"%s\"`", t.tagKey, col.Tag, col.Tag)
		} else {
			col.Tag = fmt.Sprintf("`%s:\"%s\"`", t.tagKey, col.Tag)
		}*/
		//columns = append(columns, col)
		if _, ok := tableColumns[col.TableName]; !ok {
			tableColumns[col.TableName] = []column{}
		}
		tableColumns[col.TableName] = append(tableColumns[col.TableName], col)
	}
	return
}

func (t *Table2Struct) camelCase(str string) string {
	// 是否有表前缀, 设置了就先去除表前缀
	if t.prefix != "" {
		str = strings.Replace(str, t.prefix, "", 1)
	}
	var text string
	//for _, p := range strings.Split(name, "_") {
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			// 字符长度大于1时
			if t.config.UcFirstOnly == true {
				text += strings.ToUpper(p[0:1]) + strings.ToLower(p[1:])
			} else {
				text += strings.ToUpper(p[0:1]) + p[1:]
			}
		}
	}
	return text
}
func tab(depth int) string {
	return strings.Repeat("\t", depth)
}
