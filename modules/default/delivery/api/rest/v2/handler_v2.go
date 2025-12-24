package v2

import (
	"orchid-starter/modules/default/usecase"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
)

type defaultHandler struct {
	usecase usecase.DefaultUsecaseInterface
}

func NewDefaultHandler(u usecase.DefaultUsecaseInterface) *defaultHandler {
	return &defaultHandler{
		usecase: u,
	}
}

func (base *defaultHandler) Welcome(irisCtx iris.Context) {
	ctx := irisCtx.Request().Context()
	sange.NewResponse(irisCtx, iris.StatusOK, base.usecase.WelcomeUsecase(ctx))
}
