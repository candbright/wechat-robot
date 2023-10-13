package service

import (
	"github.com/candbright/wechat-robot/internal/server/dao"
	"github.com/candbright/wechat-robot/internal/server/repo"
)

type IdiomService struct {
	IdiomDao *dao.IdiomDao
}

func (service *IdiomService) RandomIdiom() (repo.Idiom, error) {
	return service.IdiomDao.RandomIdiom()
}
