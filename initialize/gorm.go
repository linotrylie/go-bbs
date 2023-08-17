package initialize

import (
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}
