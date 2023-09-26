package initialize

import (
	"fmt"
	"go-bbs/config"
	"go-bbs/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GormMysql() *gorm.DB {
	databaseSetting := &global.CONFIG
	db, err := NewDBEngine(databaseSetting)
	if err != nil {
		return nil
	}
	return db
}

// NewDBEngine 实例化数据库连接
func NewDBEngine(server *config.Server) (*gorm.DB, error) {
	ds := server.Mysql
	logLevel := logger.Info
	/*	if utils.RunModeIsDebug() {
		logLevel = logger.Info
	}*/
	// sql 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标,前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	// 主库
	dbMasterDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		ds.Master.UserName,
		ds.Master.Password,
		ds.Master.Host,
		ds.Master.DBName,
	)
	// 从库
	/*	dbSlaveDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		ds.Slave.UserName,
		ds.Slave.Password,
		ds.Slave.Host,
		ds.Slave.DBName,
	)*/

	db, err := gorm.Open(mysql.Open(dbMasterDsn), &gorm.Config{
		Logger:                 newLogger.LogMode(logLevel),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	/*	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dbMasterDsn)},
		Replicas: []gorm.Dialector{mysql.Open(dbSlaveDsn)},
		Policy:   dbresolver.RandomPolicy{},
	}))*/

	/*if err != nil {
		return nil, err
	}*/
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(ds.Master.MaxIdleConns)
	sqlDb.SetMaxOpenConns(ds.Master.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Second * 600)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	return db, nil
}
