package handler

import (
	"github.com/candbright/go-core/rest"
	"github.com/candbright/wechat-robot/internal/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IdiomHandler struct {
	IdiomService *service.IdiomService
}

func (handler *IdiomHandler) Routes() []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "",
			Handler: handler.RandomIdiom,
		},
	}
}

func (handler *IdiomHandler) RandomIdiom(context *gin.Context) {
	rest.SimpleReq(context, func() (interface{}, error) {
		random, err := handler.IdiomService.RandomIdiom()
		if err != nil {
			return nil, err
		}
		return random, nil
	})
}
