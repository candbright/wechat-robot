package main

import (
	"github.com/candbright/go-core/rest/handler"
	"github.com/candbright/go-log/log"
	"github.com/candbright/go-log/options"
	"github.com/candbright/wechat-robot/internal/server/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	handler.AppName(config.Config.Get("application.name"))
	err := log.Init(
		options.Path(config.Config.Get("log.path")),
		options.Level(func() logrus.Level {
			level, err := logrus.ParseLevel(config.Config.Get("log.level"))
			if err != nil {
				return logrus.InfoLevel
			}
			return level
		}),
		options.Format(&logrus.TextFormatter{}),
	)
	if err != nil {
		panic(err)
	}
	engine := gin.New()
	engine.Use(handler.LogHandler())
	RegisterHandlers(engine)
	log.Debug("start application " + config.Config.Get("application.name"))
	_ = engine.Run(":" + strconv.Itoa(config.Config.GetInt("application.port")))
}
