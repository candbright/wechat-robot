package db

import (
	"github.com/candbright/wechat-robot/internal/server/config"
	"github.com/candbright/wechat-robot/internal/server/db/mysql"
	"github.com/candbright/wechat-robot/internal/server/db/options"
	"github.com/candbright/wechat-robot/internal/server/repo"
	"github.com/pkg/errors"
)

type DB interface {
	AddIdiom(data repo.Idiom) error
	GetIdioms(opts ...options.Option) ([]repo.Idiom, error)
}

func NewDB() DB {
	var db DB
	var err error
	switch config.Config.Get("db.driver") {
	case "mysql":
		db, err = mysql.NewDB()
		if err != nil {
			panic(err)
		}
	default:
		panic(errors.New("invalid db driver"))
	}
	return db
}
