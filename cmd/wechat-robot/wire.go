//go:build wireinject

package main

import (
	"github.com/candbright/go-core/rest"
	"github.com/candbright/wechat-robot/internal/server/dao"
	"github.com/candbright/wechat-robot/internal/server/db"
	"github.com/candbright/wechat-robot/internal/server/handler"
	"github.com/candbright/wechat-robot/internal/server/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func RegisterHandlers(engine *gin.Engine) {
	rest.RegisterHandler(engine, "idiom", setupIdiomHandler())
}

func setupIdiomHandler() *handler.IdiomHandler {
	wire.Build(
		wire.Struct(new(handler.IdiomHandler), "*"),
		wire.Struct(new(service.IdiomService), "*"),
		wire.Struct(new(dao.IdiomDao), "*"),
		db.NewDB,
	)
	return nil
}
