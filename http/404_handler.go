package http

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
)

func NotFoundHandler(ctx iris.Context) {
	sange.NewResponse(ctx, http.StatusNotFound, "You just got lost in orchid-starter. Nothing to see here...")
}

func NotFoundApiVersion(ctx iris.Context) {
	sange.NewResponse(ctx, http.StatusNotImplemented, "Your API version not supported in orchid-starter...")
}
