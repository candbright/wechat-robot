package dao

import (
	"github.com/candbright/wechat-robot/internal/server/db"
	"github.com/candbright/wechat-robot/internal/server/db/options"
	"github.com/candbright/wechat-robot/internal/server/repo"
)

type IdiomDao struct {
	DB db.DB
}

func (dao *IdiomDao) IsIdiom(word string) (bool, error) {
	get, err := dao.DB.GetIdioms(options.Where("word", word))
	if err != nil {
		return false, err
	}
	if len(get) == 0 {
		return false, nil
	}
	return true, nil
}

func (dao *IdiomDao) RandomIdiom() (repo.Idiom, error) {
	return dao.DB.RandomIdiom()
}
