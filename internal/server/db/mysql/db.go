package mysql

import (
	"fmt"
	"github.com/candbright/wechat-robot/internal/server/config"
	"github.com/candbright/wechat-robot/internal/server/db/mysql/model"
	"github.com/candbright/wechat-robot/internal/server/db/options"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DB struct {
	*gorm.DB
}

func (DB *DB) initTables() error {
	err := DB.AutoMigrate(&model.Idiom{})
	if err != nil {
		return errors.WithStack(err)
	}
	err = DB.AutoMigrate(&model.Quote{})
	if err != nil {
		return errors.WithStack(err)
	}
	err = DB.AutoMigrate(&model.Source{})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (DB *DB) options(opts ...options.Option) *gorm.DB {
	parseOptions := options.ParseOptions(opts...)
	query := DB.DB
	for key, value := range parseOptions.Where {
		query = query.Where(key+" = ?", value)
	}
	return query
}

func NewDB() (*DB, error) {
	var (
		ip       = config.Config.Get("db.mysql.ip")
		port     = config.Config.GetInt("db.mysql.port")
		userName = config.Config.Get("db.mysql.username")
		password = config.Config.Get("db.mysql.password")
		dbName   = config.Config.Get("db.mysql.db")
		params   = config.Config.Get("db.mysql.params")
	)
	var ssh string
	if ip == "" || port <= 0 {
		ssh = ""
	} else {
		ssh = fmt.Sprintf("tcp(%s:%d)", ip, port)
	}
	dsn := fmt.Sprintf("%s:%s@%s/%s%s", userName, password, ssh, dbName, params)
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // default size for string fields
		DisableDatetimePrecision:  true, // disable datetime precision, witch not supported before MySQL 5.6
		DontSupportRenameIndex:    true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true, // `change` when rename column, rename column not supported before MuSQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: false,
			},
		),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	instance := &DB{dbConn}
	err = instance.initTables()
	if err != nil {
		return nil, err
	}
	return instance, nil
}
