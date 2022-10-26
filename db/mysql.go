package db

import (
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/pkg/constant"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

)

var MysqlClient *gorm.DB

func InitMysql(env string) error {
	client, err := gorm.Open(mysql.Open(config.Cfg.Mysql.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	if env != constant.Prod {
		client = client.Debug()
	}

	MysqlClient = client

	sqlClient, err := MysqlClient.DB()
	if err != nil {
		return err
	}

	sqlClient.SetConnMaxLifetime(time.Hour)
	sqlClient.SetMaxOpenConns(config.Cfg.Mysql.MaxOpenConns)
	sqlClient.SetMaxIdleConns(config.Cfg.Mysql.MaxIdleConns)

	return nil
}
