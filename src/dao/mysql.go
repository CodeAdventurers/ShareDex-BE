package dao

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/goccy/go-json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBManager struct {
	Conn *gorm.DB
}

var dbManager *DBManager
var once sync.Once

func newDBManager() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ShareDex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		log.Default().Println("数据库连接错误", err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Default().Println("构造数据库连接池错误", err.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbManager = &DBManager{Conn: db}
	data, _ := json.Marshal(sqlDB.Stats()) //获得当前的SQL配置情况
	fmt.Println(string(data))
}

func NewDBManager() *DBManager {
	once.Do(newDBManager)
	return dbManager
}
