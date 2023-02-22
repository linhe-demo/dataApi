package app

import (
	"context"
	"dataApi/conf"
	"dataApi/logs"
	"dataApi/pkg/bloomfilter"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

//全局资源
var (
	MysqlClient *gorm.DB
	RedisClient *redis.Client
	BloomClient *bloomfilter.BloomFilter
)

func Init() {
	conf.Init()
	logs.InitLogger()
	initMysqlDb()
	initRedis()
	initBloom()
}

func initMysqlDb() {
	var err error
	loglevel := logger2.Error
	if conf.AppConfig.Server.Debug {
		loglevel = logger2.Info
	}
	dblogger := logger2.New(log.New(logs.GetLogWriter(), "\r\n", log.LstdFlags), logger2.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      loglevel,
		Colorful:      true,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true",
		conf.AppConfig.Mysql.User, conf.AppConfig.Mysql.Password, conf.AppConfig.Mysql.Host, conf.AppConfig.Mysql.Port, conf.AppConfig.Mysql.DbName)

	gormConfig := gorm.Config{Logger: dblogger, SkipDefaultTransaction: true}
	gormConfig.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	MysqlClient, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       dsn,   // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
			DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // smart configure based on used version
		}),
		&gormConfig,
	)
	//MysqlClient, err = gorm.Open("mysql", uri)
	if err != nil {
		logs.Logger.Error("mysql open err", err)
		return
	}

	db, err := MysqlClient.DB()
	if err != nil {
		logs.Logger.Error("mysql db err", err)
		return
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(200)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(1000)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//todo 上线前检查数据库配置
	db.SetConnMaxLifetime(time.Second)
	//db.SetConnMaxLifetime(3600 * time.Second)
}

func initRedis() {
	addr := fmt.Sprintf("%s:%s", conf.AppConfig.Redis.Host, conf.AppConfig.Redis.Port)
	opt := &redis.Options{
		Addr:     addr,
		Password: conf.AppConfig.Redis.Password,
		PoolSize: 10,
		DB:       0,
	}
	RedisClient = redis.NewClient(opt)

	pong, err := RedisClient.Ping(context.TODO()).Result()
	logs.Logger.Infow("redis init:", "pong", pong, "err", err)
}

func initBloom() {
	BloomClient = bloomfilter.NewBloomFilter()
}
