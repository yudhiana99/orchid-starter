package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
)

type Module func(ctx iris.Context) *sange.Error

type Base struct {
	Modules []Module
}

func New() *Base {
	return &Base{}
}

func (base *Base) AddModule(f Module) *Base {
	base.Modules = append(base.Modules, f)
	return base
}

func (base *Base) DeclareModule(modules []Module) *Base {
	base.Modules = append(base.Modules, modules...)
	return base
}

func (base *Base) Activated(ctx iris.Context) {
	for _, module := range base.Modules {
		if err := module(ctx); err != nil {
			if err.ExtraData != nil {
				sange.NewResponse(ctx, err.StatusCode, err.ExtraData)
				ctx.StopExecution()
				return
			}

			sange.NewResponse(ctx, err.StatusCode, err.ErrorCause)
			ctx.StopExecution()
			return
		}
	}
	ctx.Next()
}
