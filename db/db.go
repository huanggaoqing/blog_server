package db

import (
	"blog_server/module/dbModule"
	"blog_server/tools"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Person struct {
	name string
}

var dataBase *gorm.DB

// InitDb 初始化数据库
func InitDb() (*gorm.DB, error) {
	config := tools.GetSysConfig()
	dbConfig := config.DataBase
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "blog_",
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}

	// 设置数据库连接池
	poolConfig := config.DataPool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(poolConfig.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(poolConfig.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(poolConfig.MaxLifetime) * time.Hour)

	// 自动表映射
	err = db.AutoMigrate(
		&dbModule.User{},
		&dbModule.Role{},
		&dbModule.Profile{},
		&dbModule.Article{},
		&dbModule.Tag{},
		&dbModule.TagAssociate{},
	)
	if err != nil {
		return nil, err
	}

	dataBase = db
	return dataBase, nil
}

func GetDataBase() *gorm.DB {
	return dataBase
}
