package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"
)

type Options struct {
	Host string
	Username string
	Password string
	Database string
	MaxIdleConnections int
	MaxOpenConnections int
	MaxConnectionLifeTime time.Duration
	LogLevel int
}

func New(opts *Options) (*gorm.DB,error)  {
	dsn:=fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,opts.Username,opts.Password,opts.Host,opts.Database,true,"Local")

	db,err:=gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDB,err:=db.DB()
	if err != nil {
		return nil,err
	}
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	return db,nil
}
