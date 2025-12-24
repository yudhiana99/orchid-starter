package http

import (
	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
	"net/http"
)

func HealthCheckHandler(ctx iris.Context) {
	sange.NewResponse(ctx, http.StatusOK, "all good...")
}
