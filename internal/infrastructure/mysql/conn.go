package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hmdp/internal/domain/entity"
	logger2 "hmdp/pkg/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// InitDB  在中间件中初始化mysql链接
func InitDB() *gorm.DB {

	// 从配置文件中获取 MySQL 相关信息
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, database)
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		logger2.Logger.Error(fmt.Sprintf("mysql lost: %v", err))
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger2.Logger.Error(fmt.Sprintf("mysql lost: %v", err))
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	DB = db

	migration()
	return DB
}

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(
		&entity.User{},
		&entity.ShowType{},
		&entity.Blog{},
		&entity.Shop{},
		&entity.Voucher{},
	)
}

// GetDB 单例模式建立数据库连接
//func GetDB() *gorm.DB {
//	if globalDB == nil {
//		InitDB()
//	}
//	return globalDB
//}
